# KuKuPon: Walletless Blockchain Platform

KuKuPon is a walletless, forward-looking technology platform that connects sponsors with influential individuals and experienced users. By leveraging EIP-4337 technology, we aim to cover gas fees for new users and provide a zero-transaction-fee wallet in the form of coupon codes, thus lowering the barriers to entry for new blockchain users. We have also integrated airstack for on-chain data validation and Worldcoin's anonymous authentication technology to prevent duplicate claims.

## Requirements

Go v1.17

## Third-Party Packages

1. [ERC-4337 Account Abstract](https://eips.ethereum.org/EIPS/eip-4337)
2. [Worldcoin](https://worldcoin.org/)
3. [airstack](https://www.airstack.xyz/)

## Getting Started

1. Clone the repository to your local machine:

```bash=
git clone https://github.com/yourusername/kukupon.git
```

2. Change to the project's directory:

```bash=
cd kukupon
```

3. Install dependencies:

```bash=
go mod download
```

4. Copy the example configuration file `config/env.sh.example` to `config/env.sh`:

```bash=
cp config/env.sh.example config/env.sh
```

5. Modify the configuration file `config/env.sh` with the appropriate parameters. Refer to the comments in the file for more information about the parameters.
6. Run the project:

```bash=
make run
```

## Configuration Parameters

Please refer to `config/env.sh.example` for a list of configuration parameters.

## Contributing

We welcome contributions to the KuKuPon project. Please submit a pull request or create an issue to discuss your ideas.

## License

This project is licensed under the <u>MIT License</u>
