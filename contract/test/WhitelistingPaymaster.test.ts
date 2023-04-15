import "@nomiclabs/hardhat-ethers";
import { expect } from "chai";
import { ethers } from "hardhat";
import { EntryPoint, WhitelistingPaymaster, SimpleAccount } from "../src/types";
import { UserOperationStruct } from "../src/types/contracts/account/SimpleAccount";

// Add a helper function to calculate the user operation hash
// Define the getUserOpHash function
function getUserOpHash(
  op: UserOperationStruct,
  entryPoint: string,
  chainId: number
): string {
  const userOpHash = (0, ethers.utils.keccak256)(packUserOp(op, true));
  const enc = ethers.utils.defaultAbiCoder.encode(
    ["bytes32", "address", "uint256"],
    [userOpHash, entryPoint, chainId]
  );
  return (0, ethers.utils.keccak256)(enc);
}

function encode(typevalues: any, forSignature: any) {
  const types = typevalues.map((typevalue: any) =>
    typevalue.type === "bytes" && forSignature ? "bytes32" : typevalue.type
  );
  const values = typevalues.map((typevalue: any) =>
    typevalue.type === "bytes" && forSignature
      ? (0, ethers.utils.keccak256)(typevalue.val)
      : typevalue.val
  );
  return ethers.utils.defaultAbiCoder.encode(types, values);
}

function packUserOp(op: any, forSignature = true) {
  if (forSignature) {
    // lighter signature scheme (must match UserOperation#pack): do encode a zero-length signature, but strip afterwards the appended zero-length value
    const userOpType = {
      components: [
        {
          type: "address",
          name: "sender",
        },
        {
          type: "uint256",
          name: "nonce",
        },
        {
          type: "bytes",
          name: "initCode",
        },
        {
          type: "bytes",
          name: "callData",
        },
        {
          type: "uint256",
          name: "callGasLimit",
        },
        {
          type: "uint256",
          name: "verificationGasLimit",
        },
        {
          type: "uint256",
          name: "preVerificationGas",
        },
        {
          type: "uint256",
          name: "maxFeePerGas",
        },
        {
          type: "uint256",
          name: "maxPriorityFeePerGas",
        },
        {
          type: "bytes",
          name: "paymasterAndData",
        },
        {
          type: "bytes",
          name: "signature",
        },
      ],
      name: "userOp",
      type: "tuple",
    };
    // console.log('hard-coded userOpType', userOpType)
    // console.log('from ABI userOpType', UserOpType)
    let encoded = ethers.utils.defaultAbiCoder.encode(
      [userOpType],
      [Object.assign(Object.assign({}, op), { signature: "0x" })]
    );
    // remove leading word (total length) and trailing word (zero-length signature)
    encoded = "0x" + encoded.slice(66, encoded.length - 64);
    return encoded;
  }
  const typevalues = UserOpType.components.map((c) => ({
    type: c.type,
    val: op[c.name],
  }));
  return encode(typevalues, forSignature);
}

describe("WhitelistingPaymaster", function () {
  let whitelistingPaymaster: WhitelistingPaymaster,
    entryPoint: EntryPoint,
    simpleAccount: SimpleAccount;
  let owner: any, addr1: any, addr2: any;

  beforeEach(async () => {
    [owner, addr1, addr2] = await ethers.getSigners();

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

    const userOperation: UserOperationStruct = {
      sender: simpleAccount.address,
      nonce: 0,
      initCode: "0x",
      callData: simpleAccount.interface.encodeFunctionData("execute", [
        addr2.address,
        ethers.utils.parseEther("0.2"),
        "0x",
      ]),
      callGasLimit: 100000,
      verificationGasLimit: 0,
      preVerificationGas: 0,
      maxFeePerGas: ethers.utils.parseUnits("100", "gwei"),
      maxPriorityFeePerGas: ethers.utils.parseUnits("2", "gwei"),
      paymasterAndData: ethers.utils.defaultAbiCoder.encode(
        ["address"],
        [whitelistingPaymaster.address]
      ),
      signature: "0x",
    };
    // Get the user operation hash
    const userOpHash = getUserOpHash(
      userOperation,
      entryPoint.address,
      (await ethers.provider.getNetwork()).chainId
    );
    // Sign the user operation hash with addr2's private key
    const signature = await addr2.signMessage(
      ethers.utils.arrayify(userOpHash)
    );
    // Update the signature field in the user operation
    userOperation.signature = signature;

    const ops = [userOperation];
    const beneficiary = whitelistingPaymaster.address;

    // await expect()
    //   .to.emit(simpleAccount, "Transfer")
    //   .withArgs(addr1.address, addr2.address, ethers.utils.parseEther("0.2"));
    await entryPoint.connect(addr1).handleOps(ops, beneficiary);

    // Check that the SimpleAccount contract address has used some sponsored gas
    const userDetails = await whitelistingPaymaster.userDetails(
      simpleAccount.address
    );
    expect(userDetails.isWhitelisted).to.be.true;
    expect(userDetails.remainingGas).to.lessThan(
      ethers.utils.parseEther("0.1")
    );
  });
});
