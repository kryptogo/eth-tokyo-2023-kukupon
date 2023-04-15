// SPDX-License-Identifier: GPL-3.0
pragma solidity ^0.8.12;

/* solhint-disable reason-string */

import "../core/BasePaymaster.sol";
import "@openzeppelin/contracts/access/Ownable.sol";

/**
 * A paymaster which verifies a given sender address is in whitelist map and
 * has sufficient gas quota
 */
contract WhitelistingPaymaster is BasePaymaster {

    //calculated cost of the postOp
    uint256 constant public COST_OF_POST = 15000;

    using UserOperationLib for UserOperation;

    struct UserDetails {
        bool isWhitelisted;
        uint256 remainingGas;
    }
    mapping(address => UserDetails) public userDetails;

    event AddressAddedToWhitelist(address indexed addr, uint256 sponsoredGas);

    constructor(IEntryPoint _entryPoint) BasePaymaster(_entryPoint) {
    }

    function addToWhitelist(address addr, uint256 sponsoredGas) external onlyOwner {
        require(addr != address(0), "Invalid address");
        userDetails[addr] = UserDetails({
          isWhitelisted: true, 
          remainingGas: sponsoredGas
        });
        emit AddressAddedToWhitelist(addr, sponsoredGas);
    }

    function validatePaymasterUserOp(UserOperation calldata userOp, bytes32 /*userOpHash*/, uint256 maxCost)
        external view override returns (bytes memory context, uint256 deadline)
    {
        UserDetails memory user = userDetails[userOp.sender];
        if (user.isWhitelisted) {
            if (user.remainingGas >= maxCost) {
                return ("", 0);
            }
        }
        return ("", 1);
    }


    /**
     * actual charge of user.
     * this method will be called just after the user's TX with mode==OpSucceeded|OpReverted (account pays in both cases)
     * BUT: if the user changed its balance in a way that will cause  postOp to revert, then it gets called again, after reverting
     * the user's TX , back to the state it was before the transaction started (before the validatePaymasterUserOp),
     * and the transaction should succeed there.
     */
    function _postOp(PostOpMode mode, bytes calldata context, uint256 actualGasCost) internal override {
        //we don't really care about the mode, we just pay the gas with the user's tokens.
        (mode);
        (address sender, uint256 remainingGas) = abi.decode(context, (address, uint256));
        uint256 totalCost = actualGasCost + COST_OF_POST;
        userDetails[sender] = UserDetails({
          isWhitelisted: true, 
          remainingGas: remainingGas - totalCost
        });
    }

}
