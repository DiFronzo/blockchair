package blockchair

import (
	"encoding/json"
	"errors"
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

// Errors it is a set of errors returned when working with the package.
var (
	ErrTHW = errors.New("blockchair: transaction hash is wrong")
	ErrERC = errors.New("blockchair: ERC-20 token is wrong")
	ErrSC  = errors.New("blockchair: the Bitcoin-like cryptocurrency is not supported")
	ErrSCE = errors.New("blockchair: the Ethereum cryptocurrency is not supported")
	ErrSCG = errors.New("blockchair: the cryptocurrency is not supported")
	ErrCGD = errors.New("blockchair: cannot get data on url")
	ErrCRR = errors.New("blockchair: could not read answer response")
	ErrRPE = errors.New("blockchair: response parsing error")
	ErrIRS = errors.New("blockchair: incorrect response status")
	ErrMAX = errors.New("blockchair: the maximum number of addresses is 100")
	ErrETH = errors.New("blockchair: can only handle one Ethereum cryptocurrency address")
)

// GetSupportedCrypto List of supported Bitcoin-like crypto.
func GetSupportedCrypto() []string {
	return []string{"bitcoin", "bitcoin-cash", "litecoin", "bitcoin-sv", "dogecoin", "dash", "groestlcoin", "zcash", "ecash", "bitcoin/testnet"}
}

// GetSupportedCryptoEth List of supported Ethereum crypto.
func GetSupportedCryptoEth() []string {
	return []string{"ethereum/testnet", "ethereum"}
}

// GetSupportedCryptoMultichain List of supported crypto for multichain address check.
func GetSupportedCryptoMultichain() []string {
	return []string{"bitcoin", "bitcoin-cash", "litecoin", "bitcoin-sv", "dash", "groestlcoin", "zcash", "ethereum"}
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
		if c.APIKey != "" {
			options["key"] = c.APIKey
		}
		values := url.Values{}
		for k, v := range options {
			values.Set(k, v)
		}
		fullPath += "?" + (values.Encode())
	}

	req, e := http.NewRequest("GET", fullPath, nil)
	if e != nil {
		return c.err2(ErrCGD, e)
	}
	// fmt.Println("querying..." + fullPath)
	req.Header.Set("User-Agent", c.userAgent())

	resp, e := c.client.Do(req)
	if e != nil {
		return c.err3(ErrCGD, e, resp)
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(resp.Body)

	b, e := ioutil.ReadAll(resp.Body)
	if e != nil {
		return c.err3(ErrCRR, e, resp)
	}

	if resp.Status[0] != '2' {
		return c.err3(ErrIRS, e, resp)
		//return fmt.Errorf("expected status 2xx, got %s: %s", resp.Status, string(b))
	}

	if err := json.Unmarshal(b, &i); err != nil {
		fmt.Println(err)
		return c.err3(ErrRPE, e, resp)
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
