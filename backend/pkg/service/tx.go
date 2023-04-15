package service

import (
	"crypto/ecdsa"
	"encoding/hex"
	"math/big"
	"os"

	"github.com/btcsuite/btcd/btcec"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	alchemyapi "github.com/kryptogo/eth-tokyo-2023-kukuponpon/pkg/alchemy-api"
)

func composeTxData(from string) (*bind.TransactOpts, error) {
	value := big.NewInt(0) // nft transfer value is 0
	nonce, err := alchemyapi.GetNextNonce(from)
	if err != nil {
		return nil, err
	}
	opts := &bind.TransactOpts{
		From:     common.HexToAddress(from),
		Value:    value,
		GasPrice: big.NewInt(200000000000),
		GasLimit: 200000,
		Signer: func(address common.Address, tx *types.Transaction) (*types.Transaction, error) {
			return sign(tx)
		},
		Nonce: hexToBigInt(nonce),
	}
	return opts, nil
}

func PrepareTx(from, newAddress, paymasterAddress string) (*WhitelistingPaymasterSession, error) {
	url := alchemyapi.GetAPIURL()
	conn, err := ethclient.Dial(url)
	if err != nil {
		return nil, err
	}

	contractAddress := common.HexToAddress(paymasterAddress)
	a, err := NewWhitelistingPaymaster(contractAddress, conn)
	if err != nil {
		return nil, err
	}

	txOpts, err := composeTxData(from)
	if err != nil {
		return nil, err
	}

	session := WhitelistingPaymasterSession{
		Contract:     a,
		TransactOpts: *txOpts,
	}
	return &session, nil
}

func sign(tx *types.Transaction) (*types.Transaction, error) {
	privateKeyStr := os.Getenv("SIGNING_PRIVATE_KEY")
	privateKey, err := parsePrivateKey(string(privateKeyStr))
	if err != nil {
		return nil, err
	}

	chainIDBigInt := big.NewInt(137)
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
