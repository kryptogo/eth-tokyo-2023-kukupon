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
	members := getMembers(campaignID, campaign.OperationName, campaign.Query)
	for _, member := range members {
		if member == from {
			return true
		}
	}
	return false
}

func getMembers(campaignID, operationName, query string) []string {
	variables := map[string]interface{}{}

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
	// req.Header.Set("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/112.0.0.0 Safari/537.36")
	req.Header.Set("accept", "application/json, multipart/mixed")
	// req.Header.Set("authority", "api.airstack.xyz")
	// req.Header.Set("sec-fetch-mode", "cors")
	// req.Header.Set("origin", "https://app.airstack.xyz")
	// req.Header.Set("referer", "https://app.airstack.xyz/")
	// req.Header.Set("sec-ch-ua", `"Chromium";v="112", "Google Chrome";v="112", "Not:A-Brand";v="99"`)
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

	var addresses []string
	if campaignID == "eth_global_tokyo_2023" {
		addresses = getAddressesFromRespDapp(body)
	} else {
		addresses = getAddressesFromRespNFT(body)
	}
	fmt.Println("all members:", addresses)
	return addresses
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
		} `json:"TokenBalances"`
	} `json:"data"`
}

func getAddressesFromRespDapp(body []byte) []string {
	var response Response

	err := json.Unmarshal(body, &response)
	if err != nil {
		panic(err)
	}

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

func getAddressesFromRespNFT(body []byte) []string {
	var response ResponseNFT

	err := json.Unmarshal(body, &response)
	if err != nil {
		panic(err)
	}

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
