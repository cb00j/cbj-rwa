// SPDX-License-Identifier: MIT
pragma solidity ^0.8.26;

import {ERC20} from "@openzeppelin/contracts/token/ERC20/ERC20.sol";
import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import {AccessControlEnumerable} from "@openzeppelin/contracts/access/extensions/AccessControlEnumerable.sol";
import {Initializable} from "@openzeppelin/contracts/proxy/utils/Initializable.sol";
import {IERC20Metadata} from "@openzeppelin/contracts/token/ERC20/extensions/IERC20Metadata.sol";

import {ICBJToken} from "src/token/interfaces/ICBJToken.sol";
import {ICBJCompliance} from "src/compliance/interfaces/ICBJCompliance.sol";
import {ICBJTokenPauseManager} from "src/compliance/interfaces/ICBJTokenPauseManager.sol";

contract CBJToken is ERC20, AccessControlEnumerable, Initializable, ICBJToken {
    /// Role for changing the token name, symbol, compliance and token pause manager
    bytes32 public constant CONFIGURE_ROLE = keccak256("CONFIGURE_ROLE");
    /// Role for burning tokens
    bytes32 public constant BURN_ROLE = keccak256("BURN_ROLE");
    /// Role for minting tokens
    bytes32 public constant MINT_ROLE = keccak256("MINT_ROLE");
    /// Role for updating multiplier
    bytes32 public constant MULTIPLIER_UPDATE_ROLE =
        keccak256("MULTIPLIER_UPDATE_ROLE");

    /// Override for the name allowing the name to be changed
    string internal _name;
    /// Override for the symbol allowing the symbol to be changed
    string internal _symbol;

    /// Rebasing related variables
    /// @dev Defines ratio between a single share of a token to balance of a token (1e18 precision)
    uint256 public multiplier;
    /// @dev Mapping of account addresses to their share amounts
    mapping(address => uint256) private _shares;
    /// @dev Total amount of shares
    uint256 internal _totalShares;
    /// @dev Multiplier nonce for tracking updates
    uint256 public multiplierNonce;

    // ============ Compliance and Pause Management ============
    /// Compliance contract
    ICBJCompliance public cbjCompliance;
    /// Token pause manager contract
    ICBJTokenPauseManager public cbjTokenPauseManager;

    constructor() ERC20("", "") {
        // Disable initializers to prevent direct initialization of implementation
        _disableInitializers();
    }

    /**
     * @notice Initialize the token with name, symbol, compliance, and pause manager
     * @param name_ The name of the token
     * @param symbol_ The symbol of the token
     * @param _cbjCompliance The compliance contract address
     * @param _tokenPauseManager The token pause manager address
     */
    function initialize(
        string memory name_,
        string memory symbol_,
        address _cbjCompliance,
        address _tokenPauseManager
    ) external initializer {
        // Grant roles to the caller (factory)
        _grantRole(DEFAULT_ADMIN_ROLE, _msgSender());
        _grantRole(CONFIGURE_ROLE, _msgSender());
        _grantRole(MINT_ROLE, _msgSender());
        _grantRole(BURN_ROLE, _msgSender());
        _grantRole(MULTIPLIER_UPDATE_ROLE, _msgSender());

        // Set the name and symbol
        _name = name_;
        _symbol = symbol_;

        // Initialize multiplier to 1e18 (1:1 ratio)
        multiplier = 1e18;
        multiplierNonce = 0;

        // Set compliance and token pause manager
        _setCompliance(_cbjCompliance);
        _setTokenPauseManager(_tokenPauseManager);
    }

    /**
     * @notice Returns the name of the token. Overrides the default name allowing the name to be changed
     *      after deployment.
     */
    function name()
        public
        view
        override(ERC20, IERC20Metadata)
        returns (string memory)
    {
        return _name;
    }

    /**
     * @notice Sets the token name
     *
     * @param name_ New token name
     */
    function setName(string memory name_) external onlyRole(CONFIGURE_ROLE) {
        emit NameChanged(_name, name_);
        _name = name_;
    }

    /**
     * @notice Returns the symbol of the token. Overrides the default symbol allowing the symbol to be changed
     *      after deployment.
     */
    function symbol()
        public
        view
        override(ERC20, IERC20Metadata)
        returns (string memory)
    {
        return _symbol;
    }

    /**
     * @notice Sets the token symbol
     *
     * @param symbol_ New token symbol
     */
    function setSymbol(
        string memory symbol_
    ) external onlyRole(CONFIGURE_ROLE) {
        emit SymbolChanged(_symbol, symbol_);
        _symbol = symbol_;
    }

    /**
     * @notice Returns the amount of tokens owned by `account` (rebased balance)
     * @param account The account to query the balance of
     * @return The rebased token balance
     */
    function balanceOf(
        address account
    ) public view override(ERC20, IERC20) returns (uint256) {
        return _getUnderlyingAmountByShares(_shares[account], multiplier);
    }

    /**
     * @notice Returns the total amount of tokens in existence (rebased total supply)
     * @return The rebased total supply
     */
    function totalSupply()
        public
        view
        virtual
        override(ERC20, IERC20)
        returns (uint256)
    {
        return _getUnderlyingAmountByShares(_totalShares, multiplier);
    }

    /**
     * @notice Returns amount of shares owned by given account
     * @param account The account to query shares for
     * @return The amount of shares owned by the account
     */
    function sharesOf(address account) public view returns (uint256) {
        return _shares[account];
    }

    /**
     * @notice Returns the total amount of shares
     * @return The total shares amount
     */
    function totalShares() public view returns (uint256) {
        return _totalShares;
    }

    /**
     * @notice Transfers shares to a specified address
     * @param to The address to transfer shares to
     * @param sharesAmount The amount of shares to transfer
     * @return True if the transfer was successful
     */
    function transferShares(
        address to,
        uint256 sharesAmount
    ) external returns (bool) {
        uint256 underlyingAmount = _getUnderlyingAmountByShares(
            sharesAmount,
            multiplier
        );
        _transfer(_msgSender(), to, underlyingAmount);
        return true;
    }

    /**
     * @notice Returns the amount of shares that corresponds to underlying amount
     * @param underlyingAmount The underlying token amount
     * @return The corresponding shares amount
     */
    function getSharesByUnderlyingAmount(
        uint256 underlyingAmount
    ) external view returns (uint256) {
        return _getSharesByUnderlyingAmount(underlyingAmount, multiplier);
    }

    /**
     * @notice Returns the amount of underlying that corresponds to shares amount
     * @param sharesAmount The shares amount
     * @return The corresponding underlying token amount
     */
    function getUnderlyingAmountByShares(
        uint256 sharesAmount
    ) external view returns (uint256) {
        return _getUnderlyingAmountByShares(sharesAmount, multiplier);
    }

    /**
     * @notice Sets the compliance address
     *
     * @param _cbjCompliance New compliance address
     */
    function setCompliance(
        address _cbjCompliance
    ) external onlyRole(CONFIGURE_ROLE) {
        _setCompliance(_cbjCompliance);
    }

    /**
     * @notice Sets the token pause manager address
     *
     * @param _tokenPauseManager New token pause manager address
     */
    function setTokenPauseManager(
        address _tokenPauseManager
    ) external onlyRole(CONFIGURE_ROLE) {
        _setTokenPauseManager(_tokenPauseManager);
    }

    /**
     * @notice Mints a specific amount of tokens
     *
     * @param to The account who will receive the minted tokens
     * @param amount The amount of tokens to be minted
     */
    function mint(address to, uint256 amount) external onlyRole(MINT_ROLE) {
        _mint(to, amount);
    }

    /**
     * @notice Burns a specific amount of tokens
     *
     * @param amount The amount of tokens to be burned
     *
     * @dev This function can be considered an admin-burn and is only callable
     *      by an address with the `BURN_ROLE`
     */
    function burn(uint256 amount) external onlyRole(BURN_ROLE) {
        _burn(_msgSender(), amount);
    }

    /**
     * @notice Updates the multiplier value
     * @param newMultiplier The new multiplier value
     * @dev Only callable by accounts with MULTIPLIER_UPDATE_ROLE
     */
    function updateMultiplier(
        uint256 newMultiplier
    ) external onlyRole(MULTIPLIER_UPDATE_ROLE) {
        multiplier = newMultiplier;
        multiplierNonce += 1;
        emit MultiplierUpdated(newMultiplier);
    }

    /**
     * @dev Override _update to work with shares and compliance checks
     */
    function _update(
        address from,
        address to,
        uint256 amount
    ) internal override(ERC20) {
        if (from != address(0) || to != address(0)) {
            _beforeTokenTransfer(from, to, amount);
        }

        if (from == address(0)) {
            // Minting
            if (to == address(0)) revert MintToCannotBeZero();
            uint256 mintShares = _getSharesByUnderlyingAmount(
                amount,
                multiplier
            );

            _shares[to] += mintShares;
            _totalShares += mintShares;
            emit Transfer(address(0), to, amount);
            emit TransferShares(address(0), to, mintShares);
        } else if (to == address(0)) {
            // Burning
            uint256 burnShares = _getSharesByUnderlyingAmount(
                amount,
                multiplier
            );
            uint256 currentShares = _shares[from];
            if (currentShares < burnShares) revert BurnAmountExceedsBalance();
            unchecked {
                _shares[from] = currentShares - burnShares;
            }

            _totalShares -= burnShares;
            emit Transfer(from, address(0), amount);
            emit TransferShares(from, address(0), burnShares);
        } else {
            // Transferring
            uint256 transferShares = _getSharesByUnderlyingAmount(
                amount,
                multiplier
            );

            uint256 currentShares = _shares[from];
            if (currentShares < transferShares)
                revert TransferAmountExceedsBalance();
            unchecked {
                _shares[from] = currentShares - transferShares;
            }
            _shares[to] += transferShares;
            emit Transfer(from, to, amount);
            emit TransferShares(from, to, transferShares);
        }

        if (from != address(0) || to != address(0)) {
            _afterTokenTransfer(from, to, amount);
        }
    }

    function _beforeTokenTransfer(
        address from,
        address to,
        uint256 amount
    ) internal virtual {
        _checkTokenIsPaused();
        // Check constraints when `transferFrom` is called to facilitate
        // a transfer between two parties that are not `from` or `to`.
        if (from != _msgSender() && to != _msgSender()) {
            _checkIsCompliant(_msgSender());
        }

        if (from != address(0)) {
            // If not minting
            _checkIsCompliant(from);
        }

        if (to != address(0)) {
            // If not burning
            _checkIsCompliant(to);
        }
    }

    function _afterTokenTransfer(
        address from,
        address to,
        uint256 amount
    ) internal virtual {
        // Custom logic after token transfer
        // e.g. token holders tracking, interaction with other contracts, etc.
    }

    /**
     * @dev Internal function to calculate shares from underlying amount
     * @param underlyingAmount The underlying token amount
     * @param _multiplier The multiplier to use for calculation
     * @return The corresponding shares amount
     */
    function _getSharesByUnderlyingAmount(
        uint256 underlyingAmount,
        uint256 _multiplier
    ) internal pure returns (uint256) {
        return (underlyingAmount * 1e18) / _multiplier;
    }

    /**
     * @notice Checks if the token contract is paused
     * @dev Reverts with `TokenPaused` if the contract is paused
     */
    function _checkTokenIsPaused() internal view {
        if (cbjTokenPauseManager.isTokenPaused(address(this))) {
            revert TokenPaused();
        }
    }

    /**
     * @notice Checks whether an address has been blocked
     * @param account The account to check
     * @dev This function will revert if the account is not compliant
     */
    function _checkIsCompliant(address account) internal view {
        cbjCompliance.checkIsCompliant(account);
    }

    // ============ Compliance and Pause Management Functions ============
    /**
     * @notice Sets the compliance address for this token
     * @param _cbjCompliance The new compliance address
     */
    function _setCompliance(address _cbjCompliance) internal {
        if (_cbjCompliance == address(0)) {
            revert ComplianceCannotBeZero();
        }

        emit ComplianceSet(address(cbjCompliance), _cbjCompliance);
        cbjCompliance = ICBJCompliance(_cbjCompliance);
    }

    /**
     * @notice Sets the token pause manager address for this token
     * @param _tokenPauseManager The new token pause manager address
     */
    function _setTokenPauseManager(address _tokenPauseManager) internal {
        if (_tokenPauseManager == address(0))
            revert TokenPauseManagerCannotBeZero();

        emit TokenPauseManagerSet(
            address(cbjTokenPauseManager),
            _tokenPauseManager
        );

        cbjTokenPauseManager = ICBJTokenPauseManager(_tokenPauseManager);
    }

    /**
     * @dev Internal function to calculate underlying amount from shares
     * @param sharesAmount The shares amount
     * @param _multiplier The multiplier to use for calculation
     * @return The corresponding underlying token amount
     */
    function _getUnderlyingAmountByShares(
        uint256 sharesAmount,
        uint256 _multiplier
    ) internal pure returns (uint256) {
        return (sharesAmount * _multiplier) / 1e18;
    }
}
