package blockchair

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

// Hash used to verify Ethereum hash.
const (
	apiRoot   = "https://api.blockchair.com/"
	Hash      = "^0x[0-9a-f]{64}$"
	UserAgent = "blockchair-api-golang-v1"
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
	client *http.Client

	APIKey    string // API access key.
	UserAgent string // Optional additional User-Agent fragment.
}

func (c *Client) userAgent() string {
	c.UserAgent = strings.TrimSpace(c.UserAgent)
	if c.UserAgent == "" {
		return UserAgent
	}

	return UserAgent + " " + c.UserAgent
}

// LoadResponse to send a client request, which is then converted to the passed type.
func (c *Client) LoadResponse(path string, i interface{}, options map[string]string) error {
	fullPath := apiRoot + path

	if c.APIKey != "" && options == nil {
		fullPath += "?key=" + c.APIKey
	}

	if options != nil {
		options["key"] = c.APIKey
		values := url.Values{}
		for k, v := range options {
			values.Set(k, v)
		}
		fullPath += "?" + (values.Encode())
	}

	req, e := http.NewRequest("GET", fullPath, nil)
	if e != nil {
		panic(e)
	}
	fmt.Println("querying..." + fullPath)
	req.Header.Set("User-Agent", c.userAgent())

	resp, e := c.client.Do(req)
	if e != nil {
		panic(e)
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(resp.Body)

	b, e := ioutil.ReadAll(resp.Body)
	if e != nil {
		log.Fatal(e)
	}

	if resp.Status[0] != '2' {
		return fmt.Errorf("expected status 2xx, got %s: %s", resp.Status, string(b))
	}

	if err := json.Unmarshal(b, &i); err != nil {
		fmt.Printf("Error parsing JSON string - %s", err)
	}

	return nil
}

// New creates a new client instance the network internet.
func New() *Client {
	return &Client{client: &http.Client{}}
}

// SetClient http client setter.
func (c *Client) SetClient(client *http.Client) {
	if client == nil {
		panic("blockchair: impossible install the client as nil")
	}
	c.client = client
}
