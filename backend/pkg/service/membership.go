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

	value, ok := members[strings.ToLower(from)]
	if !ok {
		return false
	}
	return value
}

// type GraphqlResponse struct {
// 	Data   interface{}            `json:"data"`
// 	Errors []map[string]interface{} `json:"errors,omitempty"`
// }

func getMembers(campaignID, operationName, query string) map[string]bool {
	result := map[string]bool{}
	cursor := ""
	for {

		// fmt.Println("[getMembers] cursor:", cursor)
		newquery := strings.Replace(query, "__REPLACE__", `"`+cursor+`"`, 1)
		variables := map[string]interface{}{}

		payload := struct {
			Query         string                 `json:"query"`
			Variables     map[string]interface{} `json:"variables"`
			OperationName string                 `json:"operationName"`
		}{
			Query:         newquery,
			Variables:     variables,
			OperationName: operationName,
		}
		// fmt.Println("[getMembers] payload:", payload)
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
		var addresses map[string]bool
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
		for address := range addresses {
			result[strings.ToLower(address)] = true
		}

		// 更新 cursor
		cursor = nextCursor
		// 如果沒有下一頁了，結束 loop
		if cursor == "" {
			break
		}
	}

	fmt.Println("members size:", len(result))
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

func getAddressesFromRespDapp(response Response) map[string]bool {
	addresses := map[string]bool{}

	for _, tt := range response.Data.TokenTransfers.TokenTransfer {
		if from := tt.From.Addresses; from != nil {
			for _, a := range from {
				addresses[a] = true
			}
		}
		if to := tt.To.Addresses; to != nil {
			for _, a := range to {
				addresses[a] = true
			}
		}
	}
	return addresses
}

func getAddressesFromRespNFT(response ResponseNFT) map[string]bool {
	addresses := map[string]bool{}

	for _, tt := range response.Data.TokenBalances.TokenBalance {
		if from := tt.Owner.Addresses; from != nil {
			for _, a := range from {
				addresses[a] = true
			}
		}
	}
	return addresses
}
