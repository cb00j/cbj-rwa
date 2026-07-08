// SPDX-License-Identifier: MIT
pragma solidity ^0.8.26;
import {IAccessControl} from "@openzeppelin/contracts/access/IAccessControl.sol";
import {ICBJBlocklist} from "./ICBJBlocklist.sol";
import {ICBJSanctionsList} from "./ICBJSanctionsList.sol";

interface ICBJCompliance is IAccessControl {
    // ============ Functions ============
    /**
     * @notice Check if a user is compliant with the CBJ token's compliance rules
     * @param  cbjToken The CBJ token address
     * @param  user     The user address
     * @dev    Reverts if the user is not compliant. Does not revert for unsupported tokens.
     */
    function checkIsCompliant(address cbjToken, address user) external view;

    /**
     * @notice Simplified compliance check using _msgSender() as CBJ identifier
     * @param user The user address to check
     */
    function checkIsCompliant(address user) external view;

    /**
     * @notice Set the blocklist for a given CBJ token
     * @param  cbjToken  The CBJ token address
     * @param  blocklist The blocklist contract
     */
    function setBlockList(address cbjToken, ICBJBlocklist blocklist) external;

    /**
     * @notice Set the sanctions list for a given CBJ token
     * @param  cbjToken      The CBJ token address
     * @param  sanctionsList The sanctions list contract
     */
    function setSanctionsList(
        address cbjToken,
        ICBJSanctionsList sanctionsList
    ) external;

    // ============ View Functions ============
    /**
     * @notice Set the role for a CBJ token
     * @param  cbjToken The CBJ token address
     * @dev    This role is computed as the keccak256 hash of the CBJ token address.
     */
    function setCBJRole(address cbjToken) external;

    /**
     * @notice Returns the MASTER_CONFIGURE_ROLE constant
     * @return The bytes32 value of MASTER_CONFIGURE_ROLE
     */
    function MASTER_CONFIGURE_ROLE() external view returns (bytes32);

    /**
     * @notice Returns the CBJ_MANAGE_ROLE constant
     * @return The bytes32 value of CBJ_MANAGE_ROLE
     */
    function CBJ_MANAGE_ROLE() external view returns (bytes32);

    /**
     * @notice Returns the role for a specific CBJ token
     * @param cbjToken The CBJ token address
     * @return The role bytes32 for the token
     */
    function cbjRole(address cbjToken) external view returns (bytes32);

    /**
     * @notice Returns the blocklist contract for a specific CBJ token
     * @param cbjToken The CBJ token address
     * @return The blocklist contract
     */
    function cbjTokenToBlocklist(
        address cbjToken
    ) external view returns (ICBJBlocklist);

    function cbjTokenToSanctionsList(
        address cbjToken
    ) external view returns (ICBJSanctionsList);

    // ============ Events ============

    /**
     * @notice Emitted when the role is set for an CBJ token
     * @param  cbjToken The CBJ token address
     * @param  role     The role that was set - keccak256 hash of the CBJ token address
     */
    event CBJRoleSet(address indexed cbjToken, bytes32 role);

    /**
     * @notice Emitted when the blocklist is set for an CBJ token
     * @param  cbjToken     The CBJ token address for which the blocklist was set
     * @param  oldBlocklist The old blocklist contract
     * @param  newBlocklist The new blocklist contract
     */
    event BlocklistSet(
        address indexed cbjToken,
        ICBJBlocklist oldBlocklist,
        ICBJBlocklist newBlocklist
    );

    /**
     * @notice Emitted when the sanctions list is set for an CBJ token
     * @param  cbjToken         The CBJ token address for which the sanctions list was set
     * @param  oldSanctionsList The old sanctions list contract
     * @param  newSanctionsList The new sanctions list contract
     */
    event SanctionsListSet(
        address indexed cbjToken,
        ICBJSanctionsList oldSanctionsList,
        ICBJSanctionsList newSanctionsList
    );

    // ============ Errors ============

    /// Error thrown when a user is blocked via the blocklist
    error UserBlocked();

    /// Error thrown when a user is sanctioned via the sanctions list
    error UserSanctioned();

    /// Error thrown when trying to set an CBJ token address of 0x0
    error CBJAddressCannotBeZero();

    /// Error thrown when trying to set configure an CBJ token without the required role
    error MissingCBJOrMasterConfigurerRole();
}
