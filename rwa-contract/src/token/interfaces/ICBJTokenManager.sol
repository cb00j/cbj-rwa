// SPDX-License-Identifier: MIT

pragma solidity ^0.8.26;

/**
 * @title  ICBJTokenManager
 * @notice Comprehensive interface for the CBJTokenManager contract containing all functions, events, and errors
 */
interface ICBJTokenManager {
    // ============ Enums ============

    enum QuoteSide {
        /// Indicates that the user is buying CBJ tokens
        BUY,
        /// Indicates that the user is selling CBJ tokens
        SELL
    }

    // ============ Structs ============

    /**
     * @notice Quote struct that is signed by the attestation signer
     * @param  attestationId  The ID of the quote
     * @param  asset          The address of the CBJ token being bought or sold
     * @param  price          The price of the CBJ token in Usd with 18 decimals
     * @param  quantity       The quantity of CBJ tokens being bought or sold
     * @param  expiry         The expiration of the quote in seconds since the epoch
     * @param  side           The direction of the quote (BUY or SELL)
     * @param  nonce          The nonce for the quote
     */
    struct Quote {
        bytes32 attestationId;
        address asset;
        uint256 price;
        uint256 quantity;
        uint256 expiry;
        QuoteSide side;
        uint256 nonce;
    }

    // ============ Functions ============
    /**
     * @notice Called by users to mint CBJ tokens with USDC
     * @param  quote                The quote to mint CBJ tokens with
     * @param  signature            The signature of the quote attestation
     * @param  depositToken         The token the user is depositing (must be USDC)
     * @param  depositAmount        The amount of deposit tokens (must equal USDC value)
     * @return receivedCBJTokenAmount The amount of CBJ tokens minted
     */
    function mintWithAttestation(
        Quote calldata quote,
        bytes calldata signature,
        address depositToken,
        uint256 depositAmount
    ) external returns (uint256 receivedCBJTokenAmount);

    /**
     * @notice Called by users to redeem Anchored tokens for USDC
     * @param  quote                The quote to redeem Anchored tokens with
     * @param  signature            The signature of the quote attestation
     * @param  receiveToken         The token the user would like to receive (must be USDC)
     * @return The amount of USDC transferred
     */
    function redeemWithAttestaton(
        Quote calldata quote,
        bytes calldata signature,
        address receiveToken
    ) external returns (uint256);

    /**
     * @notice Sets whether a token is accepted for mints and redemptions on this contract
     * @param  token    The address of the token
     * @param  accepted Whether the token is accepted for mints and redemptions
     */
    function setCBJTokenRegisterationStatus(
        address token,
        bool accepted
    ) external;

    /**
     * @notice Sets the minimum amount required for a subscription
     * @param  minimumDepositAmount The minimum amount required to subscribe
     */
    function setMinimumDepositAmount(uint256 minimumDepositAmount) external;

    /**
     * @notice Sets the minimum amount to redeem
     * @param  minimumRedemptionAmount The minimum amount to redeem
     */
    function setMinimumRedemptionAmount(
        uint256 minimumRedemptionAmount
    ) external;

    /**
     * @notice Pauses minting for a specific CBJ token
     * @param  token The address of the CBJ token
     */
    function pauseCBJTokenMinting(address token) external;

    /**
     * @notice Unpauses minting for a specific CBJ token
     * @param  token The address of the CBJ token
     */
    function unpauseCBJTokenMinting(address token) external;

    /**
     * @notice Pauses redemptions for a specific CBJ token
     * @param  token The address of the CBJ token
     */
    function pauseCBJTokenRedemptions(address token) external;

    /**
     * @notice Unpauses redemptions for a specific CBJ token
     * @param  token The address of the CBJ token
     */
    function unpauseCBJTokenRedemptions(address token) external;

    /**
     * @notice Updates the multiplier for a specific CBJ token
     * @param  token The address of the CBJ token
     * @param  newMultiplier The new multiplier value
     */
    function updateMultiplier(address token, uint256 newMultiplier) external;

    /**
     * @notice Initialize function for proxy deployment
     * @param usdc The address of the USDC token
     * @param guardian The address which will be granted admin and other roles
     * @param minimumDepositAmount Minimum Usd amount required to subscribe
     * @param minimumRedemptionAmount Minimum Usd amount required to redeem
     */
    function initialize(
        address usdc,
        address guardian,
        uint256 minimumDepositAmount,
        uint256 minimumRedemptionAmount
    ) external;

    // ============ View Functions ============

    /**
     * @notice Returns whether an attestation ID has been executed
     * @param attestationId The attestation ID to check
     * @return True if the attestation has been executed
     */
    function executedAttestationIds(
        bytes32 attestationId
    ) external view returns (bool);

