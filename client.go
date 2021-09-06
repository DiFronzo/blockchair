package blockchair

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	apiRoot = "https://api.blockchair.com/"
	Hash    = "^0x[0-9a-f]{64}$"
	//UserAgent = "blockchair-api-go-v1"
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

func New(k interface{}) (*Client, error) {
	return &Client{Client: &http.Client{}, api: k}, nil
}
