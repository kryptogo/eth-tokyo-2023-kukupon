import { ethers } from "hardhat";
require("dotenv").config();

// Usage: npx hardhat run scripts/deposit-paymaster.ts --network yourNetwork
const etherToStake = "0.01";

async function main() {
  const WhitelistingPaymaster = await ethers.getContractFactory(
    "WhitelistingPaymaster"
  );
  const paymaster = await WhitelistingPaymaster.attach(
    process.env.PAYMASTER_CONTRACT_ADDRESS!
  );
  const tx = await paymaster.addStake(1000, {
    value: ethers.utils.parseEther(etherToStake),
  });
  console.log("tx hash", tx.hash);
}

main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
