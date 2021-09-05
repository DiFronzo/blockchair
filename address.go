package blockchair

import (
	"fmt"
	"strings"
)

type DataAddress struct {
	Data	map[string]AddressInfo	`json:"data"`
	Context    ContextAddress    `json:"context"`
}

type DataAddresses struct {
	Data	AddressesInfo	`json:"data"`
	Context    ContextAddress    `json:"context"`
}

type DataXpub struct {
	Data	map[string]XpubInfo	`json:"data"`
	Context    ContextAddress    `json:"context"`
}

type XpubInfo struct {
	Xpub	Xpub	`json:"xpub"`
	Addresses map[string]Address `json:"addresses"`
	Transactions []string `json:"transactions"`
	Utxo	[]Utxo	`json:"utxo"`
}

type Xpub struct {
	AddressCount       int     `json:"address_count"`
	Balance            int     `json:"balance"`
	BalanceUsd         float64 `json:"balance_usd"`
	Received           int     `json:"received"`
	Spent              int     `json:"spent"`
	OutputCount        int     `json:"output_count"`
	UnspentOutputCount int     `json:"unspent_output_count"`
	FirstSeenReceiving string  `json:"first_seen_receiving"`
	LastSeenReceiving  string  `json:"last_seen_receiving"`
	FirstSeenSpending  string  `json:"first_seen_spending"`
	LastSeenSpending   string  `json:"last_seen_spending"`
	TransactionCount   int     `json:"transaction_count"`
}

type AddressesInfo struct {
	Set Set `json:"set"`
	Addresses map[string]Address `json:"addresses"`
	Transactions []string `json:"transactions"`
	Utxo	[]Utxo	`json:"utxo"`
}

type Set struct {
	AddressCount       int         `json:"address_count"`
	Balance            int64       `json:"balance"`
	BalanceUsd         float64     `json:"balance_usd"`
	Received           int64       `json:"received"`
	Spent              int         `json:"spent"`
	OutputCount        int         `json:"output_count"`
	UnspentOutputCount int         `json:"unspent_output_count"`
	FirstSeenReceiving string      `json:"first_seen_receiving"`
	LastSeenReceiving  string      `json:"last_seen_receiving"`
	FirstSeenSpending  string	   `json:"first_seen_spending"`
	LastSeenSpending   string 	   `json:"last_seen_spending"`
	TransactionCount   int         `json:"transaction_count"`
}

type AddressInfo struct {
	Address Address `json:"address"`
	Transactions []string `json:"transactions"`
	Utxo	[]Utxo	`json:"utxo"`
}

type Address struct {
	Type               string      `json:"type"`
	ScriptHex          string      `json:"script_hex"`
	Balance            int64       `json:"balance"`
	BalanceUsd         float32     `json:"balance_usd"`
	Received           int64       `json:"received"`
	ReceivedUsd        float32     `json:"received_usd"`
	Spent              int         `json:"spent"`
	SpentUsd           float32     `json:"spent_usd"`
	OutputCount        int         `json:"output_count"`
	UnspentOutputCount int         `json:"unspent_output_count"`
	FirstSeenReceiving string      `json:"first_seen_receiving"`
	LastSeenReceiving  string      `json:"last_seen_receiving"`
	FirstSeenSpending  interface{} `json:"first_seen_spending"`
	LastSeenSpending   interface{} `json:"last_seen_spending"`
	ScripthashType     interface{} `json:"scripthash_type"`
	TransactionCount   int         `json:"transaction_count"`
}

type Utxo struct {
	BlockID         int    `json:"block_id"`
	TransactionHash string `json:"transaction_hash"`
	Index           int    `json:"index"`
	Value           int    `json:"value"`
}

// ContextAddress TODO! FIX "CHECKED" INTO A SLICE
type ContextAddress struct {
	Code       	  	int    `json:"code"`
	Source			string `json:"source"`
	Limit       	string    `json:"limit"`
	Offset       	string    `json:"offset"`
	Results			int    `json:"results"`
	Checked			[]string `json:"checked"`
	State           int    `json:"state"`
	MarketPriceUsd  int      `json:"market_price_usd"`
	Cache			*Cache  `json:"cache"`
	API				*Api  `json:"api"`
	Server			string `json:"server"`
	Time			float32    `json:"time"`
	RenderTime		float32    `json:"render_time"`
	FullTime		float32    `json:"full_time"`
	RequestCost		float32   	`json:"request_cost"`
}

func (c *Client) GetAddress(crypto string, Address string) (*DataAddress, error) {
	rsp := &DataAddress{}
	var path = crypto + "/dashboards/address/" + Address
	e := c.loadResponse(path, rsp)

	if e != nil {
		fmt.Print(e)
	}
	return rsp, e
}

func (c *Client) GetAddresses(crypto string, Addresses []string) (*DataAddresses, error) {
	rsp := &DataAddresses{}
	var path = crypto + "/dashboards/addresses/" + strings.Join(Addresses, ",")
	e := c.loadResponse(path, rsp)

	if e != nil {
		fmt.Print(e)
	}
	return rsp, e
}


func (c *Client) GetXpub(crypto string, extendedKey string) (*DataXpub, error) {
	rsp := &DataXpub{}
	var path = crypto + "/dashboards/xpub/" + extendedKey
	e := c.loadResponse(path, rsp)

	if e != nil {
		fmt.Print(e)
	}
	return rsp, e
}