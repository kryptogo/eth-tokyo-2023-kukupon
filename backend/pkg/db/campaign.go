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
			"Welcome to new users who own BAYC nft within the past three months!",
		},
		GraphqlReq: GraphqlRequest{
			Query: `query BAYCHoldersENSAndImages {
			TokenBalances(
			  input: {filter: {tokenAddress: {_eq: "0xbc4ca0eda7647a8ab7c2061c2e118a18a936f13d"}, tokenType: {_in: [ERC1155, ERC721]}}, blockchain: ethereum, limit: 49, cursor: __REPLACE__}
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
			OperationName: "BAYCHoldersENSAndImages",
		},
		Sponsor:      "Bored Ape",
		Image:        "https://i.seadn.io/gae/Ju9CkWtV-1Okvf45wo8UctR-M9He2PjILP0oOvxE89AyiPPGtrR3gysu1Zgy0hjd2xKIgjJJtWIc0ybj4Vd7wv8t3pxDGHoJBzDB?auto=format&w=3840",
		SponsorGas:   *big.NewInt(10000000000000000), // 0.01
		CouponAmount: 2,
	},
	{
		ID: "eth_global_tokyo_2023",
		RequiredCondition: []string{
			"Interacted with dapp 1inch at least 100 times",
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
			OperationName: "GetUserInteractedWithTokenAddress",
		},
		Sponsor:      "1inch",
		Image:        "https://1inch.io/assets/social-image/main-cover-2.png",
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
		Image:        "https://i.seadn.io/gcs/files/c1c0e5d22d1d52e834c8a72566c644e9.png?auto=format&w=384",
		SponsorGas:   *big.NewInt(10000000000000000), // 0.01
		CouponAmount: 3,
	},
	// {
	// 	ID: "kygc_holder",
	// 	RequiredCondition: []string{
	// 		"Hold at least 1 KYGC token",
	// 	},
	// 	GraphqlReq: GraphqlRequest{
	// 		Query: `query KGYCHoldersENSAndImages {
	// 		TokenBalances(
	// 		  input: {filter: {tokenAddress: {_eq: "0xa82fa2c0fd1fc6bb964d9302d3507b88a5f1b8d0"}, tokenType: {_in: [ERC1155, ERC721]}}, blockchain: polygon, limit: 49, cursor: __REPLACE__}
	// 		) {
	// 		  TokenBalance {
	// 			owner {
	// 			  addresses
	// 			}
	// 		  }
	// 		  pageInfo {
	// 			nextCursor
	// 			prevCursor
	// 		  }
	// 		}
	// 	  }`,
	// 		OperationName: "KGYCHoldersENSAndImages",
	// 	},
	// 	Sponsor: "KryptoGO",
	// 	Image:   "https://twnewshub.com/wp-content/uploads/2021/12/Android-topic.png",
	// 	// SponsorGas: *big.NewInt(500000000000000000), // 0.5 matic
	// 	SponsorGas:   *big.NewInt(10000000000000000), // 0.01 mumbai
	// 	CouponAmount: 1,
	// },
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
		Image:   "https://twnewshub.com/wp-content/uploads/2021/12/Android-topic.png",
		// SponsorGas: *big.NewInt(500000000000000000), // 0.5 matic
		SponsorGas:   *big.NewInt(18000000000000000), // 0.01 mumbai //TODO:
		CouponAmount: 3,
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
