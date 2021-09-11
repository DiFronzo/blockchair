package blockchair

import (
	"strings"
)

// DataUncle includes full server response to uncle request.
type DataUncle struct {
	Data    map[string]UncleInfo `json:"data"`
	Context ContextUncle         `json:"context"`
}

// UncleInfo describes the outer structure of the uncle.
type UncleInfo struct {
	Uncle Uncle `json:"uncle"`
}

// Uncle is the structure of one specific uncle block.
type Uncle struct {
	ParentBlockID    int     `json:"parent_block_id"`
	Index            int     `json:"index"`
	ID               int     `json:"id"`
	Hash             string  `json:"hash"`
	Date             string  `json:"date"`
	Time             string  `json:"time"`
	Size             int     `json:"size"`
	Miner            string  `json:"miner"`
	ExtraDataHex     string  `json:"extra_data_hex"`
	Difficulty       float32 `json:"difficulty"`
	GasUsed          int     `json:"gas_used"`
	GasLimit         int     `json:"gas_limit"`
	BaseFeePerGas    float32 `json:"base_fee_per_gas,omitempty"`
	LogsBloom        string  `json:"logs_bloom"`
	MixHash          string  `json:"mix_hash"`
	Nonce            string  `json:"nonce"`
	ReceiptsRoot     string  `json:"receipts_root"`
	Sha3Uncles       string  `json:"sha3_uncles"`
	StateRoot        string  `json:"state_root"`
	TransactionsRoot string  `json:"transactions_root"`
	Generation       string  `json:"generation"`
	GenerationUsd    float32 `json:"generation_usd"`
}

// ContextUncle the structure of context for uncle(s).
type ContextUncle struct {
	Code           int     `json:"code"`
	Source         string  `json:"source"`
	Results        int     `json:"results"`
	State          int     `json:"state"`
	StateLayer2    int     `json:"state_layer_2"`
	MarketPriceUsd float32 `json:"market_price_usd"`
	Cache          *Cache  `json:"cache"`
	API            *API    `json:"api"`
	Server         string  `json:"server"`
	Time           float64 `json:"time"`
	RenderTime     float64 `json:"render_time"`
	FullTime       float64 `json:"full_time"`
	RequestCost    float32 `json:"request_cost"`
}

// GetUncle fetch an uncle block created on Ethereum.
func (c *Client) GetUncle(crypto string, hash string) (*DataUncle, error) {
	return c.GetUncleAdv(crypto, hash, nil)
}

// GetUncleAdv fetch an uncle block created on Ethereum with options.
func (c *Client) GetUncleAdv(crypto string, hash string, options map[string]string) (resp *DataUncle, e error) {
	if e = c.ValidateCryptoEth(crypto); e != nil {
		return
	}
	if e = c.ValidateHashEth(hash); e != nil {
		return
	}

	resp = &DataUncle{}
	var path = crypto + "/dashboards/uncle/" + hash
	return resp, c.LoadResponse(path, resp, options)
}

// GetUncles fetch multiple uncle blocks created on Ethereum.
func (c *Client) GetUncles(crypto string, hashes []string) (*DataUncle, error) {
	return c.GetUnclesAdv(crypto, hashes, nil)
}

// GetUnclesAdv fetch multiple uncle blocks created on Ethereum with options.
func (c *Client) GetUnclesAdv(crypto string, hashes []string, options map[string]string) (resp *DataUncle, e error) {
	if e = c.ValidateCryptoEth(crypto); e != nil {
		return
	}
	if e = c.ValidateHashesEth(hashes); e != nil {
		return
	}

	resp = &DataUncle{}
	var path = crypto + "/dashboards/uncles/" + strings.Join(hashes, ",")
	return resp, c.LoadResponse(path, resp, options)
}
