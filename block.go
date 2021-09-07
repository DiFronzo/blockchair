package blockchair

import (
	"log"
	"strings"
)

// DataBlock includes full server response to block request.
type DataBlock struct {
	Data    map[string]DataInfo `json:"data"`
	Context Context             `json:"context"`
}

// DataBlockEth includes full server response to block request for Ethereum.
type DataBlockEth struct {
	Data    map[string]DataInfoEth `json:"data"`
	Context Context                `json:"context"`
}

// DataInfoEth describes the outer structure of the block for Ethereum.
type DataInfoEth struct {
	Block                 BlockEth `json:"block"`
	Uncles                []string `json:"uncles"`
	Transactions          []string `json:"transactions"`
	SyntheticTransactions []int64  `json:"synthetic_transactions"`
}

// DataInfo describes the outer structure of the block.
type DataInfo struct {
	Block        Block    `json:"block"`
	Transactions []string `json:"transactions"`
}

// Block the structure of one specific Bitcoin-like block.
type Block struct {
	ID               int     `json:"id"`
	Hash             string  `json:"hash"`
	Date             string  `json:"date"`
	Time             string  `json:"time"`
	MedianTime       string  `json:"median_time"`
	Size             int     `json:"size"`
	StrippedSize     int     `json:"stripped_size"`
	Weight           int     `json:"weight"`
	Version          int     `json:"version"`
	VersionHex       string  `json:"version_hex"`
	VersionBits      string  `json:"version_bits"`
	MerkleRoot       string  `json:"merkle_root"`
	Nonce            int     `json:"nonce"`
	Bits             int     `json:"bits"`
	Difficulty       float32 `json:"difficulty"`
	Chainwork        string  `json:"chainwork"`
	CoinbaseDataHex  string  `json:"coinbase_data_hex"`
	TransactionCount int     `json:"transaction_count"`
	WitnessCount     int     `json:"witness_count"`
	InputCount       int     `json:"input_count"`
	OutputCount      int     `json:"output_count"`
	InputTotal       int     `json:"input_total"`
	InputTotalUsd    float32 `json:"input_total_usd"`
	OutputTotal      int     `json:"output_total"`
	OutputTotalUsd   float32 `json:"output_total_usd"`
	FeeTotal         int     `json:"fee_total"`
	FeeTotalUsd      float32 `json:"fee_total_usd"`
	FeePerKb         float32 `json:"fee_per_kb"`
	FeePerKbUsd      float32 `json:"fee_per_kb_usd"`
	FeePerKwu        float32 `json:"fee_per_kwu"`
	FeePerKwuUsd     float32 `json:"fee_per_kwu_usd"`
	CddTotal         float32 `json:"cdd_total"`
	Generation       int     `json:"generation"`
	GenerationUsd    float32 `json:"generation_usd"`
	Reward           int     `json:"reward"`
	RewardUsd        float32 `json:"reward_usd"`
	GuessedMiner     string  `json:"guessed_miner"`
}

// BlockEth the structure of one specific Ethereum block.
type BlockEth struct {
	ID                        int     `json:"id"`
	Hash                      string  `json:"hash"`
	Date                      string  `json:"date"`
	Time                      string  `json:"time"`
	Size                      int     `json:"size"`
	Miner                     string  `json:"miner"`
	ExtraDataHex              string  `json:"extra_data_hex"`
	Difficulty                int64   `json:"difficulty"`
	GasUsed                   int     `json:"gas_used"`
	GasLimit                  int     `json:"gas_limit"`
	BaseFeePerGas             float64 `json:"base_fee_per_gas"`
	LogsBloom                 string  `json:"logs_bloom"`
	MixHash                   string  `json:"mix_hash"`
	Nonce                     string  `json:"nonce"`
	ReceiptsRoot              string  `json:"receipts_root"`
	Sha3Uncles                string  `json:"sha3_uncles"`
	StateRoot                 string  `json:"state_root"`
	TotalDifficulty           string  `json:"total_difficulty"`
	TransactionsRoot          string  `json:"transactions_root"`
	UncleCount                int     `json:"uncle_count"`
	TransactionCount          int     `json:"transaction_count"`
	SyntheticTransactionCount int     `json:"synthetic_transaction_count"`
	CallCount                 int     `json:"call_count"`
	SyntheticCallCount        int     `json:"synthetic_call_count"`
	ValueTotal                string  `json:"value_total"`
	ValueTotalUsd             float32 `json:"value_total_usd"`
	InternalValueTotal        string  `json:"internal_value_total"`
	InternalValueTotalUsd     float32 `json:"internal_value_total_usd"`
	Generation                string  `json:"generation"`
	GenerationUsd             float32 `json:"generation_usd"`
	UncleGeneration           string  `json:"uncle_generation"`
	UncleGenerationUsd        int     `json:"uncle_generation_usd"`
	FeeTotal                  string  `json:"fee_total"`
	FeeTotalUsd               float32 `json:"fee_total_usd"`
	BurnedTotal               float64 `json:"burned_total"`
	Reward                    string  `json:"reward"`
	RewardUsd                 float32 `json:"reward_usd"`
}

