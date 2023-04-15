import "@nomiclabs/hardhat-ethers";
import { expect } from "chai";
import { ethers } from "hardhat";
import { Wallet } from "ethers";
import { EntryPoint, VerifyingPaymaster } from "../src/types";

describe("VerifyingPaymaster", function () {
  let VerifyingPaymasterFac,
    verifyingPaymaster: VerifyingPaymaster,
    EntryPointFac,
    entryPoint: EntryPoint,
    verifyingSigner: Wallet;

  beforeEach(async () => {
    // Deploy the EntryPoint contract (mock)
    EntryPointFac = await ethers.getContractFactory("EntryPoint");
    entryPoint = await EntryPointFac.deploy();
    await entryPoint.deployed();

    // Set up the verifyingSigner
    const privateKey = process.env.PRIVATE_KEY!;
    verifyingSigner = new ethers.Wallet(privateKey);

    // Deploy the VerifyingPaymaster contract
    VerifyingPaymasterFac = await ethers.getContractFactory(
      "VerifyingPaymaster"
    );
    verifyingPaymaster = await VerifyingPaymasterFac.deploy(
      entryPoint.address,
      verifyingSigner.address
    );
    await verifyingPaymaster.deployed();
  });

  it("Should set the correct verifyingSigner address", async function () {
    expect(await verifyingPaymaster.verifyingSigner()).to.equal(
      verifyingSigner.address
    );
  });
});
