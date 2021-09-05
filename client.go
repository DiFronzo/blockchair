package blockchair

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	apiRoot = "https://api.blockchair.com/"
	Hash = "^0x[0-9a-f]{64}$"
)

func Contains(slice []string, item string) bool {
	set := make(map[string]struct{}, len(slice))
	for _, s := range slice {
		set[s] = struct{}{}
	}

	_, ok := set[item]
	return ok
}

func GetSupportedCrypto() []string {
	return []string{"bitcoin", "bitcoin-cash", "litecoin", "bitcoin-sv", "dogecoin", "dash", "groestlcoin", "zcash", "ecash", "bitcoin/testnet"}
}

func GetSupportedCryptoEth() []string {
	return []string{"ethereum/testnet", "ethereum"}
}

type Client struct {
	*http.Client
}

func (c *Client) loadResponse(path string, i interface{}) error {
	fullPath := apiRoot + path

	fmt.Println("querying..." + fullPath)
	rsp, e := c.Get(fullPath)
	if e != nil {
		return e
	}

	defer rsp.Body.Close()

	b, e := ioutil.ReadAll(rsp.Body)
	if e != nil {
		return e
	}
	if rsp.Status[0] != '2' {
		return fmt.Errorf("expected status 2xx, got %s: %s", rsp.Status, string(b))
	}

	err := json.Unmarshal(b, &i)
	if err != nil {
		fmt.Printf("Error parsing JSON string - %s", err)
	}
	return err
}

func New() (*Client, error) {
	return &Client{Client: &http.Client{}}, nil
}
