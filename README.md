# Kukupon
> Where Gift Cards Meet Blockchain Gaming

![badmath](https://img.shields.io/github/languages/top/lernantino/badmath)

![img](img/kukupon.png)

## Introduction
Kukupon is a revolutionary platform that combines the best of gift cards and blockchain technology to provide a seamless and secure way for users to enter the world of GameFi. 

It is a more efficient and flexible solution that can reduce operational costs while providing better user experiences and more effective promotion tracking.

For more information about our design philosophy, please refer to our [white paper](whitepaper.md).

## Features
- Supports a guest mode for new players
- Provides entry NFTs
- Combines the advantages of Faucet and POAP with GameFi onboarding
- Wallet-less solution ensures the highest security without any burden
- Allows seamless gameplay across multiple devices
- Utilizes various key technologies to provide a secure and flexible gaming experience

## How to Use
1. Obtain a voucher code from a participating game developer
2. Input the voucher code in the designated field to access the account and start playing
3. Use different devices with the same voucher code for seamless gameplay across multiple devices

## Use Cases
- Social media promotion
- Influencer marketing
- Community building
- Rewarding loyal players
- Providing in-game items, currency, or other rewards with gift cards
- Providing discounts for purchases in the game
- Redeeming rewards or participating in special item or event giveaways
- Attracting new players

## Prerequisites

Before using this repo, please make sure you have the following installed:

### Frontend Prerequisites

Before running the frontend, make sure you have the following installed:

- [Node.js](https://nodejs.org/) v14 or higher

#### Installing

To install the frontend, follow these steps:

1. Clone this repository
2. `cd` into the `frontend` directory
3. Run `npm install` to install the necessary dependencies
4. Run `npm start` to start the development server
5. Open your browser and navigate to `http://localhost:3000`

### Contract Prerequisites

Before deploying the smart contract, make sure you have the following installed:

- [Hardhat](https://hardhat.org/) v2.0 or higher
- [Solidity](https://soliditylang.org/) v0.8.0 or higher

#### Installing

To install the contract, follow these steps:

1. Clone this repository
2. `cd` into the `contract` directory
3. Run `npm install` to install the necessary dependencies
4. Run `npx hardhat compile` to compile the contract code
5. Run `npx hardhat test` to run the contract tests
6. Run `npx hardhat run scripts/deploy.js --network <network>` to deploy the contract to the specified network

### Backend Prerequisites

Before running the backend, make sure you have the following installed:

- [Go](https://golang.org/) v1.16 or higher

#### Installing

To install the backend, follow these steps:

1. Clone this repository
2. `cd` into the `backend` directory
3. Run `go mod download` to download the necessary dependencies
4. Run `go run main.go` to start the backend server
5. The server will be available at `http://localhost:3001`

## Contributors (sorted in alphabetical order)
- Alan: Backend Engineer
- Dorara: Full-Stack Engineer
- Harry: Architect
- [Kordan](https://github.com/hitripod): Product Manager
- Kuan: Frontend Engineer

### How to Contribute

Please refer to [Contributor Covenant](https://www.contributor-covenant.org/).

## Demo

https://eth-tokyo-2023-kukupon.vercel.app