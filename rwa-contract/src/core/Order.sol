// SPDX-License-Identifier: MIT
pragma solidity ^0.8.26;

import {AccessControlEnumerable} from "@openzeppelin/contracts/access/extensions/AccessControlEnumerable.sol";
import {ReentrancyGuard} from "@openzeppelin/contracts/utils/ReentrancyGuard.sol";
import {Initializable} from "@openzeppelin/contracts/proxy/utils/Initializable.sol";
import {ICBJTokenLike} from "../token/interfaces/ICBJTokenLike.sol";

contract OrderContract is
    AccessControlEnumerable,
    ReentrancyGuard,
    Initializable
{
    // ============Roles===========
    bytes32 public constant BACKEND_ROLE = keccak256("BACKEND_ROLE");

    // ============Types===========
    enum Side {
        BUY,
        SELL
    }
    enum OrderType {
        MARKET,
        LIMIT
    }

    enum Status {
        PENDING,
        EXECUTED,
        CANCELREQUESTED,
        CANCELLED
    }

    enum TimeInForce {
        DAY, // Day Order
        GTC, // Good Till Cancel
        OPG, // At the Opening
        IOC, // Immediate or Cancel
        FOK, // Fill or Kill
        GTX, // Good Till Crossing/Post Only
        GTD, // Good Till Date
        CLS // At the Close
    }

    // ============Storage============
    // USDM used as the payment token for Buy orders
    ICBJTokenLike public usdm;
    // Registered CBJToken for each symbol (payment token for Sell orders)
    mapping(string => ICBJTokenLike) public symbolToToken;
    // Global auto-increment nonce for unique orderId (used as storage key)
    uint256 public nextOrderId;
    // Per-account incrementing order sequence (used for display order number)
    mapping(address => uint) public accountOrderSeq;

    struct Order {
        uint id;
        uint orderNumber; // Display-only structured order number
        address user;
        string symbol; // Business-side symbol
        uint qty;
        address escrowAsset; // Actual escrowed asset (USDM or the symbol's CBJToken)
        uint amount; // Escrowed amount (Buy: price*qty; Sell: qty), using 18 decimals
        uint price; // For Market: user's acceptable worst execution price; for Limit: limit price
        Side side;
        OrderType orderType;
        Status status;
        TimeInForce timeInForce;
    }

    // orderId => Order
    mapping(uint => Order) public orders;

    // ============Events============
    event OrderSubmitted(
        address indexed user,
        uint indexed orderId,
        string symbol,
        uint qty,
        uint price,
        Side side,
        OrderType orderType,
        TimeInForce tif,
        uint blockTimestamp
    );
    event CancelRequested(
        address indexed user,
        uint indexed orderId,
        uint blockTimestamp
    );

    event OrderExecuted(
        uint indexed orderId,
        address indexed user,
        uint refundAmount,
        uint burnAmount,
        uint mintAmount,
        TimeInForce tif
    );

    event OrderCancelled(
        uint indexed orderId,
        address indexed user,
        address asset,
        uint refundAmount,
        Side side,
        OrderType orderType,
        TimeInForce tif,
        Status previousStatus
    );

    // ============Errors============
    error AmountZero();
    error NotOwner();
    error InvalidStatus();
    error ZeroAddress();

    error AlreadyExecuted();
    error AlreadyCancelled();
    error NotCancelRequested();
    error NotFound();

    constructor() {
        _disableInitializers();
    }

    /**
     * @notice Initialize the contract (set USDM and roles)
     * @param _usdm USDM token address (payment asset for Buy orders)
     * @param _admin Admin address (granted DEFAULT_ADMIN_ROLE)
     * @param _backend Backend address (granted BACKEND_ROLE; may be zero address)
     */
    function initialize(
        address _usdm,
        address _admin,
        address _backend
    ) external initializer {
        if (_usdm == address(0)) revert ZeroAddress();
        if (_admin == address(0)) revert ZeroAddress();

        usdm = ICBJTokenLike(_usdm);

        _grantRole(DEFAULT_ADMIN_ROLE, _admin);
        if (_backend != address(0)) {
            _grantRole(BACKEND_ROLE, _backend);
        }
    }

    // ============Admin============
    /**
     * @notice Register/update the CBJToken for a given symbol
     */

    function setSymbolToken(
        string calldata symbol,
        address token
    ) external onlyRole(DEFAULT_ADMIN_ROLE) {
        if (token == address(0)) revert ZeroAddress();
        symbolToToken[symbol] = ICBJTokenLike(token);
    }

    function setBackend(
        address backend_
    ) external onlyRole(DEFAULT_ADMIN_ROLE) {
        if (backend_ == address(0)) revert ZeroAddress();
        _grantRole(BACKEND_ROLE, backend_);
    }

    // ============ User Flow ============
    /**
     * @notice Submit an order and escrow funds to the contract
     * @param symbol Trading symbol (string), used for events and queries
     * @param qty Order quantity (18 decimals); for Buy, escrowed funds are price*qty; for Sell, escrowed funds are qty
     * @param side Buy/Sell
     * @param orderType Market/Limit
     * @param price Price (18 decimals); for Market, it's the acceptable worst price; for Limit, it is the limit price
     * @dev The user must pre-approve this contract to spend the corresponding asset (Buy: USDM price*qty; Sell: the symbol's PocToken qty)
     */
    function submitOrder(
        string calldata symbol,
        uint qty,
        uint price,
        Side side,
        OrderType orderType,
        TimeInForce tif
    ) external nonReentrant returns (uint orderId) {
        if (qty == 0) revert AmountZero();
        if (price == 0) revert AmountZero();

        ICBJTokenLike token = side == Side.BUY ? usdm : symbolToToken[symbol];
        if (address(0) == address(token)) revert ZeroAddress();

        uint amount = side == Side.BUY ? price * qty : qty;
        if (amount == 0) revert AmountZero();

        require(
            token.transferFrom(msg.sender, address(this), amount),
            "TRANSFER_FROM_FAIL"
        );

        // Create order with globally unique auto-increment ID
        uint seq = ++accountOrderSeq[msg.sender];
        orderId = ++nextOrderId;

        orders[orderId] = Order({
            id: orderId,
            orderNumber: _composeOrderId(msg.sender, orderType, seq),
            user: msg.sender,
            symbol: symbol,
            qty: qty,
            escrowAsset: address(token),
            amount: amount,
            price: price,
            side: side,
            orderType: orderType,
            status: Status.PENDING,
            timeInForce: tif
        });

        emit OrderSubmitted(
            msg.sender,
            orderId,
            symbol,
            qty,
            price,
            side,
            orderType,
            tif,
            block.timestamp
        );
    }

    /**
     * @notice User initiates a cancellation intent (only Pending can initiate)
     */
    function cancelOrderIntent(uint orderId) external {
        Order storage order = orders[orderId];
        if (order.user != msg.sender) revert NotOwner();
        if (order.status != Status.PENDING) revert InvalidStatus();
        order.status = Status.CANCELREQUESTED;
        emit CancelRequested(msg.sender, orderId, block.timestamp);
    }

    // ========== Backend Flow ============
    /**
     * @notice Backend marks the order as executed and burn/mint the corresponding tokens.
     * @dev Requires this contract to be granted the BURNER_ROLE for the escrow asset (USDM or symbol token)
     */
    function markExecuted(
        uint orderId,
        uint256 returnAmount,
        uint256 mintAmount
    ) external onlyRole(BACKEND_ROLE) nonReentrant {
        Order storage order = orders[orderId];
        if (order.user == address(0)) revert NotFound();
        if (order.status == Status.EXECUTED) revert AlreadyExecuted();
        if (order.status == Status.CANCELLED) revert AlreadyCancelled();
        // Set status to Executed
        order.status = Status.EXECUTED;
        // Refund any excess amount (if present)
        if (returnAmount > 0) {
            require(
                ICBJTokenLike(order.escrowAsset).transfer(
                    order.user,
                    returnAmount
                ),
                "TRANSFER_FAIL"
            );
        }

        // Burn the escrowed funds in this contract (USDM or symbol token)
        uint256 burnAmount = order.amount - returnAmount;
        if (burnAmount > 0) {
            ICBJTokenLike(order.escrowAsset).burn(burnAmount);
        }

        // Mint the corresponding tokens to the user (USDM or symbol token)
        if (mintAmount > 0) {
            ICBJTokenLike mintToken = order.side == Side.BUY
                ? symbolToToken[order.symbol]
                : usdm;
            mintToken.mint(order.user, mintAmount);
        }

        emit OrderExecuted(
            orderId,
            order.user,
            returnAmount,
            burnAmount,
            mintAmount,
            order.timeInForce
        );
    }

    /**
     * @notice Backend finally cancels the order and refunds all escrowed funds to the user (only when in CancelRequested)
     */
    function cancelOrder(
        uint orderId
    ) external onlyRole(BACKEND_ROLE) nonReentrant {
        Order storage order = orders[orderId];
        if (order.user == address(0)) revert NotFound();
        if (order.status != Status.CANCELREQUESTED) revert NotCancelRequested();
        Status previousStatus = order.status;
        order.status = Status.CANCELLED;
        require(
            ICBJTokenLike(order.escrowAsset).transfer(order.user, order.amount),
            "TRANSFER_FAIL"
        );
        emit OrderCancelled(
            orderId,
            order.user,
            order.escrowAsset,
            order.amount,
            order.side,
            order.orderType,
            order.timeInForce,
            previousStatus
        );
    }

    // ============ Views ============
    function getOrder(uint orderId) external view returns (Order memory) {
        Order storage order = orders[orderId];
        if (order.user == address(0)) revert NotFound();
        return order;
    }

    function getOrderNumber(uint orderId) external view returns (uint) {
        Order storage order = orders[orderId];
        if (order.user == address(0)) revert NotFound();
        return order.orderNumber;
    }

    // ============ Internal ============
    function _composeOrderId(
        address user,
        OrderType orderType,
        uint seq
    ) internal pure returns (uint) {
        return (uint(uint160(user)) << 96) | (uint(orderType) << 88) | seq;
    }
}
