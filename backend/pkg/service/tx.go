package service

import (
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"
	"math/big"
	"os"

	"github.com/btcsuite/btcd/btcec"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	alchemyapi "github.com/kryptogo/eth-tokyo-2023-kukuponpon/pkg/alchemy-api"
)

func composeTxData(from, to string) (*bind.TransactOpts, error) {
	value := big.NewInt(0) // nft transfer value is 0
	fmt.Println("composeTxData before GetNextNonce")

	fmt.Println("composeTxData after GetNextNonce")
	var gasPrice *big.Int
	if os.Getenv("CHAIN_ID") == "polygon" {
		gasPrice = big.NewInt(200000000000) //matic
	} else {
		gasPrice = big.NewInt(20000000000) // mumbai
	}
	nonce := new(big.Int).Set(alchemyapi.Nonce)
	opts := &bind.TransactOpts{
		From:     common.HexToAddress(from),
		Value:    value,
		GasPrice: gasPrice,
		GasLimit: 200000,
		Signer: func(address common.Address, tx *types.Transaction) (*types.Transaction, error) {
			return sign(tx)
		},
		Nonce: nonce,
	}
	fmt.Println("[composeTxData], nonce: ", nonce)
	alchemyapi.Nonce.Add(alchemyapi.Nonce, big.NewInt(1))
	fmt.Println("[composeTxData] after TransactOpts", opts)
	return opts, nil
}

func PrepareTxPayment(from, contract string) (*WhitelistingPaymaster, *bind.TransactOpts, error) {
	url := alchemyapi.GetAPIURL()
	conn, err := ethclient.Dial(url)
	if err != nil {
		return nil, nil, err
	}

	contractAddress := common.HexToAddress(contract)
	fmt.Println("[PrepareTxPayment] contractAddress: ", contractAddress.Hex())
	a, err := NewWhitelistingPaymaster(contractAddress, conn)
	if err != nil {
		return nil, nil, err
	}

	txOpts, err := composeTxData(from, contract)
	if err != nil {
		return nil, nil, err
	}

	return a, txOpts, nil
}

func PrepareTxAccountFactory(from, contract string) (*SimpleAccountFactory, *bind.TransactOpts, error) {
	url := alchemyapi.GetAPIURL()
	conn, err := ethclient.Dial(url)
	if err != nil {
		return nil, nil, err
	}
	fmt.Println("PrepareTxAccountFactory init")
	contractAddress := common.HexToAddress(contract)
	a, err := NewSimpleAccountFactory(contractAddress, conn)
	fmt.Println("PrepareTxAccountFactory NewSimpleAccountFactory")

	if err != nil {
		return nil, nil, err
	}
	fmt.Println("PrepareTxAccountFactory before composeTxData")
	// txOpts, err := composeTxData(from, contract)
	if err != nil {
		return nil, nil, err
	}
	fmt.Println("PrepareTxAccountFactory after composeTxData")

	return a, nil, nil
}

func sign(tx *types.Transaction) (*types.Transaction, error) {
	privateKeyStr := os.Getenv("SIGNING_PRIVATE_KEY")
	privateKey, err := parsePrivateKey(string(privateKeyStr))
	if err != nil {
		return nil, err
	}
	var chainIDBigInt *big.Int
	if os.Getenv("CHAIN_ID") == "polygon" {
		chainIDBigInt = big.NewInt(137)
	} else {
		chainIDBigInt = big.NewInt(80001)
	}

	signer := types.NewLondonSigner(chainIDBigInt)
	return types.SignTx(tx, signer, privateKey)
}

// hexToBigInt convert hex string to big.Int
func hexToBigInt(s string) *big.Int {
	if len(s) > 2 && s[:2] == "0x" {
		s = s[2:]
	}
	n := new(big.Int)
	n.SetString(s, 16)
	return n
}

func parsePrivateKey(privKeyHex string) (*ecdsa.PrivateKey, error) {
	curve := btcec.S256() // secp256k1 curve used by  go-ethereum

	privKeyBytes, err := hex.DecodeString(privKeyHex)
	if err != nil {
		return nil, err
	}

	privKey := new(ecdsa.PrivateKey)
	privKey.D = new(big.Int).SetBytes(privKeyBytes)
	privKey.PublicKey.Curve = curve
	privKey.PublicKey.X, privKey.PublicKey.Y = curve.ScalarBaseMult(privKeyBytes)

	return privKey, nil
}