    /**
     * @notice Returns whether minting is paused for a specific CBJ token
     * @param token The address of the CBJ token
     * @return True if minting is paused for the token
     */
    function cbjTokenMintingPaused(address token) external view returns (bool);

    /**
     * @notice Returns whether redemptions are paused for a specific CBJ token
     * @param token The address of the CBJ token
     * @return True if redemptions are paused for the token
     */
    function cbjTokenRedemptionsPaused(
        address token
    ) external view returns (bool);

    /**
     * @notice Returns whether an CBJ token is accepted for minting and redemptions
     * @param token The address of the CBJ token
     * @return True if the token is accepted
     */
    function cbjTokenAccepted(address token) external view returns (bool);

    // ============ Events ============

    /**
     * @notice Event emitted when a trade is executed with an attestation
     * @param  asset          The address of the Anchored token being bought or sold
     * @param  user           The user executing the trade
     * @param  side           The direction of the quote (BUY or SELL)
     * @param  quantity       The quantity of Anchored tokens being bought or sold
     * @param  price          The price of the Anchored token in Usd with 18 decimals
     */
    event TradeExecuted(
        address asset,
        address user,
        QuoteSide side,
        uint256 quantity,
        uint256 price
    );

    /**
     * @notice Event emitted when subscription minimum is set
     * @param  oldMinDepositAmount Old subscription minimum
     * @param  newMinDepositAmount New subscription minimum
     */
    event MinimumDepositAmountSet(
        uint256 indexed oldMinDepositAmount,
        uint256 indexed newMinDepositAmount
    );

    /**
     * @notice Event emitted when redeem minimum is set
     * @param  oldMinRedemptionAmount Old redeem minimum
     * @param  newMinRedemptionAmount New redeem minimum
     */
    event MinimumRedemptionAmountSet(
        uint256 indexed oldMinRedemptionAmount,
        uint256 indexed newMinRedemptionAmount
    );

    /**
     * @notice Event emitted when the accepted CBJ token is set
     * @param  cbjToken    The address of the CBJ token
     * @param  registered Whether the CBJ token is registered
     */
    event CBJTokenRegistered(address indexed cbjToken, bool indexed registered);

    /**
     * @notice Event emitted when minting is paused for a specific CBJ token
     * @param cbjToken The address of the CBJ token
     */
    event CBJTokenMintingPaused(address indexed cbjToken);

    /**
     * @notice Event emitted when minting is unpaused for a specific CBJ token
     * @param cbjToken The address of the CBJ token
     */
    event CBJTokenMintingUnpaused(address indexed cbjToken);

    /**
     * @notice Event emitted when redemption is paused for a specific CBJ token
     * @param cbjToken The address of the CBJ token
     */
    event CBJTokenRedemptionsPaused(address indexed cbjToken);

    /**
     * @notice Event emitted when redemption is unpaused for a specific CBJ token
     * @param cbjToken The address of the CBJ token
     */
    event CBJTokenRedemptionsUnpaused(address indexed cbjToken);

    /**
     * @notice Event emitted when multiplier is updated for a specific CBJ token
     * @param cbjToken The address of the CBJ token
     * @param newMultiplier The new multiplier value
     */
    event MultiplierUpdated(
        address indexed cbjToken,
        uint256 indexed newMultiplier
    );

    // ============ Errors ============

    /// Error emitted when the token address is zero
    error TokenAddressCannotBeZero();

    /// Error emitted when the deposit amount is too small
    error DepositAmountTooSmall();

    /// @notice Thrown when the attestation has already been executed
    error AttestationAlreadyExecuted();

    /// @notice Thrown when the quote has expired
    error QuoteExpired();

    /// @notice Thrown when the signer is invalid
    error InvalidSigner();

    /// @notice Thrown when the token is not accepted
    error TokenNotAccepted();

    /// @notice Thrown when minting is paused
    error MintingPaused();

    /// @notice Thrown when redemption is paused
    error RedemptionPaused();

    /// @notice Thrown when the deposit token is invalid
    error InvalidDepositToken();

    /// @notice Thrown when the receive token is invalid
    error InvalidReceiveToken();

    /// @notice Thrown when the deposit amount is invalid
    error InvalidDepositAmount();

    /// @notice Thrown when the guardian address is zero
    error GuardianAddressCannotBeZero();

    /// Error emitted when the redemption amount is too small
    error RedemptionAmountTooSmall();

    /// Custom error for invalid quote direction
    error InvalidQuoteSide();

    /// Error emitted when the CBJ Token is not registered for minting/redemption
    error CBJTokenNotRegistered();

    /// Error emitted when attempting to set the `USDC` address to zero
    error UsdcAddressCannotBeZero();
}
