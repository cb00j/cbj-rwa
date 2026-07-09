// SPDX-License-Identifier: MIT
pragma solidity ^0.8.26;

import {ICBJBlocklist} from "src/compliance/interfaces/ICBJBlocklist.sol";

import {AccessControlEnumerable} from "@openzeppelin/contracts/access/extensions/AccessControlEnumerable.sol";
import {Initializable} from "@openzeppelin/contracts/proxy/utils/Initializable.sol";

/**
 * @title Blocklist
 * @notice This contract manages the blocklist status for accounts using granular role-based access control.
 */
contract CBJBlocklist is ICBJBlocklist, AccessControlEnumerable, Initializable {
    // Role constants
    bytes32 public constant BLOCKLIST_ADD_ROLE =
        keccak256("BLOCKLIST_ADD_ROLE");
    bytes32 public constant BLOCKLIST_REMOVE_ROLE =
        keccak256("BLOCKLIST_REMOVE_ROLE");

    // {<address> => is account blocked}
    mapping(address => bool) private blockedAddresses;

    constructor() {
        _disableInitializers();
    }

    /**
     * @notice Initialize function for proxy deployment
     * @param admin The address which will be granted admin and other roles
     */
    function initialize(address admin) public initializer {
        _grantRole(DEFAULT_ADMIN_ROLE, admin);
        _grantRole(BLOCKLIST_ADD_ROLE, admin);
        _grantRole(BLOCKLIST_REMOVE_ROLE, admin);
    }

    /**
     * @notice Function to add a list of accounts to the blocklist
     * @dev Can be called by DEFAULT_ADMIN_ROLE or BLOCKLIST_ADD_ROLE
     * @param accounts Array of addresses to block
     */
    function addToBlocklist(
        address[] calldata accounts
    ) public onlyRole(BLOCKLIST_ADD_ROLE) {
        for (uint256 i = 0; i < accounts.length; i++) {
            address addr = accounts[i];
            if (addr == address(0)) revert BlocklistAddAddressCannotBeZero();
            blockedAddresses[addr] = true;
        }
        emit BlockedAddressesAdded(accounts);
    }

    /**
     * @notice Function to remove a list of accounts from the blocklist
     * @dev Can be called by DEFAULT_ADMIN_ROLE or BLOCKLIST_REMOVE_ROLE
     * @param accounts Array of addresses to unblock
     */
    function removeFromBlocklist(
        address[] calldata accounts
    ) external onlyRole(BLOCKLIST_REMOVE_ROLE) {
        for (uint256 i; i < accounts.length; ++i) {
            if (accounts[i] == address(0))
                revert BlocklistRemoveAddressCannotBeZero();
            blockedAddresses[accounts[i]] = false;
        }
        emit BlockedAddressesRemoved(accounts);
    }

    /**
     * @notice Function to check if an account is blocked
     *
     * @param addr Address to check
     *
     * @return True if account is blocked, false otherwise
     */
    function isBlocked(address addr) external view returns (bool) {
        return blockedAddresses[addr];
    }
}
