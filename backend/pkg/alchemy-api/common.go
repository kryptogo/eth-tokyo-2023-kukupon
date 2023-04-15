package alchemyapi

import (
	"errors"
	"fmt"
	"math/big"
	"os"

	"github.com/go-resty/resty/v2"
)

var (
	httpClient *resty.Client
	Nonce      *big.Int
)

func init() {
	httpClient = resty.New()
	nonce, err := GetNextNonce(os.Getenv("SIGNING_WALLET_ADDRESS"))
	if err != nil {
		panic(err)
	}
	Nonce = hexToBigInt(nonce)

}

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
	fmt.Println("GetNextNonce respData: ", respData)
	return respData.Result, nil
}

func GetAPIURL() string {
	apiKey := os.Getenv("ALCHEMY_API_KEY")
	if os.Getenv("CHAIN_ID") == "polygon" {
		url := fmt.Sprintf("https://polygon-mainnet.g.alchemyapi.io/v2/%s/%s", apiKey, "")
		return url
	} else {
		url := fmt.Sprintf("https://polygon-mumbai.g.alchemyapi.io/v2/%s/%s", apiKey, "")
		return url
	}
}

// hexToBigInt convert hex string to big.Int
func hexToBigInt(s string) *big.Int {
	if len(s) > 2 && s[:2] == "0x" {
		s = s[2:]
	}
	n := new(big.Int)
	n.SetString(s, 16)
	return n
}
