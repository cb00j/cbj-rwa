// SPDX-License-Identifier: MIT
pragma solidity ^0.8.23;

import {Script} from "forge-std/Script.sol";
import {console2} from "forge-std/console2.sol";
import {Strings} from "@openzeppelin/contracts/utils/Strings.sol";

abstract contract BaseScript is Script {
    uint256 internal deployerPrivateKey;
    address internal deployer;
    address internal factory = 0x4e59b44847b379578588920cA78FbF26c0B4956C;

    // 定义统一的部署记录文件路径
    string constant DEPLOYMENT_PATH = "./deployments/contract_addresses.json";

    function setUp() public virtual {
        deployerPrivateKey = vm.envUint("PRIVATE_KEY");
        deployer = vm.addr(deployerPrivateKey);
    }

    // 保存合约地址
    function saveContract(string memory name, address addr) public {
        string memory network = _getNetworkName();

        // 1. 确保文件存在，如果不存在则初始化为空 JSON
        if (!vm.exists(DEPLOYMENT_PATH)) {
            /// forge-lint: disable-next-line(unsafe-cheatcode)
            vm.writeFile(DEPLOYMENT_PATH, "{}");
        }

        // 2. 构建 JSON 路径 (例如: .local.CBJToken)
        // 注意：路径必须以 . 开头
        string memory jsonPath = string.concat(".", network, ".", name);

        // 3. 直接将地址字符串写入该路径
        // writeJson 会自动处理：如果路径不存在则追加，如果存在则覆盖该值
        vm.writeJson(vm.toString(addr), DEPLOYMENT_PATH, jsonPath);

        console2.log(string.concat("Saved ", name, " on ", network, ":"), addr);
    }

    // 修改后的获取地址方法
    function getAddress(string memory name) public view returns (address) {
        string memory network = _getNetworkName();
        /// forge-lint: disable-next-line(unsafe-cheatcode)
        string memory jsonContent = vm.readFile(DEPLOYMENT_PATH);

        // 根据路径读取：.networkName.contractName
        string memory path = string.concat(".", network, ".", name);
        address result = vm.parseJsonAddress(jsonContent, path);

        return result;
    }

    // 内部辅助方法：将 ChainID 映射为网络名称
    function _getNetworkName() internal view returns (string memory) {
        uint256 chainId = block.chainid;
        if (chainId == 31337) return "local";
        if (chainId == 11155111) return "sepolia";
        if (chainId == 421614) return "arb_sepolia";
        if (chainId == 1) return "mainnet";

        return Strings.toString(chainId); // 默认返回 chainId 字符串
    }

    modifier broadcaster() {
        vm.startBroadcast(deployerPrivateKey);
        _;
        vm.stopBroadcast();
    }
}
