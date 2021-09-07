package blockchair

import (
	"log"
	"strings"
)

// DataAddress includes full server response to address request.
type DataAddress struct {
	Data    map[string]AddressInfo `json:"data"`
	Context ContextAddress         `json:"context"`
}

// DataAddressEth includes full server response to address request for Ethereum.
type DataAddressEth struct {
	Data    map[string]AddressInfoEth `json:"data"`
	Context ContextAddress            `json:"context"`
}

// DataAddresses includes full server response to addresses request.
type DataAddresses struct {
	Data    AddressesInfo  `json:"data"`
	Context ContextAddress `json:"context"`
}

// DataXpub includes full server response to xpub request.
type DataXpub struct {
	Data    map[string]XpubInfo `json:"data"`
	Context ContextAddress      `json:"context"`
}

// XpubInfo describes the outer structure of the x/z/v-pub.
type XpubInfo struct {
	Xpub         Xpub               `json:"xpub"`
	Addresses    map[string]Address `json:"addresses"`
	Transactions []string           `json:"transactions"`
	Utxo         []Utxo             `json:"utxo"`
}

// Xpub describes the inner structure of the x/z/v-pub.
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

// AddressesInfo describes the outer structure of the addresses.
type AddressesInfo struct {
	Set          Set                `json:"set"`
	Addresses    map[string]Address `json:"addresses"`
	Transactions []string           `json:"transactions"`
	Utxo         []Utxo             `json:"utxo"`
}

// AddressInfoEth the structure of the set of address and calls for Ethereum.
type AddressInfoEth struct {
	Address AddressEth     `json:"address"`
	Calls   []CallsAddress `json:"calls"`
}

// Set the structure of the set for Bitcoin-like address.
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

// AddressInfo structure of the set of address, transactions, and utxo.
type AddressInfo struct {
	Address      Address  `json:"address"`
	Transactions []string `json:"transactions"`
	Utxo         []Utxo   `json:"utxo"`
}

// AddressEth is the structure of one specific Ethereum address.
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

// Address the structure of one specific Bitcoin-like address.
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

// Utxo the structure of utxo.
type Utxo struct {
	BlockID         int    `json:"block_id"`
	TransactionHash string `json:"transaction_hash"`
	Index           int    `json:"index"`
	Value           int    `json:"value"`
}

// CallsAddress is the structures of calls.
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
// the structure of context for address(es).
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
	API            *API     `json:"api"`
	Server         string   `json:"server"`
	Time           float32  `json:"time"`
	RenderTime     float32  `json:"render_time"`
	FullTime       float32  `json:"full_time"`
	RequestCost    float32  `json:"request_cost"`
}

// GetAddress get the address by type of crypto and address hash.
func (c *Client) GetAddress(crypto string, address string) (*DataAddress, error) {
	return c.GetAddressAdv(crypto, address, nil)
}

// GetAddressAdv get the address by type of crypto, address hash, and options.
func (c *Client) GetAddressAdv(crypto string, address string, options map[string]string) (resp *DataAddress, e error) {
	if !Contains(GetSupportedCrypto(), crypto) {
		log.Fatalf("error: %v is not supported", crypto)
	}
	resp = &DataAddress{}
	var path = crypto + "/dashboards/address/" + address
	return resp, c.LoadResponse(path, resp, options)
}

// GetAddresses get the addresses by type of crypto and addresses hash.
func (c *Client) GetAddresses(crypto string, addresses []string) (*DataAddresses, error) {
	return c.GetAddressesAdv(crypto, addresses, nil)
}

// GetAddressesAdv get the addresses by type of crypto, addresses hash and options.
func (c *Client) GetAddressesAdv(crypto string, addresses []string, options map[string]string) (resp *DataAddresses, e error) {
	if !Contains(GetSupportedCrypto(), crypto) {
		log.Fatalf("error: %v is not supported", crypto)
	}

	resp = &DataAddresses{}
	var path = crypto + "/dashboards/addresses/" + strings.Join(addresses, ",")
	return resp, c.LoadResponse(path, resp, options)
}

// GetXpub get the xpub/ypub/zpub by type of crypto and the extended key.
// xpub (supported for all blockchains), ypub (supported for Bitcoin, Litecoin, Groestlcoin, and Bitcoin Testnet only) and zpub (supported for Bitcoin, Litecoin, Groestlcoin, and Bitcoin Testnet only)
func (c *Client) GetXpub(crypto string, extendedKey string) (*DataXpub, error) {
	return c.GetXpubAdv(crypto, extendedKey, nil)
}

// GetXpubAdv get the xpub/ypub/zpub by type of crypto, the extended key, and options.
// xpub (supported for all blockchains), ypub (supported for Bitcoin, Litecoin, Groestlcoin, and Bitcoin Testnet only) and zpub (supported for Bitcoin, Litecoin, Groestlcoin, and Bitcoin Testnet only)
func (c *Client) GetXpubAdv(crypto string, extendedKey string, options map[string]string) (resp *DataXpub, e error) {

	resp = &DataXpub{}
	var path = crypto + "/dashboards/xpub/" + extendedKey
	return resp, c.LoadResponse(path, resp, options)
}

// GetAddressEth get the address by type of crypto and address hash for Ethereum.
func (c *Client) GetAddressEth(crypto string, address string) (*DataAddressEth, error) {
	return c.GetAddressEthAdv(crypto, address, nil)
}

// GetAddressEthAdv get the address by type of crypto, address hash, and options for Ethereum.
func (c *Client) GetAddressEthAdv(crypto string, address string, options map[string]string) (resp *DataAddressEth, e error) {
	if !Contains(GetSupportedCryptoEth(), crypto) {
		log.Fatalf("error: %v is not supported", crypto)
	}

	resp = &DataAddressEth{}
	var path = crypto + "/dashboards/address/" + address
	return resp, c.LoadResponse(path, resp, options)
}
