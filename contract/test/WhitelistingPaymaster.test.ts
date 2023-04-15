import "@nomiclabs/hardhat-ethers";
import { expect } from "chai";
import { ethers } from "hardhat";
import { EntryPoint, WhitelistingPaymaster, SimpleAccount } from "../src/types";
import { UserOperationStruct } from "../src/types/contracts/account/SimpleAccount";

describe("WhitelistingPaymaster", function () {
  let whitelistingPaymaster: WhitelistingPaymaster,
    entryPoint: EntryPoint,
    simpleAccount: SimpleAccount;
  let owner: any, addr1: any;

  beforeEach(async () => {
    [owner, addr1] = await ethers.getSigners();

    // Deploy the EntryPoint contract
    const EntryPointFac = await ethers.getContractFactory("EntryPoint", owner);
    entryPoint = await EntryPointFac.deploy();
    await entryPoint.deployed();

    // Deploy the WhitelistingPaymaster contract
    const WhitelistingPaymasterFac = await ethers.getContractFactory(
      "WhitelistingPaymaster",
      owner
    );
    whitelistingPaymaster = await WhitelistingPaymasterFac.deploy(
      entryPoint.address
    );
    await whitelistingPaymaster.deployed();

    // Deploy the simple account factory
    const SimpleAccountFactory = await ethers.getContractFactory(
      "SimpleAccount",
      addr1
    );
    simpleAccount = (await SimpleAccountFactory.deploy(
      entryPoint.address
    )) as SimpleAccount;
    await simpleAccount.deployed();
  });

  it("should successfully deposit 1 ETH, add SimpleAccount contract address to the whitelist, and pay 0.1 ETH", async () => {
    // Send 1 ETH to the deposit() method of the paymaster contract
    await whitelistingPaymaster.deposit({
      value: ethers.utils.parseEther("1"),
    });

    // Check that the entrypoint balance has increased by 1 ETH
    const entrypointBalance = (
      await entryPoint.getDepositInfo(whitelistingPaymaster.address)
    ).deposit;
    expect(entrypointBalance).to.equal(ethers.utils.parseEther("1"));

    // Add the SimpleAccount contract address to the whitelist and pay amount of 0.1 ETH
    await whitelistingPaymaster.addToWhitelist(
      simpleAccount.address,
      ethers.utils.parseEther("0.1")
    );

    // Check that the SimpleAccount contract address has been added to the whitelist and has the correct sponsored gas
    const userDetails = await whitelistingPaymaster.userDetails(
      simpleAccount.address
    );
    expect(userDetails.isWhitelisted).to.be.true;
    expect(userDetails.remainingGas).to.equal(ethers.utils.parseEther("0.1"));
  });

  it("should execute user ops sponsored by paymaster contract", async () => {
    // Send 1 ETH to the deposit() method of the paymaster contract
    await whitelistingPaymaster.deposit({
      value: ethers.utils.parseEther("1"),
    });

    // Add the SimpleAccount contract address to the whitelist and pay amount of 0.1 ETH
    await whitelistingPaymaster.addToWhitelist(
      simpleAccount.address,
      ethers.utils.parseEther("0.1")
    );

    // construct user ops to send to entrypoint
    const addr2 = (await ethers.getSigners())[2];

    const userOperation: UserOperationStruct = {
      sender: simpleAccount.address,
      nonce: 0,
      initCode: "0x",
      callData: simpleAccount.interface.encodeFunctionData("transfer", [addr2.address, ethers.utils.parseEther("0.2")]),
      callGasLimit: 100000,
      verificationGasLimit: 0,
      preVerificationGas: 0,
      maxFeePerGas: ethers.utils.parseUnits("100", "gwei"),
      maxPriorityFeePerGas: ethers.utils.parseUnits("2", "gwei"),
      paymasterAndData: ethers.utils.defaultAbiCoder.encode(["address"], [whitelistingPaymaster.address]),
      signature: "0x",
    };

    const ops = [userOperation];
    const beneficiary = whitelistingPaymaster.address;

    await expect(entryPoint.connect(addr1).handleOps(ops, beneficiary))
      .to.emit(simpleAccount, "Transfer")
      .withArgs(addr1.address, addr2.address, ethers.utils.parseEther("0.2"));
  });
});
