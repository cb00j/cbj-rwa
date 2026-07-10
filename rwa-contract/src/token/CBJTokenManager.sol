// SPDX-License-Identifier: MIT
pragma solidity ^0.8.26;

import {ICBJTokenManager} from "token/interfaces/ICBJTokenManager.sol";
import {ICBJTokenLike} from "token/interfaces/ICBJTokenLike.sol";
import {AccessControlEnumerable} from "@openzeppelin/contracts/access/extensions/AccessControlEnumerable.sol";
import {Initializable} from "@openzeppelin/contracts/proxy/utils/Initializable.sol";
import {ReentrancyGuard} from "@openzeppelin/contracts/utils/ReentrancyGuard.sol";
import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import {SafeERC20} from "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";
import {ECDSA} from "@openzeppelin/contracts/utils/cryptography/ECDSA.sol";
import {MessageHashUtils} from "@openzeppelin/contracts/utils/cryptography/MessageHashUtils.sol";

contract CBJTokenManager is
    ICBJTokenManager,
    AccessControlEnumerable,
    ReentrancyGuard,
    Initializable
{}
