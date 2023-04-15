package service

import (
	"fmt"
	"testing"
)

func TestGenerateWallets(t *testing.T) {
	coupon, wallet := GenerateWallets(1)
	fmt.Println(coupon[0])
	fmt.Println(wallet[0])

	wallets4337 := GenerateWallets4337(wallet)
	fmt.Println(wallets4337[0])
}

func TestGetCoupons(t *testing.T) {
	camapignId := "kygc_holder"
	GetCoupons(camapignId)
}
