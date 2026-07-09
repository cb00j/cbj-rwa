// SPDX-License-Identifier: MIT
pragma solidity ^0.8.26;

import {AccessControlEnumerable} from "@openzeppelin/contracts/access/extensions/AccessControlEnumerable.sol";
import {Initializable} from "@openzeppelin/contracts/proxy/utils/Initializable.sol";
import {ICBJCompliance} from "src/compliance/interfaces/ICBJCompliance.sol";
import {ICBJBlocklist} from "src/compliance/interfaces/ICBJBlocklist.sol";
import {ICBJSanctionsList} from "src/compliance/interfaces/ICBJSanctionsList.sol";

/**
 * @title  CBJCompliance
 * @notice This contract is responsible for enforcing compliance rules for CBJ tokens and
 *         associated systems. It allows for setting blocklists and
 *         sanctions lists for CBJ tokens.
 *         Roles:
 *          - MASTER_CONFIGURE_ROLE
 *          - CBJ_MANAGE_ROLE
 */
contract CBJCompliance is
    ICBJCompliance,
    AccessControlEnumerable,
    Initializable
{
    /// Role to set the blocklist or sanctions list for any CBJ token
    bytes32 public constant MASTER_CONFIGURE_ROLE =
        keccak256("MASTER_CONFIGURE_ROLE");

    /// Role admin for roles for setting blocklists and sanctions lists for specific CBJ tokens
    bytes32 public constant CBJ_MANAGE_ROLE = keccak256("CBJ_MANAGE_ROLE");

    /**
     * @notice Mapping of CBJ token address to the role that can set the blocklist or sanctions list
     *         for specific CBJ tokens
     */
    mapping(address /*cbjToken*/ => bytes32) public cbjRole;

    /// Mapping of CBJ token address to the blocklist contract
    mapping(address => ICBJBlocklist) public cbjTokenToBlocklist;

    /// Mapping of CBJ token address to the sanctions list contract
    mapping(address => ICBJSanctionsList) public cbjTokenToSanctionsList;

    /**
     * @notice Constructor for implementation contract
     */
    constructor() {
        _disableInitializers();
    }

    /**
     * @notice Initialize function for proxy deployment
     * @param admin The address that will be granted the default admin role
     */
    function initialize(address admin) external initializer {
        _grantRole(DEFAULT_ADMIN_ROLE, admin);
    }

    /**
     * @notice Check if a user is compliant with the CBJ token's compliance rules
     * @param  cbjToken The CBJ token address
     * @param  user     The user address
     * @dev    Reverts if the user is not compliant. Does not revert for unsupported tokens.
     */
    function checkIsCompliant(
        address cbjToken,
        address user
    ) external view override {
        if (
            cbjTokenToBlocklist[cbjToken] != ICBJBlocklist(address(0)) &&
            cbjTokenToBlocklist[cbjToken].isBlocked(user)
        ) {
            revert UserBlocked();
        }

        if (
            cbjTokenToSanctionsList[cbjToken] !=
            ICBJSanctionsList(address(0)) &&
            cbjTokenToSanctionsList[cbjToken].isSanctioned(user)
        ) {
            revert UserSanctioned();
        }
    }

    /**
     * @notice Set the blocklist for a given CBJ token
     * @param  cbjToken  The CBJ token address
     * @param  blocklist The blocklist contract
     */
    function setBlocklist(address cbjToken, ICBJBlocklist blocklist) external {
        if (
            !(hasRole(cbjRole[cbjToken], _msgSender()) ||
                hasRole(MASTER_CONFIGURE_ROLE, _msgSender()))
        ) {
            revert MissingCBJOrMasterConfigurerRole();
        }

        if (cbjToken == address(0)) revert CBJAddressCannotBeZero();

        emit BlocklistSet(cbjToken, cbjTokenToBlocklist[cbjToken], blocklist);
        cbjTokenToBlocklist[cbjToken] = blocklist;
    }

    /**
     * @notice Set the sanctions list for a given CBJ token
     * @param  cbjToken      The CBJ token address
     * @param  sanctionsList The sanctions list contract
     */
    function setSanctionsList(
        address cbjToken,
        ICBJSanctionsList sanctionsList
    ) external {
        if (
            !(hasRole(cbjRole[cbjToken], _msgSender()) ||
                hasRole(MASTER_CONFIGURE_ROLE, _msgSender()))
        ) {
            revert MissingCBJOrMasterConfigurerRole();
        }

        if (cbjToken == address(0)) revert CBJAddressCannotBeZero();

        emit SanctionsListSet(
            cbjToken,
            cbjTokenToSanctionsList[cbjToken],
            sanctionsList
        );
        cbjTokenToSanctionsList[cbjToken] = sanctionsList;
    }

    /**
     * @notice Set the role for a CBJ token
     * @param  cbjToken The CBJ token address
     * @dev    This role is computed as the keccak256 hash of the CBJ token address.
     */
    function setCBJRole(address cbjToken) external onlyRole(CBJ_MANAGE_ROLE) {
        if (cbjToken == address(0)) revert CBJAddressCannotBeZero();
        bytes32 role = keccak256(abi.encodePacked(cbjToken)); // generate a unique role for the CBJ token based on its address
        _setRoleAdmin(role, CBJ_MANAGE_ROLE); // only CBJ_MANAGE_ROLE can manage this role,and will invoke function _grantRole granting the role to someone
        cbjRole[cbjToken] = role; // only the one who granted this role or admin can manage the token(setBlocklist/setSanctionsList)
        emit CBJRoleSet(cbjToken, role);
    }

    /**
     * @notice Simplified compliance check using _msgSender() as CBJ identifier
     * @param user The user address to check
     */
    function checkIsCompliant(address user) external view {
        this.checkIsCompliant(_msgSender(), user);
    }
}
