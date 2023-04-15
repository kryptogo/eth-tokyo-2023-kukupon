import { HardhatUserConfig, task } from "hardhat/config";
import "@nomiclabs/hardhat-etherscan";
import "@typechain/hardhat";
import "@nomicfoundation/hardhat-toolbox";
require("hardhat-abi-exporter");
require("dotenv").config();
require("@nomiclabs/hardhat-ethers");

const { MUMBAI_API_KEY, POLYGON_API_KEY, PRIVATE_KEY, PRIVATE_KEY_LOCAL } =
  process.env;

module.exports = {
  solidity: {
    version: "0.8.18",
    settings: {
      optimizer: {
        enabled: true,
        runs: 1000,
      },
    },
  },
  networks: {
    localhost: {
      url: "http://127.0.0.1:8545",
      // accounts: [PRIVATE_KEY_LOCAL],
    },
    mumbai: {
      url: `https://polygon-mumbai.infura.io/v3/${MUMBAI_API_KEY}`,
      accounts: [PRIVATE_KEY],
      gasPrice: 8000000000,
    },
    polygon: {
      url: `https://polygon-mainnet.infura.io/v3/${POLYGON_API_KEY}`,
      accounts: [PRIVATE_KEY],
      gasPrice: 80000000000,
    },
  },
  etherscan: {
    apiKey: {
      polygonMumbai: process.env.POLYGONSCAN_API_KEY,
      polygon: process.env.POLYGONSCAN_API_KEY,
    },
  },
  typechain: {
    outDir: "src/types",
    target: "ethers-v5",
  },

  allowUnlimitedContractSize: true,

  abiExporter: {
    path: "./data/abi",
    runOnCompile: true,
    clear: true,
    flat: true,
    // only: [":ERC20$"],
    spacing: 2,
    format: "json",
  },
};
