// SPDX-License-Identifier: MIT
pragma solidity ^0.8.26;

import {PendingToken} from "src/gateway/PendingToken.sol";

/**
 * @title  PendingCbjUSDC
 * @notice A non-transferable token representing pending CBJ USDC during deposit processing
 */
contract PendingCbjUSDC is PendingToken {
    /**
     * @notice Constructor
     * @param _gatewayContract The address of the Gateway contract
     */
    constructor(
        address _gatewayContract
    ) PendingToken("Pending CBJ USDC", "pendingCbjUSDC", _gatewayContract) {
        // All initialization is handled by the parent contract
    }
}
