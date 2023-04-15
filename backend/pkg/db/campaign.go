package db

import (
	"fmt"
	"math/big"
)

type Campaign struct {
	ID                string   `json:"id"`
	Query             string   `json:"query"`
	RequiredCondition []string `json:"required_condition"`
	Sponsor           string   `json:"sponsor"`
	Image             string   `json:"image"`
	SponsorGas        big.Int  `json:"sponsor_gas"`
}

var Campaigns = []Campaign{
	{
		ID: "eth_global_tokyo_2023",
		RequiredCondition: []string{
			"Interacted with dapp 1inch at least 100 times",
			"Total transaction amount exceeds 3000 USDC.",
		},
		Query: fmt.Sprintf(`
		query GetUserInteractedWithTokenAddress {
			TokenTransfers(
				input: {
					filter: {
						_or: [
							{from: {_eq: "%s"}},
							{to: {_eq: "%s"}}
						]
					},
					blockchain: ethereum,
					limit: 10
				}
			) {
				TokenTransfer {
					from {
						addresses
					}
					to {
						addresses
					}
				}
			}
		}
	`, "0x111111111117dC0aa78b770fA6A738034120C302", "0x111111111117dC0aa78b770fA6A738034120C302"),
		Sponsor:    "1inch",
		Image:      "https://1inch.io/assets/social-image/main-cover-2.png",
		SponsorGas: *big.NewInt(500000000000000000), // 0.5
	},
	{
		ID: "kygc_holder",
		RequiredCondition: []string{
			"Hold at least 1 KYGC token",
		},
		Query: fmt.Sprintf(`
		query GetUserInteractedWithTokenAddress {
			TokenTransfers(
				input: {
					filter: {
						_or: [
							{from: {_eq: "%s"}},
							{to: {_eq: "%s"}}
						]
					},
					blockchain: ethereum,
					limit: 10
				}
			) {
				TokenTransfer {
					from {
						addresses
					}
					to {
						addresses
					}
				}
			}
		}
	`, "0xa82fa2c0fd1fc6bb964d9302d3507b88a5f1b8d0", "0xa82fa2c0fd1fc6bb964d9302d3507b88a5f1b8d0"),
		Sponsor:    "KryptoGO",
		Image:      "https://twnewshub.com/wp-content/uploads/2021/12/Android-topic.png",
		SponsorGas: *big.NewInt(500000000000000000), // 0.5
	},
}

func GetCampaigns() []Campaign {
	return Campaigns
}

func GetCampaign(campaignID string) *Campaign {
	for _, campaign := range Campaigns {
		if campaign.ID == campaignID {
			return &campaign
		}
	}
	return nil
}
