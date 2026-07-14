// SPDX-License-Identifier: MIT
pragma solidity ^0.8.26;

import {PendingToken} from "gateway/PendingToken.sol";

/**
 * @title  PendingUSDC
 * @notice A non-transferable token representing pending USDC during withdrawal processing
 */
contract PendingUSDC is PendingToken {
    /**
     * @notice Constructor
     * @param _gatewayContract The address of the Gateway contract
     */
    constructor(
        address _gatewayContract
    ) PendingToken("Pending USDC", "pendingUSDC", _gatewayContract) {
        // All initialization is handled by the parent contract
    }
}
