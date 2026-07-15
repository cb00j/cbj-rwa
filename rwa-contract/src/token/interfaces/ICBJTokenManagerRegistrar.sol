// SPDX-License-Identifier: MIT
pragma solidity ^0.8.26;

import {ICBJRegistrar} from "src/token/interfaces/ICBJRegistrar.sol";
import {CBJTokenManager} from "src/token/CBJTokenManager.sol";

/**
 * @title  ICBJTokenManagerRegistrar
 * @notice Interface for the CBJTokenManagerRegistrar contract containing all functions, events, and errors
 */
interface ICBJTokenManagerRegistrar is ICBJRegistrar {
    // ============ Functions ============

    /**
     * @notice Pauses the registrar, disabling registration
     */
    function pause() external;

    /**
     * @notice Unpauses the registrar, enabling registration
     */
    function unpause() external;

    /**
     * @notice Sets or updates the CBJ Token Manager address
     * @param  _cbjTokenManager The new CBJ Token Manager address
     */
    function setCBJTokenManager(address _cbjTokenManager) external;

    /**
     * @notice Returns the current CBJ Token Manager address
     * @return The CBJTokenManager contract instance
     */
    function cbjTokenManager() external view returns (CBJTokenManager);

    /**
     * @notice Returns the CONFIGURE_ROLE constant
     * @return The bytes32 value of CONFIGURE_ROLE
     */
    function CONFIGURE_ROLE() external view returns (bytes32);

    /**
     * @notice Returns the TOKEN_FACTORY_ROLE constant
     * @return The bytes32 value of TOKEN_FACTORY_ROLE
     */
    function TOKEN_REGISTER_ROLE() external view returns (bytes32);

    /**
     * @notice Returns the PAUSER_ROLE constant
     * @return The bytes32 value of PAUSER_ROLE
     */
    function PAUSE_ROLE() external view returns (bytes32);

    /**
     * @notice Returns the UNPAUSER_ROLE constant
     * @return The bytes32 value of UNPAUSER_ROLE
     */
    function UNPAUSE_ROLE() external view returns (bytes32);

    // ============ Events ============

    /**
     * @notice Emitted when the `CBJTokenManager` is set
     * @param  oldManager The old `CBJTokenManager` address
     * @param  newManager The new `CBJTokenManager` address
     */
    event CBJTokenManagerSet(
        address indexed oldManager,
        address indexed newManager
    );

    /**
     * @notice Emitted when a new token is registered
     * @param token The address of the token that was registered following a deployment
     */
    event TokenRegistered(address indexed token);

    // ============ Errors ============

    /// Error thrown when attempting to set the CBJ Token Manager to zero address
    error CBJTokenManagerCannotBeZero();

    /// Error thrown when attempting to register a token with zero address
    error TokenAddressCannotBeZero();
}
