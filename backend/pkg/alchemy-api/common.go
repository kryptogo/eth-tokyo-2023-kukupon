package alchemyapi

import (
	"errors"
	"fmt"
	"os"

	"github.com/go-resty/resty/v2"
)

var (
	httpClient *resty.Client
)

type JsonRpcResp struct {
	JSONRpc string                 `json:"jsonrpc"`
	ID      int                    `json:"id"`
	Result  string                 `json:"result"`
	Error   map[string]interface{} `json:"error,omitempty"`
}

func GetNextNonce(from string) (string, error) {
	respData := JsonRpcResp{}
	data := map[string]interface{}{
		"jsonprc": "2.0",
		"method":  "eth_getTransactionCount",
		"params":  []string{from, "latest"},
		"id":      1,
	}
	resp, err := httpClient.R().
		SetHeader("Content-Type", "application/json").
		SetBody(data).
		SetResult(&respData).
		Post(GetAPIURL())

	if err != nil {
		return "", err
	}
	if resp.StatusCode() >= 400 {
		return "", errors.New(resp.String())
	}
	if respData.Error != nil {
		return "", fmt.Errorf("%v", respData.Error)
	}
	return respData.Result, nil
}

func GetAPIURL() string {
	apiKey := os.Getenv("ALCHEMY_API_KEY")

	url := fmt.Sprintf("https://polygon-mainnet.g.alchemyapi.io/v2/%s/%s", apiKey, "")
	return url
}
