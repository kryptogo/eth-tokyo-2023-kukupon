import { ethers } from "hardhat";
require("dotenv").config();

// Usage: npx hardhat run scripts/deposit-paymaster.ts --network yourNetwork
const sponsoredGas = "0.01";

async function main() {
  const WhitelistingPaymaster = await ethers.getContractFactory(
    "WhitelistingPaymaster"
  );
  const paymaster = await WhitelistingPaymaster.attach(
    process.env.PAYMASTER_CONTRACT_ADDRESS!
  );
  const tx = await paymaster.addToWhitelist(
    "0x8c491bDA59470aB7E229DB69A6a4E184010C55eb",
    ethers.utils.parseEther(sponsoredGas)
  );
  console.log("tx hash", tx.hash);
}

main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