// Context common context for all requests
type Context struct {
	Code           int     `json:"code"`
	Source         string  `json:"source"`
	Limit          int     `json:"limit"`
	Offset         int     `json:"offset"`
	Results        int     `json:"results"`
	State          int     `json:"state"`
	MarketPriceUsd float64 `json:"market_price_usd"`
	Cache          *Cache  `json:"cache"`
	API            *API    `json:"api"`
	Server         string  `json:"server"`
	Time           float32 `json:"time"`
	RenderTime     float32 `json:"render_time"`
	FullTime       float32 `json:"full_time"`
	RequestCost    float32 `json:"request_cost"`
}

// Cache common cache for all requests
type Cache struct {
	Live     bool    `json:"live"`
	Duration int     `json:"duration"`
	Since    string  `json:"since"`
	Until    string  `json:"until"`
	Time     float32 `json:"time"`
}

// API common API for all requests
type API struct {
	Version         string `json:"version"`
	LastMajorUpdate string `json:"last_major_update"`
	NextMajorUpdate string `json:"next_major_update"`
	Documentation   string `json:"documentation"`
	Notice          string `json:"notice"`
}

// GetBlock fetch a Bitcoin-like block.
func (c *Client) GetBlock(crypto string, blockID string) (*DataBlock, error) {
	return c.GetBlockAdv(crypto, blockID, nil)
}

// GetBlockAdv fetch a Bitcoin-like block with options.
func (c *Client) GetBlockAdv(crypto string, blockID string, options map[string]string) (resp *DataBlock, e error) {
	if !Contains(GetSupportedCrypto(), crypto) {
		log.Fatalf("error: %v is not supported", crypto)
	}
	resp = &DataBlock{}

	var path = crypto + "/dashboards/block/" + blockID
	return resp, c.LoadResponse(path, resp, options)
}

// GetBlocks fetches multiple Bitcoin-like blocks.
func (c *Client) GetBlocks(crypto string, blockIDs []string) (*DataBlock, error) {
	return c.GetBlocksAdv(crypto, blockIDs, nil)
}

// GetBlocksAdv fetches multiple Bitcoin-like blocks with options
func (c *Client) GetBlocksAdv(crypto string, blockIDs []string, options map[string]string) (resp *DataBlock, e error) {
	if !Contains(GetSupportedCrypto(), crypto) {
		log.Fatalf("error: %v is not supported", crypto)
	}

	resp = &DataBlock{}
	var path = crypto + "/dashboards/blocks/" + strings.Join(blockIDs, ",")
	return resp, c.LoadResponse(path, resp, options)
}

// GetBlockEth fetch an Ethereum block.
func (c *Client) GetBlockEth(crypto string, blockID string) (*DataBlockEth, error) {
	return c.GetBlockEthAdv(crypto, blockID, nil)
}

// GetBlockEthAdv fetch an Ethereum block with options.
func (c *Client) GetBlockEthAdv(crypto string, blockID string, options map[string]string) (resp *DataBlockEth, e error) {
	if !Contains(GetSupportedCryptoEth(), crypto) {
		log.Fatalf("error: %v is not supported", crypto)
	}
	resp = &DataBlockEth{}

	var path = crypto + "/dashboards/block/" + blockID
	return resp, c.LoadResponse(path, resp, options)
}

// GetBlocksEth fetches multiple Ethereum blocks.
func (c *Client) GetBlocksEth(crypto string, blockIDs []string) (*DataBlockEth, error) {
	return c.GetBlocksEthAdv(crypto, blockIDs, nil)
}

// GetBlocksEthAdv fetches multiple Ethereum blocks with options.
func (c *Client) GetBlocksEthAdv(crypto string, blockIDs []string, options map[string]string) (resp *DataBlockEth, e error) {
	if !Contains(GetSupportedCryptoEth(), crypto) {
		log.Fatalf("error: %v is not supported", crypto)
	}
	resp = &DataBlockEth{}
	var path = crypto + "/dashboards/blocks/" + strings.Join(blockIDs, ",")
	return resp, c.LoadResponse(path, resp, options)
}
