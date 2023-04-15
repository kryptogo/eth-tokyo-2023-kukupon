package db

import (
	"math/big"
)

type GraphqlRequest struct {
	Query         string            `json:"query"`
	Variables     map[string]string `json:"variables,omitempty"`
	OperationName string            `json:"operationName,omitempty"`
}

type Campaign struct {
	ID                string         `json:"id"`
	RequiredCondition []string       `json:"required_condition"`
	Sponsor           string         `json:"sponsor"`
	Image             string         `json:"image"`
	SponsorGas        big.Int        `json:"sponsor_gas"`
	GraphqlReq        GraphqlRequest `json:"graphql_req"`
	CouponAmount      int            `json:"coupon_amount"`
}

var Campaigns = []Campaign{
	{
		ID: "ape_2023",
		RequiredCondition: []string{
			"New BAYC holders within the past three months!",
		},
		GraphqlReq: GraphqlRequest{
			Query: `query BAYCHoldersENSAndImages {
			TokenBalances(
			  input: {filter: {tokenAddress: {_eq: "0xbc4ca0eda7647a8ab7c2061c2e118a18a936f13d"}, tokenType: {_in: [ERC1155, ERC721]}}, blockchain: ethereum, limit: 49, cursor: __REPLACE__}
			) {
			  TokenBalance {
				owner {
				  addresses
				  lastUpdatedTimestamp
				}
			  }
			  pageInfo {
				nextCursor
				prevCursor
			  }
			}
		  }`,
			OperationName: "BAYCHoldersENSAndImages",
		},
		Sponsor:      "Bored Ape",
		Image:        "https://wallet-static.kryptogo.com/public/announcement/kg_2023/kukupon/coupon/bayc.png",
		SponsorGas:   *big.NewInt(10000000000000000), // 0.01
		CouponAmount: 2,
	},
	{
		ID: "1inch",
		RequiredCondition: []string{
			"Total transaction amount exceeds 3000 USDC.",
		},
		GraphqlReq: GraphqlRequest{
			Query: `
		query GetUserInteractedWithTokenAddress {
			TokenTransfers(
				input: {
					filter: {
						_or: [
							{from: {_eq: "0x111111111117dC0aa78b770fA6A738034120C302"}},
							{to: {_eq: "0x111111111117dC0aa78b770fA6A738034120C302"}}
						],
						{formattedAmount: {_gte: 3000}}
					},
					blockchain: ethereum,
					limit: 30
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
	`,
			OperationName: "GetUserInteractedWithTokenAddress",
		},
		Sponsor:      "1inch",
		Image:        "https://wallet-static.kryptogo.com/public/announcement/kg_2023/kukupon/coupon/1inch.png",
		SponsorGas:   *big.NewInt(10000000000000000), // 0.01
		CouponAmount: 2,
	},
	{
		ID: "azuki_2023",
		RequiredCondition: []string{
			"To all Azuki NFT holders!",
		},
		GraphqlReq: GraphqlRequest{
			Query: `query AzukiHoldersENSAndImages {
			TokenBalances(
			  input: {filter: {tokenAddress: {_eq: "0xed5af388653567af2f388e6224dc7c4b3241c544"}, tokenType: {_in: [ERC1155, ERC721]}}, blockchain: ethereum, limit: 49, cursor: __REPLACE__}
			) {
			  TokenBalance {
				owner {
				  addresses
				}
			  }
			  pageInfo {
				nextCursor
				prevCursor
			  }
			}
		  }`,
			OperationName: "AzukiHoldersENSAndImages",
		},
		Sponsor:      "Azuki",
		Image:        "https://wallet-static.kryptogo.com/public/announcement/kg_2023/kukupon/coupon/azuki.png",
		SponsorGas:   *big.NewInt(10000000000000000), // 0.01
		CouponAmount: 3,
	},
	{
		ID: "dev",
		RequiredCondition: []string{
			"Hold at least 1 KYGC token",
		},
		GraphqlReq: GraphqlRequest{
			Query: `query KGYCHoldersENSAndImages {
			TokenBalances(
			  input: {filter: {tokenAddress: {_eq: "0xa82fa2c0fd1fc6bb964d9302d3507b88a5f1b8d0"}, tokenType: {_in: [ERC1155, ERC721]}}, blockchain: polygon, limit: 30}
			) {
			  TokenBalance {
				owner {
				  addresses
				}
			  }
			  pageInfo {
				nextCursor
				prevCursor
			  }
			}
		  }`,
			OperationName: "KGYCHoldersENSAndImages",
		},
		Sponsor: "KryptoGO",
		Image:   "https://wallet-static.kryptogo.com/public/announcement/kg_2023/kukupon/coupon/kryptogo.png",
		// SponsorGas: *big.NewInt(500000000000000000), // 0.5 matic
		SponsorGas:   *big.NewInt(18000000000000000), // 0.01 mumbai //TODO:
		CouponAmount: 3,
	},

	{
		ID: "curve",
		RequiredCondition: []string{
			"Stake a transaction of over 1000 U!",
		},
		GraphqlReq: GraphqlRequest{
			Query: `
	query Curve {
		TokenTransfers(
			input: {
				filter: {
					_or: [
						{from: {_eq: "0x45F783CCE6B7FF23B2ab2D70e416cdb7D6055f51"}},
						{to: {_eq: "0x45F783CCE6B7FF23B2ab2D70e416cdb7D6055f51"}}
					],
					{formattedAmount: {_gte: 3000}}
				},
				blockchain: ethereum,
				limit: 30
			}
		) {
			TokenTransfer {
				from {
					addresses
				}
				to {
					addresses
				}
				type
			}
		}
	}
`,
			OperationName: "Curve",
		},
		Sponsor:      "Curve",
		Image:        "https://wallet-static.kryptogo.com/public/announcement/kg_2023/kukupon/coupon/curve.png",
		SponsorGas:   *big.NewInt(10000000000000000), // 0.01
		CouponAmount: 2,
	},
	{
		ID: "ens",
		RequiredCondition: []string{
			"Holding an ENS <3",
		},
		GraphqlReq: GraphqlRequest{
			Query: `
	query Ens {
		TokenTransfers(
			input: {
				filter: {
					_or: [
						{from: {_eq: "0xC18360217D8F7Ab5e7c516566761Ea12Ce7F9D72"}},
						{to: {_eq: "0xC18360217D8F7Ab5e7c516566761Ea12Ce7F9D72"}}
					]
				},
				blockchain: ethereum,
				limit: 30
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
`,
			OperationName: "Ens",
		},
		Sponsor:      "ENS",
		Image:        "https://wallet-static.kryptogo.com/public/announcement/kg_2023/kukupon/coupon/ens.png",
		SponsorGas:   *big.NewInt(10000000000000000), // 0.01
		CouponAmount: 2,
	},
	{
		ID: "LooksRare",
		RequiredCondition: []string{
			"With five transactions on LooksRare.",
		},
		GraphqlReq: GraphqlRequest{
			Query: `
	query LooksRare {
		TokenTransfers(
			input: {
				filter: {
					_or: [
						{from: {_eq: "0x0000000000E655fAe4d56241588680F86E3b2377"}},
						{to: {_eq: "0x0000000000E655fAe4d56241588680F86E3b2377"}}
					]
				},
				blockchain: ethereum,
				limit: 30
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
`,
			OperationName: "LooksRare",
		},
		Sponsor:      "LooksRare",
		Image:        "https://wallet-static.kryptogo.com/public/announcement/kg_2023/kukupon/coupon/looksrare.png",
		SponsorGas:   *big.NewInt(10000000000000000), // 0.01
		CouponAmount: 2,
	},
	{
		ID: "Lido",
		RequiredCondition: []string{
			"With five transactions on Lido.",
		},
		GraphqlReq: GraphqlRequest{
			Query: `
	query Lido {
		TokenTransfers(
			input: {
				filter: {
					_or: [
						{from: {_eq: "0x5A98FcBEA516Cf06857215779Fd812CA3beF1B32"}},
						{to: {_eq: "0x5A98FcBEA516Cf06857215779Fd812CA3beF1B32"}}
					]
				},
				blockchain: ethereum,
				limit: 30
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
`,
			OperationName: "Lido",
		},
		Sponsor:      "Lido",
		Image:        "https://wallet-static.kryptogo.com/public/announcement/kg_2023/kukupon/coupon/lido.png",
		SponsorGas:   *big.NewInt(10000000000000000), // 0.01
		CouponAmount: 2,
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
