package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/kryptogo/eth-tokyo-2023-kukuponpon/pkg/db"
)

func VerifyMembership(campaignID, from string) bool {
	campaign := db.GetCampaign(campaignID)
	members := getMembers(campaign.Query)
	for _, member := range members {
		if member == from {
			return true
		}
	}
	return false
}

func getMembers(query string) []string {
	variables := map[string]interface{}{}
	operationName := "GetUserInteractedWithTokenAddress"

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
	req.Header.Set("Accept", "application/json, multipart/mixed")
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
	}
	addresses := getAddressesFromResp(body)
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

func getAddressesFromResp(body []byte) []string {
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
