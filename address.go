package blockchair

import (
	"encoding/json"
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

// DataMultichain includes full server response to multichain address check request.
type DataMultichain struct {
	Data    MultichainInfo `json:"data"`
	Context Context        `json:"context"`
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

// MultichainInfo describes the outer structure of the multichain address check.
type MultichainInfo struct {
	Set          SetMultichain                `json:"set"`
	Addresses    map[string]MultichainAddress `json:"addresses"`
	Transactions []TransactionsMultichain     `json:"transactions"`
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
	FirstSeenReceiving string  `json:"first_seen_receiving,omitempty"`
	LastSeenReceiving  string  `json:"last_seen_receiving,omitempty"`
	FirstSeenSpending  string  `json:"first_seen_spending,omitempty"`
	LastSeenSpending   string  `json:"last_seen_spending,omitempty"`
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
	BalanceUsd         float32 `json:"balance_usd"`
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

// SetMultichain the structure of the set for the multichain address check.
type SetMultichain struct {
	AddressCount       int     `json:"address_count"`
	BalanceUsd         float32 `json:"balance_usd"`
	ReceivedUsd        float32 `json:"received_usd"`
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
	ContractCodeHex     string  `json:"contract_code_hex,omitempty"`
	ContractCreated     string  `json:"contract_created,omitempty"`
	ContractDestroyed   string  `json:"contract_destroyed,omitempty"`
	Balance             string  `json:"balance"`
	BalanceUsd          float32 `json:"balance_usd"`
	ReceivedApproximate string  `json:"received_approximate"`
	ReceivedUsd         float32 `json:"received_usd"`
	SpentApproximate    string  `json:"spent_approximate"`
	SpentUsd            float32 `json:"spent_usd"`
	FeesApproximate     string  `json:"fees_approximate"`
	FeesUsd             float32 `json:"fees_usd"`
	ReceivingCallCount  int     `json:"receiving_call_count"`
	SpendingCallCount   int     `json:"spending_call_count"`
	CallCount           int     `json:"call_count"`
	TransactionCount    int     `json:"transaction_count"`
	FirstSeenReceiving  string  `json:"first_seen_receiving"`
	LastSeenReceiving   string  `json:"last_seen_receiving"`
	FirstSeenSpending   string  `json:"first_seen_spending,omitempty"`
	LastSeenSpending    string  `json:"last_seen_spending,omitempty"`
	Nonce               int     `json:"nonce,omitempty"`
}

// Address the structure of one specific Bitcoin-like address.
type Address struct {
	Path               string  `json:"path,omitempty"`
	Type               string  `json:"type"`
	ScriptHex          string  `json:"script_hex"`
	Balance            float32 `json:"balance"`
	BalanceUsd         float32 `json:"balance_usd"`
	Received           float32 `json:"received"`
	ReceivedUsd        float32 `json:"received_usd"`
	Spent              float32 `json:"spent"`
	SpentUsd           float32 `json:"spent_usd"`
	OutputCount        int     `json:"output_count"`
	UnspentOutputCount int     `json:"unspent_output_count"`
	FirstSeenReceiving string  `json:"first_seen_receiving,omitempty"`
	LastSeenReceiving  string  `json:"last_seen_receiving,omitempty"`
	FirstSeenSpending  string  `json:"first_seen_spending,omitempty"`
	LastSeenSpending   string  `json:"last_seen_spending,omitempty"`
	ScripthashType     string  `json:"scripthash_type,omitempty"`
	TransactionCount   int     `json:"transaction_count,omitempty"`
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
	TransactionHash string  `json:"transaction_hash,omitempty"`
	Index           string  `json:"index"`
	Time            string  `json:"time"`
	Sender          string  `json:"sender,omitempty"`
	Recipient       string  `json:"recipient"`
	Value           float64 `json:"value"`
	ValueUsd        float64 `json:"value_usd"`
	Transferred     bool    `json:"transferred"`
}

// MultichainAddress is the structures of multichain address for Bitcoin-like and Ethereum.
type MultichainAddress struct {
	Chain               string      `json:"chain"`
	Address             string      `json:"address"`
	Type                string      `json:"type"`
	ScriptHex           string      `json:"script_hex,omitempty"`
	ContractCodeHex     string      `json:"contract_code_hex,omitempty"`
	ContractCreated     string      `json:"contract_created,omitempty"`
	ContractDestroyed   string      `json:"contract_destroyed,omitempty"`
	Balance             json.Number `json:"balance"` //int64 for Bitcoin-like and string for Ethereum TODO!
	BalanceUsd          float32     `json:"balance_usd"`
	Received            float32     `json:"received,omitempty"`
	ReceivedApproximate string      `json:"received_approximate,omitempty"`
	ReceivedUsd         float32     `json:"received_usd"`
	Spent               float32     `json:"spent,omitempty"`
	SpentApproximate    string      `json:"spent_approximate,omitempty"`
	SpentUsd            float32     `json:"spent_usd"`
	OutputCount         int         `json:"output_count,omitempty"`
	UnspentOutputCount  int         `json:"unspent_output_count,omitempty"`
	FeesApproximate     string      `json:"fees_approximate,omitempty"`
	FeesUsd             float32     `json:"fees_usd,omitempty"`
	ReceivingCallCount  int         `json:"receiving_call_count,omitempty"`
	SpendingCallCount   int         `json:"spending_call_count,omitempty"`
	CallCount           int         `json:"call_count,omitempty"`
	TransactionCount    int         `json:"transaction_count,omitempty"`
	FirstSeenReceiving  string      `json:"first_seen_receiving"`
	LastSeenReceiving   string      `json:"last_seen_receiving"`
	FirstSeenSpending   string      `json:"first_seen_spending,omitempty"`
	LastSeenSpending    string      `json:"last_seen_spending,omitempty"`
	Nonce               int         `json:"nonce,omitempty"`
}

// TransactionsMultichain is the structure of one specific transaction.
type TransactionsMultichain struct {
	Chain         string  `json:"chain"`
	Address       string  `json:"address"`
	BlockID       int32   `json:"block_id"`
	Hash          string  `json:"hash"`
	Time          string  `json:"time"`
	BalanceChange float32 `json:"balance_change"`
}

// ContextAddress the structure of context for address(es).
// TODO! FIX "CHECKED" INTO A SLICE
type ContextAddress struct {
	Code           int      `json:"code"`
	Source         string   `json:"source"`
	Limit          string   `json:"limit"`
	Offset         string   `json:"offset"`
	Results        int      `json:"results"`
	Checked        []string `json:"checked,omitempty"`
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
	if e = c.ValidateCrypto(crypto); e != nil {
		return
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
	if e = c.ValidateCrypto(crypto); e != nil {
		return
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
	if e = c.ValidateCryptoBoth(crypto); e != nil {
		return
	}

	resp = &DataXpub{}
	var path = crypto + "/dashboards/xpub/" + extendedKey
	return resp, c.LoadResponse(path, resp, options)
}

// GetAddressEth get the address by type of crypto and address hash for Ethereum.
func (c *Client) GetAddressEth(crypto string, address string) (*DataAddressEth, error) {
	return c.GetAddressEthAdv(crypto, address, nil)
}

// GetAddressEthAdv get the address by type of crypto, address hash, and options for Ethereum.
// TODO! Validate address
func (c *Client) GetAddressEthAdv(crypto string, address string, options map[string]string) (resp *DataAddressEth, e error) {
	if e = c.ValidateCryptoEth(crypto); e != nil {
		return
	}

	resp = &DataAddressEth{}
	var path = crypto + "/dashboards/address/" + address
	return resp, c.LoadResponse(path, resp, options)
}

// MutliAddress struct for usage with GetMutlichainAddressCheck(Adv), sends type of crypto and address.
type MutliAddress []struct {
	currency, address string
}

// GetMutlichainAddressCheck check multiple addresses from different blockchain via just one request. This can be useful if you're monitoring your own wallet or portfolio.
func (c *Client) GetMutlichainAddressCheck(mutliAddress MutliAddress) (resp *DataMultichain, e error) {
	return c.GetMutlichainAddressCheckAdv(mutliAddress, nil)
}

// GetMutlichainAddressCheckAdv check multiple addresses from different blockchain via just one request. This can be useful if you're monitoring your own wallet or portfolio with options.
// TODO! CHECK THAT ONLY 1 OR NON ETH ADDR. IS ADDED. ALLOW USAGE OF "?transaction_details=true".
func (c *Client) GetMutlichainAddressCheckAdv(mutliAddress MutliAddress, options map[string]string) (resp *DataMultichain, e error) {
	if len(mutliAddress) > 100 { // max 100 addresses
		return nil, c.err1(ErrMAX)
	}
	var formatMultiAddr []string
	for j := range mutliAddress {
		if e = c.ValidateCryptoMultichain(mutliAddress[j].currency); e != nil {
			return
		}

		formatMultiAddr = append(formatMultiAddr, mutliAddress[j].currency+":"+mutliAddress[j].address)
	}

	resp = &DataMultichain{}
	var path = "multi/dashboards/addresses/" + strings.Join(formatMultiAddr, ",")
	return resp, c.LoadResponse(path, resp, options)
}
