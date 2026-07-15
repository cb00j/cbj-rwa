// SPDX-License-Identifier: MIT
pragma solidity ^0.8.26;

import {ICBJTokenFactory} from "src/token/interfaces/ICBJTokenFactory.sol";
import {ICBJRegistrar} from "src/token/interfaces/ICBJRegistrar.sol";
import {CBJToken} from "src/token/CBJToken.sol";
import {AccessControlEnumerable} from "@openzeppelin/contracts/access/extensions/AccessControlEnumerable.sol";
import {Initializable} from "@openzeppelin/contracts/proxy/utils/Initializable.sol";
import {ReentrancyGuard} from "@openzeppelin/contracts/utils/ReentrancyGuard.sol";
import {Pausable} from "@openzeppelin/contracts/utils/Pausable.sol";
import {UpgradeableBeacon} from "@openzeppelin/contracts/proxy/beacon/UpgradeableBeacon.sol";
import {BeaconProxy} from "@openzeppelin/contracts/proxy/beacon/BeaconProxy.sol";

/**
 * @title  CBJTokenFactory
 * @notice This contract serves as a factory for deploying and configuring CBJ tokens with
 *         built-in compliance and pause management.
 *
 *         This contract allows for:
 *         - Deploying new CBJ tokens with preconfigured compliance and pause management
 *         - Registering the tokens for use via with bridge and token manager registrars
 *         - Isolated token deployments without registrar integration
 *
 * @dev    The contract uses OpenZeppelin's AccessControl for role-based permissions:
 *         - DEFAULT_ADMIN_ROLE: Can grant/revoke other roles
 *         - DEPLOY_ROLE: Can deploy new tokens
 *         - CONFIGURE_ROLE: Can update compliance and registrar settings
 */
