package service

import (
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"math/big"
	"os"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/kryptogo/eth-tokyo-2023-kukupon/pkg/db"
)

func GetCoupons(camapignId string) []string {
	camapign := db.GetCampaign(camapignId)
	number := camapign.CouponAmount
	coupons, newWallets := GenerateWallets(number)

	// Generate 4337 wallets
	newWallets4337 := GenerateWallets4337(newWallets)
	fmt.Println("[GetCoupons] newWallets4337: ", newWallets4337)
	// addToWhitelist to paymaster
	for _, newWallet := range newWallets4337 {
		paymasterAddress := os.Getenv("PAYMASTER_ADDRESS")
		sponsorGas := camapign.SponsorGas
		UpdatePayMaster(paymasterAddress, newWallet, sponsorGas)
		time.Sleep(50 * time.Millisecond)
	}
	return coupons
}

func GenerateWallets(number int) ([]string, []string) {
	coupons := []string{}
	wallets := []string{}
	for i := 0; i < number; i++ {
		randomString := getRandomString()
		coupons = append(coupons, randomString)

		// 2. Hash the random string using SHA-256 to create a private key
		hash := sha256.Sum256([]byte(randomString))
		privateKey := crypto.ToECDSAUnsafe(hash[:])

		// 3. Derive the public key from the private key
		publicKey := privateKey.Public().(*ecdsa.PublicKey)

		// 4. Get 4337 Address from privateKey
		ethereumAddress := crypto.PubkeyToAddress(*publicKey)
		wallets = append(wallets, common.BytesToAddress(ethereumAddress.Bytes()).Hex())
	}
	return coupons, wallets
}

func GenerateWallets4337(wallets []string) []string {
	hostWalletAddress := os.Getenv("SIGNING_WALLET_ADDRESS")
	SimpleAccountFactoryAddress := os.Getenv("SIMPLE_ACCOUNT_FACTORY_ADDRESS")
	a, _, err := PrepareTxAccountFactory(hostWalletAddress, SimpleAccountFactoryAddress)
	if err != nil {
		panic(err)
	}

	session := SimpleAccountFactorySession{
		Contract: a,
		// TransactOpts: *txOpts,
	}
	salt := big.NewInt(0)
	wallets4337 := []string{}
	for _, wallet := range wallets {
		walletAddress := common.HexToAddress(wallet)
		wallet4337, err := session.GetAddress(walletAddress, salt)
		if err != nil {
			panic(err)
		}
		wallets4337 = append(wallets4337, wallet4337.Hex())
	}
	return wallets4337
}

func UpdatePayMaster(paymasterAddress, newAddress string, sponsorGas big.Int) {
	fmt.Println("[UpdatePayMaster] begin")
	hostWalletAddress := os.Getenv("SIGNING_WALLET_ADDRESS")
	a, txOpts, err := PrepareTxPayment(hostWalletAddress, paymasterAddress)
	if err != nil {
		fmt.Println("PrepareTxPayment error: ", err)
		panic(err)
	}
	fmt.Println("[UpdatePayMaster] after PrepareTxPayment")

	session := WhitelistingPaymasterSession{
		Contract:     a,
		TransactOpts: *txOpts,
	}
	fmt.Println("[UpdatePayMaster] WhitelistingPaymasterSession Contract", a)
	fmt.Println("[UpdatePayMaster] WhitelistingPaymasterSession TransactOpts", *txOpts)

	fmt.Println("[UpdatePayMaster] after session")

	if err != nil {
		panic(err)
	}

	newWalletAddress := common.HexToAddress(newAddress)
	fmt.Println("[UpdatePayMaster] newWalletAddress", newWalletAddress)
	tx, err := session.AddToWhitelist(newWalletAddress, &sponsorGas)
	if err != nil {
		fmt.Println("AddToWhitelist error: ", err)
		panic(err)
	}
	fmt.Println("[UpdatePayMaster] done, hash!!!:", tx.Hash().Hex())

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
