// SPDX-License-Identifier: MIT
pragma solidity ^0.8.23;

import {BaseScript} from "./BaseScript.s.sol";
import {console2} from "forge-std/console2.sol";
import {MockUSDC} from "script/mocks/MockUSDC.sol";
import {CBJToken} from "src/token/CBJToken.sol";
import {CBJTokenPauseManager} from "src/token/CBJTokenPauseManager.sol";
import {TransparentUpgradeableProxy} from "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";
import {CBJBlocklist} from "src/compliance/CBJBlocklist.sol";
import {CBJSanctionsList} from "src/compliance/CBJSanctionsList.sol";
import {CBJCompliance} from "src/compliance/CBJCompliance.sol";
import {CBJCompliance} from "src/compliance/CBJCompliance.sol";
import {OrderContract} from "src/order/Order.sol";
import {CBJGateway} from "src/gateway/CBJGateway.sol";
import {CBJTokenFactory} from "src/token/CBJTokenFactory.sol";

contract DeployAll is BaseScript {
    bytes32 constant MINT_ROLE = keccak256("MINT_ROLE");
    bytes32 constant BURN_ROLE = keccak256("BURN_ROLE");

    address public backendAddress;
    address public proxyAdmin;

    function run() public broadcaster {
        // Backend address: defaults to deployer if not set
        backendAddress = vm.envOr("BACKEND_ADDRESS", deployer);
        // Proxy admin: defaults to deployer if not set
        proxyAdmin = vm.envOr("PROXY_ADMIN_ADDRESS", deployer);

        console2.log("Deployer:", deployer);
        console2.log("Backend:", backendAddress);
        console2.log("ProxyAdmin:", proxyAdmin);

        // =============================================
        // Step 1: Deploy MockUSDC
        // =============================================
        MockUSDC mockUSDC = new MockUSDC();
        address mockUSDCAddress = address(mockUSDC);
        saveContract("MockUSDC", mockUSDCAddress);
        console2.log("MockUSDC deployed at:", mockUSDCAddress);

        // Mint 1,000,000 USDC (6 decimals) to deployer
        mockUSDC.mint(deployer, 1_000_000 * 1e6);
        console2.log("Minted 1,000,000 USDC to deployer");

        // =============================================
        // Step 2: 合规基础设施 —— CBJToken 的必需配套,不是可选项
        // 顺序很重要:PauseManager / Blocklist / SanctionsList / Compliance
        // 必须先于 CBJToken 部署,因为 CBJToken.initialize() 需要它们的地址
        // =============================================

        CBJTokenPauseManager pauseManagerImpl = new CBJTokenPauseManager();
        bytes memory pauseManagerInitData = abi.encodeWithSelector(
            CBJTokenPauseManager.initialize.selector,
            deployer
        );
        address pauseManagerProxy = address(
            new TransparentUpgradeableProxy(
                address(pauseManagerImpl),
                proxyAdmin,
                pauseManagerInitData
            )
        );

        saveContract("CBJTokenPauseManager", pauseManagerProxy);
        console2.log("CBJTokenPauseManager deployed at:", pauseManagerProxy);

        CBJBlocklist blocklistImpl = new CBJBlocklist();
        bytes memory blocklistInitData = abi.encodeWithSelector(
            CBJBlocklist.initialize.selector,
            deployer
        );
        address blocklistProxy = address(
            new TransparentUpgradeableProxy(
                address(blocklistImpl),
                proxyAdmin,
                blocklistInitData
            )
        );
        saveContract("CBJBlocklist", blocklistProxy);
        console2.log("CBJBlocklist deployed at:", blocklistProxy);

        CBJSanctionsList sanctionsListImpl = new CBJSanctionsList();
        bytes memory sanctionsListInitData = abi.encodeWithSelector(
            CBJSanctionsList.initialize.selector,
            deployer
        );
        address sanctionsListProxy = address(
            new TransparentUpgradeableProxy(
                address(sanctionsListImpl),
                proxyAdmin,
                sanctionsListInitData
            )
        );
        saveContract("CBJSanctionsList", sanctionsListProxy);
        console2.log("CBJSanctionsList deployed at:", sanctionsListProxy);

        CBJCompliance complianceImpl = new CBJCompliance();
        bytes memory complianceInitData = abi.encodeWithSelector(
            CBJCompliance.initialize.selector,
            deployer
        );
        address complianceProxy = address(
            new TransparentUpgradeableProxy(
                address(complianceImpl),
                proxyAdmin,
                complianceInitData
            )
        );
        saveContract("CBJCompliance", complianceProxy);
        console2.log("CBJCompliance deployed at:", complianceProxy);

        CBJCompliance compliance = CBJCompliance(complianceProxy);
        compliance.grantRole(compliance.MASTER_CONFIGURE_ROLE(), deployer);

        // =============================================
        // Step 3: 使用CBJTokenFactory创建CBJToken实例 —— 一个 Beacon,三个实例(USDM / AAPL.cbj / TSLA.cbj)
        // USDM 既是 Order.sol 的支付资产,也是 CBJGateway 铸销的那个稳定币,三方共用同一个实例
        // =============================================
        CBJTokenFactory tokenFactoryImplContract = new CBJTokenFactory();
        address factoryAddress = address(tokenFactoryImplContract);

        // initialize(_guardian, _cbjCompliance, _tokenPauseManager, _tokenManagerRegistrar)
        bytes memory factoryInitData = abi.encodeWithSelector(
            CBJTokenFactory.initialize.selector,
            deployer,
            complianceProxy,
            pauseManagerProxy,
            address(0)
        );
        address tokenFactoryProxy = address(
            new TransparentUpgradeableProxy(
                factoryAddress,
                proxyAdmin,
                factoryInitData
            )
        );
        saveContract("CBJTokenFactory", tokenFactoryProxy);
        console2.log("CBJTokenFactory deployed at:", tokenFactoryProxy);

        CBJTokenFactory factory = CBJTokenFactory(tokenFactoryProxy);
        // USDM 既是 Order.sol 的支付资产,也是 CBJGateway 铸销的那个稳定币,三方共用同一个实例
        address usdmProxy = factory.deployCBJTokenIsolated(
            "USDM",
            "USDM",
            deployer
        );
        saveContract("USDM", usdmProxy);
        console2.log("USDM deployed at:", usdmProxy);

        address aaplProxy = factory.deployCBJTokenIsolated(
            "AAPL.cbj",
            "AAPL.cbj",
            deployer
        );
        saveContract("AAPL.cbj", aaplProxy);
        console2.log("AAPL.cbj deployed at:", aaplProxy);

        address tslaProxy = factory.deployCBJTokenIsolated(
            "TSLA.cbj",
            "TSLA.cbj",
            deployer
        );
        saveContract("TSLA.cbj", tslaProxy);
        console2.log("TSLA.cbj deployed at:", tslaProxy);

        // =============================================
        // Step 4: OrderContract(交易引擎)
        // =============================================
        OrderContract orderImplContract = new OrderContract();
        address orderImpl = address(orderImplContract);
        // initialize(usdm_, admin_, backend_)
        bytes memory orderInitData = abi.encodeWithSelector(
            OrderContract.initialize.selector,
            usdmProxy,
            deployer,
            backendAddress
        );
        address orderProxy = address(
            new TransparentUpgradeableProxy(
                orderImpl,
                proxyAdmin,
                orderInitData
            )
        );
        saveContract("OrderContract", orderProxy);
        console2.log("OrderContract deployed at:", orderProxy);

        // =============================================
        // Step 5: CBJGateway(USDC与USDM之间的两段式入金出金网关)
        // =============================================
        CBJGateway gatewayImplContract = new CBJGateway(
            mockUSDCAddress,
            usdmProxy
        );
        address gatewayImpl = address(gatewayImplContract);
        bytes memory gatewayInitData = abi.encodeWithSelector(
            CBJGateway.initialize.selector,
            mockUSDCAddress,
            usdmProxy,
            deployer,
            uint256(0),
            uint256(0)
        );
        address gatewayProxy = address(
            new TransparentUpgradeableProxy(
                gatewayImpl,
                proxyAdmin,
                gatewayInitData
            )
        );
        saveContract("CBJGateway", gatewayProxy);
        console2.log("CBJGateway deployed at:", gatewayProxy);

        // =============================================
        // Step 6: 权限配置
        // =============================================
        CBJToken usdm = CBJToken(usdmProxy);
        CBJToken aapl = CBJToken(aaplProxy);
        CBJToken tsla = CBJToken(tslaProxy);
        OrderContract order = OrderContract(orderProxy);

        // 6a. OrderContract.markExecuted 现在内部直接 mint+burn(买单烧USDM铸股票代币,
        //     卖单烧股票代币铸USDM),必须在 USDM 和每只股票代币上同时拿到两个角色
        usdm.grantRole(MINT_ROLE, orderProxy);
        usdm.grantRole(BURN_ROLE, orderProxy);
        aapl.grantRole(MINT_ROLE, orderProxy);
        aapl.grantRole(BURN_ROLE, orderProxy);
        tsla.grantRole(MINT_ROLE, orderProxy);
        tsla.grantRole(BURN_ROLE, orderProxy);
        console2.log(
            "OrderContract: granted MINT_ROLE + BURN_ROLE on USDM / AAPL.cbj / TSLA.cbj"
        );

        // 6b. CBJGateway 需要在 USDM 上有 mint(充值时)和 burn(提现时)的权限
        usdm.grantRole(MINT_ROLE, gatewayProxy);
        usdm.grantRole(BURN_ROLE, gatewayProxy);
        console2.log("CBJGateway: granted MINT_ROLE + BURN_ROLE on USDM");

        // 6c. 把股票代币登记进 OrderContract
        order.setSymbolToken("AAPL", aaplProxy);
        order.setSymbolToken("TSLA", tslaProxy);
        console2.log("Registered AAPL / TSLA symbol tokens on OrderContract");

        // 6d. 把 Blocklist/SanctionsList 接到每只代币上 —— 名单目前是空的,不影响任何人使用,
        //     后面要拉黑谁,直接对 CBJBlocklist/CBJSanctionsList 操作,不需要重新部署任何东西
        compliance.setBlocklist(usdmProxy, CBJBlocklist(blocklistProxy));
        compliance.setSanctionsList(
            usdmProxy,
            CBJSanctionsList(sanctionsListProxy)
        );
        compliance.setBlocklist(aaplProxy, CBJBlocklist(blocklistProxy));
        compliance.setSanctionsList(
            aaplProxy,
            CBJSanctionsList(sanctionsListProxy)
        );
        compliance.setBlocklist(tslaProxy, CBJBlocklist(blocklistProxy));
        compliance.setSanctionsList(
            tslaProxy,
            CBJSanctionsList(sanctionsListProxy)
        );
        console2.log(
            "Wired Blocklist/SanctionsList to USDM / AAPL.cbj / TSLA.cbj (lists are empty by default)"
        );

        console2.log("Deployment completed successfully");
    }
}
