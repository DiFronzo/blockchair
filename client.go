package blockchair

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Hash used to verify Ethereum hash.
const (
	apiRoot = "https://api.blockchair.com/"
	Hash    = "^0x[0-9a-f]{64}$"
	//UserAgent = "blockchair-api-go-v1"
)

// Contains used with GetSupportedCrypto and/or GetSupportedCryptoEth to verify correct crypto.
func Contains(slice []string, item string) bool {
	set := make(map[string]struct{}, len(slice))
	for _, s := range slice {
		set[s] = struct{}{}
	}

	_, ok := set[item]
	return ok
}

// GetSupportedCrypto List of supported Bitcoin-like crypto.
func GetSupportedCrypto() []string {
	return []string{"bitcoin", "bitcoin-cash", "litecoin", "bitcoin-sv", "dogecoin", "dash", "groestlcoin", "zcash", "ecash", "bitcoin/testnet"}
}

// GetSupportedCryptoEth List of supported Ethereum crypto.
func GetSupportedCryptoEth() []string {
	return []string{"ethereum/testnet", "ethereum"}
}

// Client specifies the mechanism by which individual API requests are made.
type Client struct {
	*http.Client
	api interface{}
}

func (c *Client) loadResponse(path string, i interface{}) error {
	fullPath := apiRoot + path
	if c.api != nil {
		fullPath = fullPath + "?api=" + fmt.Sprintf("%v", c.api)
	}

	fmt.Println("querying..." + fullPath)
	resp, e := c.Get(fullPath)
	if e != nil {
		panic(e)
	}

	defer func() {
		err := resp.Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	b, e := ioutil.ReadAll(resp.Body)
	if e != nil {
		log.Fatal(e)
	}
	if resp.Status[0] != '2' {
		return fmt.Errorf("expected status 2xx, got %s: %s", resp.Status, string(b))
	}

	err := json.Unmarshal(b, &i)
	if err != nil {
		fmt.Printf("Error parsing JSON string - %s", err)
	}
	return err
}

// New creates a new client instance the network internet.
func New(k interface{}) (*Client, error) {
	return &Client{Client: &http.Client{}, api: k}, nil
}
