import { ethers } from "hardhat";
require("dotenv").config();

async function main() {
  const WhitelistingPaymasterFac = await ethers.getContractFactory(
    "WhitelistingPaymaster"
  );
  const paymaster = await WhitelistingPaymasterFac.deploy(
    process.env.ENTRYPOINT_CONTRACT_ADDRESS!
  );
  await paymaster.deployed();
  console.log(paymaster.address);
}

// We recommend this pattern to be able to use async/await everywhere
// and properly handle errors.
main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
