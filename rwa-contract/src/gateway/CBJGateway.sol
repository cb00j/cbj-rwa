// SPDX-License-Identifier: MIT
pragma solidity ^0.8.26;
import {ICBJGateway} from "gateway/interfaces/ICBJGateway.sol";

import {PendingUSDC} from "gateway/PendingUSDC.sol";
import {PendingCbjUSDC} from "gateway/PendingCbjUSDC.sol";
import {AccessControlEnumerable} from "@openzeppelin/contracts/access/extensions/AccessControlEnumerable.sol";
import {Initializable} from "@openzeppelin/contracts/proxy/utils/Initializable.sol";
import {ReentrancyGuard} from "@openzeppelin/contracts/utils/ReentrancyGuard.sol";
import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import {IERC20Metadata} from "@openzeppelin/contracts/token/ERC20/extensions/IERC20Metadata.sol";
import {SafeERC20} from "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";
import {ICBJTokenLike} from "token/interfaces/ICBJTokenLike.sol";
import {PendingToken} from "gateway/PendingToken.sol";

/**
 * @title  CBJGateway
 * @notice This contract manages the deposit and withdrawal flows for USDC and cbjUSDC
 */
contract CBJGateway is
    ICBJGateway,
    AccessControlEnumerable,
    ReentrancyGuard,
    Initializable
{
    using SafeERC20 for IERC20;

    /// @notice Role identifier for those who can configure the contract
    bytes32 public constant CONFIGURE_ROLE = keccak256("CONFIGURE_ROLE");

    /// @notice Role identifier for those who can pause operations
    bytes32 public constant PAUSE_ROLE = keccak256("PAUSE_ROLE");

    /// @notice Role identifier for those who can process operations (backend)
    bytes32 public constant PROCESSOR_ROLE = keccak256("PROCESSOR_ROLE");

    /// @notice The address of the USDC token
    address public immutable USDC;

    /// @notice The address of the cbjUSDC token
    address public immutable CBJ_USDC;

    /// @notice The address of the pending cbjUSDC token
    address public PENDING_CBJ_USDC;

    /// @notice The address of the pending USDC token
    address public PENDING_USDC;

    /// @notice Minimum USDC amount required for deposits
    uint256 public minimumDepositAmount;

    /// @notice Minimum cbjUSDC amount required for withdrawals
    uint256 public minimumWithdrawalAmount;

    /// @notice Whether deposits are paused
    bool public depositsArePaused;

    /// @notice Whether withdrawals are paused
    bool public withdrawalsArePaused;

    /// @notice Counter for generating unique operation IDs
    uint256 private _operationCounter;

    /// @notice Mapping of operation ID to deposit operations
    mapping(bytes32 => DepositOperation) public depositOperations;

    /// @notice Mapping of operation ID to withdrawal operations
    mapping(bytes32 => WithdrawalOperation) public withdrawalOperations;

    /**
     * @notice Constructor for implementation contract
     * @param _usdc The address of the USDC token
     * @param _cbjUSDC The address of the cbjUSDC token
     */
    constructor(address _usdc, address _cbjUSDC) {
        _disableInitializers();

        // Initialize immutable variables
        USDC = _usdc;
        CBJ_USDC = _cbjUSDC;
    }

    /**
     * @notice Initialize function for proxy deployment
     * @param _usdc The USDC token address
     * @param _cbjUSDC The cbjUSDC token address
     * @param _guardian The guardian address
     * @param _minimumDepositAmount The minimum deposit amount
     * @param _minimumWithdrawalAmount The minimum withdrawal amount
     */
    function initialize(
        address _usdc,
        address _cbjUSDC,
        address _guardian,
        uint256 _minimumDepositAmount,
        uint256 _minimumWithdrawalAmount
    ) external initializer {
        if (_usdc == address(0)) revert AddressCannotBeZero();
        if (_cbjUSDC == address(0)) revert AddressCannotBeZero();
        if (_guardian == address(0)) revert AddressCannotBeZero();

        minimumDepositAmount = _minimumDepositAmount;
        minimumWithdrawalAmount = _minimumWithdrawalAmount;

        // Create pending token contracts
        PENDING_CBJ_USDC = address(new PendingCbjUSDC(address(this)));
        PENDING_USDC = address(new PendingUSDC(address(this)));

        _grantRole(DEFAULT_ADMIN_ROLE, _guardian);
        _grantRole(CONFIGURE_ROLE, _guardian);
        _grantRole(PAUSE_ROLE, _guardian);
        _grantRole(PROCESSOR_ROLE, _guardian);
    }

    /**
     * @notice Modifier to check if deposits are not paused
     */
    modifier whenDepositsNotPaused() {
        if (depositsArePaused) revert DepositsArePaused();
        _;
    }

    /**
     * @notice Modifier to check if withdrawals are not paused
     */
    modifier whenWithdrawalsNotPaused() {
        if (withdrawalsArePaused) revert WithdrawalsArePaused();
        _;
    }

    /**
     * @notice Deposit USDC and receive pending CBJ USDC
     * @param amount The amount of USDC to deposit
     * @return operationId The ID of the deposit operation
     */
    function deposit(
        uint256 amount
    )
        external
        nonReentrant
        whenDepositsNotPaused
        returns (bytes32 operationId)
    {
        if (amount == 0) revert AmountCannotBeZero();
        if (amount < minimumDepositAmount) revert DepositAmountTooSmall();

        // Generate unique operation ID
        operationId = _generateOperationId();

        // Transfer USDC from user to contract
        IERC20(USDC).safeTransferFrom(_msgSender(), address(this), amount);

        // Mint pending cbjUSDC to user (convert from USDC decimals to cbjUSDC decimals)
        uint256 mintpendingAmount = _convertDecimals(
            amount,
            USDC,
            PENDING_CBJ_USDC
        );
        ICBJTokenLike(PENDING_CBJ_USDC).mint(_msgSender(), mintpendingAmount);

        // Store deposit operation
        depositOperations[operationId] = DepositOperation({
            user: _msgSender(),
            usdcAmount: amount,
            pendingCbjUSDCAmount: mintpendingAmount,
            status: OperationStatus.PENDING,
            timestamp: block.timestamp
        });

        emit PendingDeposit(
            operationId,
            _msgSender(),
            amount,
            mintpendingAmount
        );
    }

    /**
     * @notice Process pending deposit (backend function)
     * @param operationId The ID of the deposit operation to process
     * @param cbjUSDCAmount The amount of cbjUSDC to mint
     */
    function processDeposit(
        bytes32 operationId,
        uint256 cbjUSDCAmount
    ) external nonReentrant onlyRole(PROCESSOR_ROLE) {
        DepositOperation storage depositOperation = depositOperations[
            operationId
        ];
        if (depositOperation.user == address(0)) revert InvalidOperationId();
        if (depositOperation.status != OperationStatus.PENDING)
            revert InvalidOperationStatus();

        // TODO: There should be some logic to transfer USDC to broker

        // Burn pending cbjUSDC directly from user's balance
        PendingToken(PENDING_CBJ_USDC).burnFrom(
            depositOperation.user,
            depositOperation.pendingCbjUSDCAmount
        );

        // Mint cbjUSDC to user
        ICBJTokenLike(CBJ_USDC).mint(depositOperation.user, cbjUSDCAmount);

        depositOperation.status = OperationStatus.ACTIVE;
        emit DepositProcessed(
            operationId,
            depositOperation.user,
            cbjUSDCAmount
        );
    }

    /**
     * @notice Withdraw cbjUSDC and receive pending USDC
     * @param cbjUSDCAmount The amount of cbjUSDC to withdraw
     * @return operationId The ID of the withdrawal operation
     */
    function withdraw(
        uint256 cbjUSDCAmount
    )
        external
        nonReentrant
        whenWithdrawalsNotPaused
        returns (bytes32 operationId)
    {
        if (cbjUSDCAmount == 0) revert AmountCannotBeZero();
        if (cbjUSDCAmount < minimumWithdrawalAmount)
            revert WithdrawalAmountTooSmall();

        // Generate unique operation ID
        operationId = _generateOperationId();
        // Burn cbjUSDC from user
        IERC20(CBJ_USDC).safeTransferFrom(
            _msgSender(),
            address(this),
            cbjUSDCAmount
        );
        ICBJTokenLike(CBJ_USDC).burn(cbjUSDCAmount);

        uint256 pendingUSDCAmount = _convertDecimals(
            cbjUSDCAmount,
            CBJ_USDC,
            USDC
        );
        // Mint pending USDC to user (convert from cbjUSDC decimals to USDC decimals)
        ICBJTokenLike(PENDING_USDC).mint(_msgSender(), pendingUSDCAmount);

        // Store withdrawal operation
        withdrawalOperations[operationId] = WithdrawalOperation({
            user: _msgSender(),
            cbjUSDCAmount: cbjUSDCAmount,
            pendingUSDCAmount: pendingUSDCAmount,
            status: OperationStatus.PENDING,
            timestamp: block.timestamp
        });

        emit PendingWithdraw(
            operationId,
            _msgSender(),
            cbjUSDCAmount,
            pendingUSDCAmount
        );
    }

    /**
     * @notice Process pending withdrawal (backend function)
     * @param operationId The ID of the withdrawal operation to process
     * @param usdcAmount The amount of USDC to transfer back
     */
    function processWithdrawal(
        bytes32 operationId,
        uint256 usdcAmount
    ) external nonReentrant onlyRole(PROCESSOR_ROLE) {
        WithdrawalOperation storage operation = withdrawalOperations[
            operationId
        ];
        if (operation.user == address(0)) revert InvalidOperationId();
        if (operation.status != OperationStatus.PENDING)
            revert OperationAlreadyProcessed();
        if (usdcAmount == 0) revert AmountCannotBeZero();

        // Burn pending USDC directly from user's balance
        PendingToken(PENDING_USDC).burnFrom(
            operation.user,
            operation.pendingUSDCAmount
        );
        // Update operation status
        operation.status = OperationStatus.ACTIVE;
        // Transfer USDC to user
        IERC20(USDC).safeTransfer(operation.user, usdcAmount);

        emit WithdrawalProcessed(operationId, operation.user, usdcAmount);
    }

    /**
     * @notice Set the minimum deposit amount
     * @param amount The minimum deposit amount in USDC
     */
    function setMinimumDepositAmount(
        uint256 amount
    ) external onlyRole(CONFIGURE_ROLE) {
        emit MinimumDepositAmountSet(minimumDepositAmount, amount);
        minimumDepositAmount = amount;
    }

    /**
     * @notice Set the minimum withdrawal amount
     * @param amount The minimum withdrawal amount in cbjUSDC
     */
    function setMinimumWithdrawalAmount(
        uint256 amount
    ) external onlyRole(CONFIGURE_ROLE) {
        emit MinimumWithdrawalAmountSet(minimumWithdrawalAmount, amount);
        minimumWithdrawalAmount = amount;
    }

    /**
     * @notice Pause deposits
     */
    function pauseDeposits() external onlyRole(PAUSE_ROLE) {
        depositsArePaused = true;
        emit DepositsPaused();
    }

    /**
     * @notice Unpause deposits
     */
    function unpauseDeposits() external onlyRole(PAUSE_ROLE) {
        depositsArePaused = false;
        emit DepositsUnpaused();
    }

    /**
     * @notice Pause withdrawals
     */
    function pauseWithdrawals() external onlyRole(PAUSE_ROLE) {
        withdrawalsArePaused = true;
        emit WithdrawalsPaused();
    }

    /**
     * @notice Unpause withdrawals
     */
    function unpauseWithdrawals() external onlyRole(PAUSE_ROLE) {
        withdrawalsArePaused = false;
        emit WithdrawalsUnpaused();
    }

    /**
     * @notice Get deposit operation details
     * @param operationId The operation ID
     * @return The deposit operation struct
     */
    function getDepositOperation(
        bytes32 operationId
    ) external view returns (DepositOperation memory) {
        return depositOperations[operationId];
    }

    /**
     * @notice Get withdrawal operation details
     * @param operationId The operation ID
     * @return The withdrawal operation struct
     */
    function getWithdrawalOperation(
        bytes32 operationId
    ) external view returns (WithdrawalOperation memory) {
        return withdrawalOperations[operationId];
    }

    /**
     * @notice Convert amount from one token's decimals to another token's decimals
     * @param amount The amount to convert
     * @param fromToken The source token address
     * @param toToken The target token address
     * @return The converted amount
     */
    function _convertDecimals(
        uint256 amount,
        address fromToken,
        address toToken
    ) internal view returns (uint256) {
        uint8 fromDecimals = IERC20Metadata(fromToken).decimals();
        uint8 toDecimals = IERC20Metadata(toToken).decimals();

        if (fromDecimals == toDecimals) {
            return amount;
        } else if (fromDecimals < toDecimals) {
            // Scale up: multiply by 10^(toDecimals - fromDecimals)
            return amount * (10 ** (toDecimals - fromDecimals));
        } else {
            // Scale down: divide by 10^(fromDecimals - toDecimals)
            return amount / (10 ** (fromDecimals - toDecimals));
        }
    }

    /**
     * @notice Generate a unique operation ID
     * @return The generated operation ID
     */
    function _generateOperationId() internal returns (bytes32) {
        return
            keccak256(
                abi.encodePacked(
                    block.timestamp,
                    block.number,
                    msg.sender,
                    ++_operationCounter
                )
            );
    }
}
