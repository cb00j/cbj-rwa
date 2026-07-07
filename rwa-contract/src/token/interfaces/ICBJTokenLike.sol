// SPDX-License-Identifier: MIT
pragma solidity ^0.8.26;
import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";

interface ICBJTokenLike is IERC20 {
    function mint(address _to, uint256 _amount) external;

    function burn(uint256 _amount) external;

    function burnFrom(address _from, uint256 _amount) external;

    function updateMultiplier(address _newMultiplier) external;
}
