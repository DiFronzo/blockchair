package blockchair

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	apiRoot = "https://api.blockchair.com/"
)

//func getSupportedCrypto() []string {
//	return []string{"bitcoin", "bitcoin-cash", "litecoin", "bitcoin-sv", "dogecoin", "dash", "groestlcoin", "zcash", "ecash", "ethereum"}
//}

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

	return json.Unmarshal(b, &i)
}

func New() (*Client, error) {
	return &Client{Client: &http.Client{}}, nil
}
