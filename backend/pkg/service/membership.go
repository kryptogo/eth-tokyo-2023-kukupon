package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/kryptogo/eth-tokyo-2023-kukuponpon/pkg/db"
)

func VerifyMembership(campaignID, from string) bool {
	campaign := db.GetCampaign(campaignID)
	members := getMembers(campaignID, campaign.GraphqlReq.OperationName, campaign.GraphqlReq.Query)
	for _, member := range members {
		if member == from {
			return true
		}
	}
	return false
}

// type GraphqlResponse struct {
// 	Data   interface{}            `json:"data"`
// 	Errors []map[string]interface{} `json:"errors,omitempty"`
// }

func getMembers(campaignID, operationName, query string) []string {
	var result []string
	cursor := ""

	for {
		variables := map[string]interface{}{
			"cursor": "",
		}

		payload := struct {
			Query         string                 `json:"query"`
			Variables     map[string]interface{} `json:"variables"`
			OperationName string                 `json:"operationName"`
		}{
			Query:         query,
			Variables:     variables,
			OperationName: operationName,
		}

		requestBody, err := json.Marshal(payload)
		if err != nil {
			fmt.Println("error:", err)
			return nil
		}

		req, err := http.NewRequest("POST", "https://api.airstack.xyz/gql", bytes.NewBuffer(requestBody))
		if err != nil {
			fmt.Println("error:", err)
			return nil
		}

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("accept", "application/json, multipart/mixed")
		req.Header.Set("Authorization", os.Getenv("AIRSTACK_API_KEY"))

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			fmt.Println("error:", err)
			return nil
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("error:", err)
			return nil
		} else if strings.Contains(string(body), "errors") {
			fmt.Println("error:", string(body))
			return nil
		}
		var nextCursor string
		var addresses []string
		if campaignID == "eth_global_tokyo_2023" {
			var response Response

			err := json.Unmarshal(body, &response)
			if err != nil {
				panic(err)
			}
			addresses = getAddressesFromRespDapp(response)
			nextCursor = response.Data.TokenTransfers.PageInfo.NextCursor
		} else {
			var response ResponseNFT

			err := json.Unmarshal(body, &response)
			if err != nil {
				panic(err)
			}
			addresses = getAddressesFromRespNFT(response)
			nextCursor = response.Data.TokenBalances.PageInfo.NextCursor
		}
		result = append(result, addresses...)
		// 更新 cursor
		cursor = nextCursor

		// 如果沒有下一頁了，結束 loop
		if cursor == "" {
			break
		}
	}

	fmt.Println("all members:", result)
	return result
}

type Response struct {
	Data struct {
		TokenTransfers struct {
			TokenTransfer []struct {
				From struct {
					Addresses []string `json:"addresses"`
				} `json:"from"`
				To struct {
					Addresses []string `json:"addresses"`
				} `json:"to"`
			} `json:"TokenTransfer"`
			PageInfo struct {
				NextCursor string `json:"nextCursor"`
			} `json:"pageInfo"`
		} `json:"TokenTransfers"`
	} `json:"data"`
}

type ResponseNFT struct {
	Data struct {
		TokenBalances struct {
			TokenBalance []struct {
				Owner struct {
					Addresses []string `json:"addresses"`
				} `json:"owner"`
			} `json:"TokenBalance"`
			PageInfo struct {
				NextCursor string `json:"nextCursor"`
			} `json:"pageInfo"`
		} `json:"TokenBalances"`
	} `json:"data"`
}

func getAddressesFromRespDapp(response Response) []string {
	var addresses []string
	for _, tt := range response.Data.TokenTransfers.TokenTransfer {
		if from := tt.From.Addresses; from != nil {
			for _, a := range from {
				addresses = append(addresses, a)
			}
		}
		if to := tt.To.Addresses; to != nil {
			for _, a := range to {
				addresses = append(addresses, a)
			}
		}
	}
	return addresses
}

func getAddressesFromRespNFT(response ResponseNFT) []string {
	var addresses []string
	for _, tt := range response.Data.TokenBalances.TokenBalance {
		if from := tt.Owner.Addresses; from != nil {
			for _, a := range from {
				addresses = append(addresses, a)
			}
		}
	}
	return addresses
}
