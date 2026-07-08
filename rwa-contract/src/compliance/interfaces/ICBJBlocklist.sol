// SPDX-License-Identifier: MIT
pragma solidity ^0.8.26;

/**
 * @title  ICBJBlocklist
 * @notice Comprehensive interface for the CBJBlocklist contract containing all functions, events, and errors
 */
interface ICBJBlocklist {
    // ============ Functions ============

    /**
     * @notice Function to add a list of accounts to the blocklist
     * @param accounts Array of addresses to block
     */
    function addToBlocklist(address[] calldata accounts) external;

    /**
     * @notice Function to remove a list of accounts from the blocklist
     * @param accounts Array of addresses to unblock
     */
    function removeFromBlocklist(address[] calldata accounts) external;

    // ============ View Functions ============

    /**
     * @notice Function to check if an account is blocked
     * @param addr Address to check
     * @return True if account is blocked, false otherwise
     */
    function isBlocked(address addr) external view returns (bool);

    // ============ Events ============

    /**
     * @notice Event emitted when addresses are added to the blocklist
     * @param accounts The addresses that were added to the blocklist
     */
    event BlockedAddressesAdded(address[] accounts);

    /**
     * @notice Event emitted when addresses are removed from the blocklist
     * @param accounts The addresses that were removed from the blocklist
     */
    event BlockedAddressesRemoved(address[] accounts);

    // ============ Errors ============

    /// Error thrown when attempting to add zero address to blocklist
    error BlocklistAddAddressCannotBeZero();

    /// Error thrown when attempting to remove zero address from blocklist
    error BlocklistRemoveAddressCannotBeZero();
}
