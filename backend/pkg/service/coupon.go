package service

import (
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/kryptogo/eth-tokyo-2023-kukuponpon/pkg/db"
)

func GetCoupons(camapignId string) []string {
	number := 6 // FIXME: get from db
	coupons, privateKeys := GenerateWallets(number)

	// Generate 4337 wallets
	newWallets := GenerateWallets4337(privateKeys)

	// addToWhitelist to paymaster
	for _, newWallet := range newWallets {
		camapign := db.GetCampaign(camapignId)
		paymasterAddress := camapign.PaymasterAddress
		sponsorGas := camapign.SponsorGas
		UpdatePayMaster(paymasterAddress, newWallet, sponsorGas)
	}
	return coupons
}

func GenerateWallets(number int) ([]string, []*ecdsa.PrivateKey) {
	coupons := []string{}
	privateKeys := []*ecdsa.PrivateKey{}
	for i := 0; i < number; i++ {
		randomString := getRandomString()
		coupons = append(coupons, randomString)

		// 2. Hash the random string using SHA-256 to create a private key
		hash := sha256.Sum256([]byte(randomString))
		privateKey := crypto.ToECDSAUnsafe(hash[:])
		privateKeys = append(privateKeys, privateKey)
	}
	return coupons, privateKeys
}

func GenerateWallets4337(privateKeys []*ecdsa.PrivateKey) []string {
	return []string{"1", "TODO:"}
}

func UpdatePayMaster(paymasterAddress, newAddress string, sponsorGas big.Int) {
	hostWalletAddress := os.Getenv("SIGNING_WALLET_ADDRESS")
	sesssion, err := PrepareTx(hostWalletAddress, newAddress, paymasterAddress)
	if err != nil {
		panic(err)
	}

	newWalletAddress := common.HexToAddress(newAddress)
	_, err = sesssion.AddToWhitelist(newWalletAddress, &sponsorGas)
	if err != nil {
		panic(err)
	}
}

func getRandomString() string {
	// 產生32 bytes的隨機字串
	randomBytes := make([]byte, 32)
	_, err := rand.Read(randomBytes)
	if err != nil {
		panic(err)
	}

	// 將bytes轉換為base64編碼的字串
	randomString := base64.URLEncoding.EncodeToString(randomBytes)
	return randomString
}
