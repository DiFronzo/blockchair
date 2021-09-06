package blockchair

import (
	"fmt"
	"log"
	"strings"
)

type DataAddress struct {
	Data    map[string]AddressInfo `json:"data"`
	Context ContextAddress         `json:"context"`
}

type DataAddressEth struct {
	Data    map[string]AddressInfoEth `json:"data"`
	Context ContextAddress            `json:"context"`
}

type DataAddresses struct {
	Data    AddressesInfo  `json:"data"`
	Context ContextAddress `json:"context"`
}

type DataXpub struct {
	Data    map[string]XpubInfo `json:"data"`
	Context ContextAddress      `json:"context"`
}

type XpubInfo struct {
	Xpub         Xpub               `json:"xpub"`
	Addresses    map[string]Address `json:"addresses"`
	Transactions []string           `json:"transactions"`
	Utxo         []Utxo             `json:"utxo"`
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
	Set          Set                `json:"set"`
	Addresses    map[string]Address `json:"addresses"`
	Transactions []string           `json:"transactions"`
	Utxo         []Utxo             `json:"utxo"`
}

type AddressInfoEth struct {
	Address AddressEth     `json:"address"`
	Calls   []CallsAddress `json:"calls"`
}

type Set struct {
	AddressCount       int     `json:"address_count"`
	Balance            int64   `json:"balance"`
	BalanceUsd         float64 `json:"balance_usd"`
	Received           int64   `json:"received"`
	Spent              int     `json:"spent"`
	OutputCount        int     `json:"output_count"`
	UnspentOutputCount int     `json:"unspent_output_count"`
	FirstSeenReceiving string  `json:"first_seen_receiving"`
	LastSeenReceiving  string  `json:"last_seen_receiving"`
	FirstSeenSpending  string  `json:"first_seen_spending"`
	LastSeenSpending   string  `json:"last_seen_spending"`
	TransactionCount   int     `json:"transaction_count"`
}

type AddressInfo struct {
	Address      Address  `json:"address"`
	Transactions []string `json:"transactions"`
	Utxo         []Utxo   `json:"utxo"`
}

type AddressEth struct {
	Type                string  `json:"type"`
	ContractCodeHex     string  `json:"contract_code_hex"`
	ContractCreated     string  `json:"contract_created"`
	ContractDestroyed   string  `json:"contract_destroyed"`
	Balance             string  `json:"balance"`
	BalanceUsd          float64 `json:"balance_usd"`
	ReceivedApproximate string  `json:"received_approximate"`
	ReceivedUsd         float64 `json:"received_usd"`
	SpentApproximate    string  `json:"spent_approximate"`
	SpentUsd            float64 `json:"spent_usd"`
	FeesApproximate     string  `json:"fees_approximate"`
	FeesUsd             float64 `json:"fees_usd"`
	ReceivingCallCount  int     `json:"receiving_call_count"`
	SpendingCallCount   int     `json:"spending_call_count"`
	CallCount           int     `json:"call_count"`
	TransactionCount    int     `json:"transaction_count"`
	FirstSeenReceiving  string  `json:"first_seen_receiving"`
	LastSeenReceiving   string  `json:"last_seen_receiving"`
	FirstSeenSpending   string  `json:"first_seen_spending"`
	LastSeenSpending    string  `json:"last_seen_spending"`
	Nonce               int     `json:"nonce"`
}

type Address struct {
	Type               string  `json:"type"`
	ScriptHex          string  `json:"script_hex"`
	Balance            int64   `json:"balance"`
	BalanceUsd         float32 `json:"balance_usd"`
	Received           int64   `json:"received"`
	ReceivedUsd        float32 `json:"received_usd"`
	Spent              int     `json:"spent"`
	SpentUsd           float32 `json:"spent_usd"`
	OutputCount        int     `json:"output_count"`
	UnspentOutputCount int     `json:"unspent_output_count"`
	FirstSeenReceiving string  `json:"first_seen_receiving"`
	LastSeenReceiving  string  `json:"last_seen_receiving"`
	FirstSeenSpending  string  `json:"first_seen_spending"`
	LastSeenSpending   string  `json:"last_seen_spending"`
	ScripthashType     string  `json:"scripthash_type"`
	TransactionCount   int     `json:"transaction_count"`
}

type Utxo struct {
	BlockID         int    `json:"block_id"`
	TransactionHash string `json:"transaction_hash"`
	Index           int    `json:"index"`
	Value           int    `json:"value"`
}

type CallsAddress struct {
	BlockID         int     `json:"block_id"`
	TransactionHash string  `json:"transaction_hash"`
	Index           string  `json:"index"`
	Time            string  `json:"time"`
	Sender          string  `json:"sender"`
	Recipient       string  `json:"recipient"`
	Value           float64 `json:"value"`
	ValueUsd        float64 `json:"value_usd"`
	Transferred     bool    `json:"transferred"`
}

// ContextAddress TODO! FIX "CHECKED" INTO A SLICE
type ContextAddress struct {
	Code           int      `json:"code"`
	Source         string   `json:"source"`
	Limit          string   `json:"limit"`
	Offset         string   `json:"offset"`
	Results        int      `json:"results"`
	Checked        []string `json:"checked"`
	State          int      `json:"state"`
	MarketPriceUsd float32  `json:"market_price_usd"`
	Cache          *Cache   `json:"cache"`
	API            *Api     `json:"api"`
	Server         string   `json:"server"`
	Time           float32  `json:"time"`
	RenderTime     float32  `json:"render_time"`
	FullTime       float32  `json:"full_time"`
	RequestCost    float32  `json:"request_cost"`
}

func (c *Client) GetAddress(crypto string, address string) (*DataAddress, error) {
	if !Contains(GetSupportedCrypto(), crypto) {
		log.Fatalf("error: %v is not supported", crypto)
	}
	rsp := &DataAddress{}
	var path = crypto + "/dashboards/address/" + address
	e := c.loadResponse(path, rsp)

	if e != nil {
		fmt.Print(e)
	}
	return rsp, e
}

func (c *Client) GetAddresses(crypto string, addresses []string) (*DataAddresses, error) {
	if !Contains(GetSupportedCrypto(), crypto) {
		log.Fatalf("error: %v is not supported", crypto)
	}
	rsp := &DataAddresses{}
	var path = crypto + "/dashboards/addresses/" + strings.Join(addresses, ",")
	e := c.loadResponse(path, rsp)

	if e != nil {
		fmt.Print(e)
	}
	return rsp, e
}

func (c *Client) GetXpub(crypto string, extendedKey string) (*DataXpub, error) {
	// xpub (supported for all blockchains), ypub (supported for Bitcoin, Litecoin, Groestlcoin, and Bitcoin Testnet only)
	// zpub (supported for Bitcoin, Litecoin, Groestlcoin, and Bitcoin Testnet only)
	rsp := &DataXpub{}
	var path = crypto + "/dashboards/xpub/" + extendedKey
	e := c.loadResponse(path, rsp)

	if e != nil {
		fmt.Print(e)
	}
	return rsp, e
}

func (c *Client) GetAddressEth(crypto string, address string) (*DataAddressEth, error) {
	if !Contains(GetSupportedCryptoEth(), crypto) {
		log.Fatalf("error: %v is not supported", crypto)
	}

	rsp := &DataAddressEth{}
	var path = crypto + "/dashboards/address/" + address
	e := c.loadResponse(path, rsp)

	if e != nil {
		fmt.Print(e)
	}
	return rsp, e
}
