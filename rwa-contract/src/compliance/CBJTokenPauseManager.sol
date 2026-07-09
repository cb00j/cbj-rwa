// SPDX-License-Identifier: MIT
pragma solidity ^0.8.26;

import {AccessControlEnumerable} from "@openzeppelin/contracts/access/extensions/AccessControlEnumerable.sol";
import {Initializable} from "@openzeppelin/contracts/proxy/utils/Initializable.sol";
import {ICBJTokenPauseManager} from "src/compliance/interfaces/ICBJTokenPauseManager.sol";

/**
 * @title TokenPauseManager
 * @notice This contract manages the pausing and unpausing of client contracts that inherit from
 *         `PauseManagerClient`. Pausers require the 'PAUSE_TOKEN_ROLE' and unpausers require the
 *         'UNPAUSE_TOKEN_ROLE'.
 */
contract CBJTokenPauseManager is
    ICBJTokenPauseManager,
    AccessControlEnumerable,
    Initializable
{
    /// Role required to pause tokens
    bytes32 public constant PAUSE_TOKEN_ROLE = keccak256("PAUSE_TOKEN_ROLE");

    /// Role required to unpause tokens
    bytes32 public constant UNPAUSE_TOKEN_ROLE =
        keccak256("UNPAUSE_TOKEN_ROLE");

    /// Tracks paused state of individual tokens
    mapping(address => bool) public pausedTokens;

    /// Tracks global paused state for all token recipient tokens
    bool public allTokensPaused;

    /**
     * @notice Constructor for implementation contract
     */
    constructor() {
        _disableInitializers();
    }

    /**
     * @notice Initialize function for proxy deployment
     * @param guardian The address which will be granted admin, pause and unpause roles
     */
    function initialize(address guardian) public initializer {
        _grantRole(DEFAULT_ADMIN_ROLE, guardian);
        _grantRole(PAUSE_TOKEN_ROLE, guardian);
        _grantRole(UNPAUSE_TOKEN_ROLE, guardian);
    }

    /**
     * @notice Pauses a token
     * @param  token The address of the token to pause
     * @dev    Only callable by addresses with the `PAUSE_TOKEN_ROLE`
     */
    function pauseToken(address token) external onlyRole(PAUSE_TOKEN_ROLE) {
        pausedTokens[token] = true;
        emit TokenPaused(token, _msgSender());
    }

    /**
     * @notice Unpauses a token
     * @param  token The address of the token to unpause
     * @dev    Only callable by addresses with the `UNPAUSE_TOKEN_ROLE`
     */
    function unpauseToken(address token) external onlyRole(UNPAUSE_TOKEN_ROLE) {
        pausedTokens[token] = false;
        emit TokenUnpaused(token, _msgSender());
    }

    /**
     * @notice Pauses all tokens
     * @dev    Only callable by addresses with the `PAUSE_TOKEN_ROLE`
     */
    function pauseAllTokens() external onlyRole(PAUSE_TOKEN_ROLE) {
        allTokensPaused = true;
        emit AllTokensPaused(_msgSender());
    }

    /**
     * @notice Unpauses all tokens
     * @dev    Only callable by addresses with the `UNPAUSE_TOKEN_ROLE`.
     */
    function unpauseAllTokens() external onlyRole(UNPAUSE_TOKEN_ROLE) {
        allTokensPaused = false;
        emit AllTokensUnpaused(_msgSender());
    }

    /**
     * @notice Checks if a specific token is paused
     * @param  token The address of the token to check
     * @return True if the token is paused, false otherwise
     * @dev    Returns true if the specific token is paused or if all tokens are paused
     */
    function isTokenPaused(
        address token
    ) external view override returns (bool) {
        return pausedTokens[token] || allTokensPaused;
    }
}
