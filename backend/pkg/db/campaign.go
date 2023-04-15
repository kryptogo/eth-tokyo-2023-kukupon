package db

import "fmt"

type Campaign struct {
	ID                string
	RequiredCondition string
	PaymasterAddress  string
	Sponsor           string
}

var Campaigns = []Campaign{
	{
		ID: "eth_global_tokyo_2023",
		RequiredCondition: fmt.Sprintf(`
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
		PaymasterAddress: "0x000",
		Sponsor:          "Sugar Daddy",
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
