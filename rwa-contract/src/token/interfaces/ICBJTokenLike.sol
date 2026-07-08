// SPDX-License-Identifier: MIT
pragma solidity ^0.8.26;
import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";

// This interface is not inherited directly by this project,instead, it is a
// subset of functions provided by all project tokens that the other contracts interact with
interface ICBJTokenLike is IERC20 {
    function mint(address to, uint256 amount) external;

    function burn(uint256 amount) external;

    function updateMultiplier(address newMultiplier) external;
}