contract CBJTokenFactory is
    ICBJTokenFactory,
    AccessControlEnumerable,
    ReentrancyGuard,
    Pausable,
    Initializable
{
    // Role used for deploying new tokens
    bytes32 public constant DEPLOY_ROLE = keccak256("DEPLOY_ROLE");
    // Role used for configuring the factory
    bytes32 public constant CONFIGURE_ROLE = keccak256("CONFIGURE_ROLE");
    // Role used to pause the factory
    bytes32 public constant PAUSE_ROLE = keccak256("PAUSE_ROLE");
    // Role used to unpause the factory
    bytes32 public constant UNPAUSE_ROLE = keccak256("UNPAUSE_ROLE");

    // Address of the BEACON contract used for proxy deployments
    address public immutable BEACON;

    // Address of the CBJCompliance contract
    address public cbjCompliance;
    // Address of the CBJTokenPauseManager contract
    address public cbjTokenPauseManager;
    // Address of the token registrar contract
    ICBJRegistrar public tokenManagerRegistrar;

    // Indicates if a token with the same symbol already exists
    mapping(string => bool) public symbolExists;

    /**
     * @notice Constructor for implementation contract
     */
    constructor() {
        _disableInitializers();

        // Initialize BEACON (immutable variable)
        address cbjTokenImplementation = address(new CBJToken());
        // Use a temporary owner for the BEACON, will be transferred in initialize()
        UpgradeableBeacon _beacon = new UpgradeableBeacon(
            cbjTokenImplementation,
            address(this)
        );
        BEACON = address(_beacon);
    }

    /**
     * @notice Initialize function for proxy deployment
     * @param _guardian The address which will be granted admin and other roles
     * @param _cbjCompliance The address of the CBJCompliance contract
     * @param _tokenPauseManager The address of the token pause manager contract
     * @param _tokenManagerRegistrar The address of the token manager registrar contract,maybe empty
     */
    function initialize(
        address _guardian,
        address _cbjCompliance,
        address _tokenPauseManager,
        address _tokenManagerRegistrar
    ) external initializer {
        if (_cbjCompliance == address(0)) revert ComplianceCannotBeZero();
        if (_tokenPauseManager == address(0))
            revert TokenPauseManagerCannotBeZero();

        _grantRole(DEFAULT_ADMIN_ROLE, _guardian);
        _grantRole(DEPLOY_ROLE, _guardian);
        _grantRole(CONFIGURE_ROLE, _guardian);
        _grantRole(PAUSE_ROLE, _guardian);
        _grantRole(UNPAUSE_ROLE, _guardian);

        cbjCompliance = _cbjCompliance;
        cbjTokenPauseManager = _tokenPauseManager;
        tokenManagerRegistrar = ICBJRegistrar(_tokenManagerRegistrar);
    }

    /**
     * @notice Pauses the factory, disabling new deployments
     */
    function pause() external onlyRole(PAUSE_ROLE) {
        _pause();
    }

    /**
     * @notice Unpauses the factory, enabling new deployments
     */
    function unpause() external onlyRole(UNPAUSE_ROLE) {
        _unpause();
    }

    /**
     * @notice Deploys a new CBJ token and registers it with both bridge and token manager
     * @param  name       The name of the token
     * @param  symbol     The token symbol
     * @param  tokenAdmin The address that will receive admin rights on the token
     * @return            The address of the deployed token proxy
     */
    function deployAndRegisterToken(
        string calldata name,
        string calldata symbol,
        address tokenAdmin
    )
        public
        override
        nonReentrant
        onlyRole(DEPLOY_ROLE)
        whenNotPaused
        returns (address)
    {
        if (tokenAdmin == address(0))
            revert TokenManagerRegistrarCannotBeZero();

        CBJToken token = CBJToken(_deployCBJToken(name, symbol));
        token.grantRole(DEFAULT_ADMIN_ROLE, address(tokenManagerRegistrar));
        tokenManagerRegistrar.register(address(token));
        token.revokeRole(DEFAULT_ADMIN_ROLE, address(tokenManagerRegistrar));

        token.grantRole(DEFAULT_ADMIN_ROLE, tokenAdmin);
        token.renounceRole(DEFAULT_ADMIN_ROLE, address(this));
        return address(token);
    }

    /**
     * @notice Deploys a new CBJ token without registering it anywhere
     * @param  name       The name of the token
     * @param  symbol     The token symbol
     * @param  tokenAdmin The address that will receive admin rights on the token
     * @return            The address of the deployed token proxy
     */
    function deployCBJTokenIsolated(
        string calldata name,
        string calldata symbol,
        address tokenAdmin
    )
        public
        override
        nonReentrant
        onlyRole(DEPLOY_ROLE)
        whenNotPaused
        returns (address)
    {
        CBJToken token = CBJToken(_deployCBJToken(name, symbol));
        token.grantRole(DEFAULT_ADMIN_ROLE, tokenAdmin);
        token.renounceRole(DEFAULT_ADMIN_ROLE, address(this));
        return address(token);
    }

    /**
     * @notice Updates the compliance contract address
     * @param  _cbjCompliance The new compliance contract address
     */
    function setCompliance(
        address _cbjCompliance
    ) external onlyRole(CONFIGURE_ROLE) {
        if (_cbjCompliance == address(0)) revert ComplianceCannotBeZero();
        emit NewComplianceSet(cbjCompliance, _cbjCompliance);
        cbjCompliance = _cbjCompliance;
    }

    /**
     * @notice Updates the token pause manager address
     * @param  _tokenPauseManager The new token pause manager address
     */
    function setTokenPauseManager(
        address _tokenPauseManager
    ) external onlyRole(CONFIGURE_ROLE) {
        if (_tokenPauseManager == address(0))
            revert TokenPauseManagerCannotBeZero();
        emit NewTokenPauseManagerSet(cbjTokenPauseManager, _tokenPauseManager);
        cbjTokenPauseManager = _tokenPauseManager;
    }

    /**
     * @notice Updates the token manager registrar address
     * @param  _tokenManagerRegistrar The new token manager registrar address
     */
    function setTokenManagerRegistrar(
        address _tokenManagerRegistrar
    ) external onlyRole(CONFIGURE_ROLE) {
        if (_tokenManagerRegistrar == address(0)) {
            revert TokenManagerRegistrarCannotBeZero();
        }
        emit NewTokenManagerRegistrarSet(
            address(tokenManagerRegistrar),
            _tokenManagerRegistrar
        );
        tokenManagerRegistrar = ICBJRegistrar(_tokenManagerRegistrar);
    }

    /**
     * @notice Clears a symbol in the edge case where we need to deploy a token
     *         with the same symbol as a previously deployed token
     * @param symbol The symbol to clear
     */
    function clearSymbol(
        string calldata symbol
    ) external onlyRole(CONFIGURE_ROLE) {
        symbolExists[symbol] = false;
        emit SymbolSet(symbol, false);
    }

    /**
     * @notice Internal function to deploy a new CBJ token
     * @param  name   The name of the token
     * @param  symbol The token symbol
     * @return        The address of the deployed token proxy
     */
    function _deployCBJToken(
        string calldata name,
        string calldata symbol
    ) internal returns (address) {
        if (symbolExists[symbol]) revert SymbolAlreadyExists();
        BeaconProxy cbjTokenProxy = new BeaconProxy(BEACON, "");
        CBJToken cbjTokenProxied = CBJToken(address(cbjTokenProxy));
        cbjTokenProxied.initialize(
            name,
            symbol,
            cbjCompliance,
            cbjTokenPauseManager
        );
        symbolExists[symbol] = true;
        emit SymbolSet(symbol, true);

        emit NewCBJTokenDeployed(
            address(cbjTokenProxied),
            BEACON,
            name,
            symbol,
            cbjCompliance,
            cbjTokenPauseManager
        );

        return address(cbjTokenProxied);
    }
}
