package blockchair

import (
	"fmt"
	"strings"
)

type DataBlock struct {
	Data       map[string]DataInfo `json:"data"`
	Context    Context    `json:"context"`
}

type DataBlockEth struct {
	Data       map[string]DataInfoEth `json:"data"`
	Context    Context    `json:"context"`
}

type DataBlocks struct {
	Data       map[string]DataInfo `json:"data"`
	Context    Context    `json:"context"`
}

type DataInfoEth struct {
	Block       	BlockEth   `json:"block"`
	Uncles	[]string `json:"uncles"` //??
	Transactions	[]string `json:"transactions"`
	SyntheticTransactions []int64 `json:"synthetic_transactions"`
}

type DataInfo struct {
	Block       	Block   `json:"block"`
	Transactions	[]string `json:"transactions"`
}

type Block struct {
	ID       	  	int    `json:"id"`
	Hash          	string `json:"hash"`
	Date          	string `json:"date"`
	Time  		  	string `json:"time"`
	MedianTime    	string `json:"median_time"`
	Size		  	int    `json:"size"`
	StrippedSize  	int    `json:"stripped_size"`
	Weight		  	int	 `json:"weight"`
	Version		  	int	 `json:"version"`
	VersionHex	  	string `json:"version_hex"`
	VersionBits	  	string `json:"version_bits"`
	MerkleRoot	  	string `json:"merkle_root"`
	Nonce		  	int	 `json:"nonce"`
	Bits		  	int	 `json:"bits"`
	Difficulty	  	float32	 `json:"difficulty"`
	Chainwork	    string `json:"chainwork"`
	CoinbaseDataHex string `json:"coinbase_data_hex"`
	TransactionCount int `json:"transaction_count"`
	WitnessCount int `json:"witness_count"`
	InputCount int `json:"input_count"`
	OutputCount int `json:"output_count"`
	InputTotal int `json:"input_total"`
	InputTotalUsd float32 `json:"input_total_usd"`
	OutputTotal int `json:"output_total"`
	OutputTotalUsd float32 `json:"output_total_usd"`
	FeeTotal int `json:"fee_total"`
	FeeTotalUsd float32 `json:"fee_total_usd"`
	FeePerKb float32 `json:"fee_per_kb"`
	FeePerKbUsd float32 `json:"fee_per_kb_usd"`
	FeePerKwu float32 `json:"fee_per_kwu"`
	FeePerKwuUsd float32 `json:"fee_per_kwu_usd"`
	CddTotal float32 `json:"cdd_total"`
	Generation int `json:"generation"`
	GenerationUsd float32 `json:"generation_usd"`
	Reward int `json:"reward"`
	RewardUsd float32  `json:"reward_usd"`
	GuessedMiner string `json:"guessed_miner"`
}

type BlockEth struct {
	ID                        int         `json:"id"`
	Hash                      string      `json:"hash"`
	Date                      string      `json:"date"`
	Time                      string      `json:"time"`
	Size                      int         `json:"size"`
	Miner                     string      `json:"miner"`
	ExtraDataHex              string      `json:"extra_data_hex"`
	Difficulty                int64       `json:"difficulty"`
	GasUsed                   int         `json:"gas_used"`
	GasLimit                  int         `json:"gas_limit"`
	BaseFeePerGas             interface{} `json:"base_fee_per_gas"`
	LogsBloom                 string      `json:"logs_bloom"`
	MixHash                   string      `json:"mix_hash"`
	Nonce                     string      `json:"nonce"`
	ReceiptsRoot              string      `json:"receipts_root"`
	Sha3Uncles                string      `json:"sha3_uncles"`
	StateRoot                 string      `json:"state_root"`
	TotalDifficulty           string      `json:"total_difficulty"`
	TransactionsRoot          string      `json:"transactions_root"`
	UncleCount                int         `json:"uncle_count"`
	TransactionCount          int         `json:"transaction_count"`
	SyntheticTransactionCount int         `json:"synthetic_transaction_count"`
	CallCount                 int         `json:"call_count"`
	SyntheticCallCount        int         `json:"synthetic_call_count"`
	ValueTotal                string      `json:"value_total"`
	ValueTotalUsd             float32     `json:"value_total_usd"`
	InternalValueTotal        string      `json:"internal_value_total"`
	InternalValueTotalUsd     float32     `json:"internal_value_total_usd"`
	Generation                string      `json:"generation"`
	GenerationUsd             float32     `json:"generation_usd"`
	UncleGeneration           string      `json:"uncle_generation"`
	UncleGenerationUsd        int         `json:"uncle_generation_usd"`
	FeeTotal                  string      `json:"fee_total"`
	FeeTotalUsd               float32     `json:"fee_total_usd"`
	BurnedTotal               interface{} `json:"burned_total"`
	Reward                    string      `json:"reward"`
	RewardUsd                 float32     `json:"reward_usd"`
}

type Context struct {
	Code       	  	int    `json:"code"`
	Source			string `json:"source"`
	Limit       	int    `json:"limit"`
	Offset       	int    `json:"offset"`
	Results			int    `json:"results"`
	State           int    `json:"state"`
	MarketPriceUsd  float64      `json:"market_price_usd"`
	Cache			*Cache  `json:"cache"`
	API				*Api  `json:"api"`
	Server			string `json:"server"`
	Time			float32    `json:"time"`
	RenderTime		float32    `json:"render_time"`
	FullTime		float32    `json:"full_time"`
	RequestCost		float32   	`json:"request_cost"`
}

type Cache struct {
	Live			bool   `json:"live"`
	Duration	    int    `json:"duration"`
	Since			string `json:"since"`
	Until			string `json:"until"`
	Time			float32 `json:"time"`
}

type Api struct {
	Version		  	string `json:"version"`
	LastMajorUpdate	string `json:"last_major_update"`
	NextMajorUpdate	string `json:"next_major_update"`
	Documentation	string `json:"documentation"`
	Notice			string `json:"notice"`
}


func (c *Client) GetBlock(crypto string, blockID string) (*DataBlock, error) {
	rsp := &DataBlock{}

	var path = crypto + "/dashboards/block/" + blockID
	e := c.loadResponse(path, rsp)

	if e != nil {
		fmt.Print(e)
	}
	return rsp, e
}

func (c *Client) GetBlocks(crypto string, blockIDs []string) (*DataBlocks, error) {
	rsp := &DataBlocks{}
	var path = crypto + "/dashboards/blocks/" + strings.Join(blockIDs, ",")
	e := c.loadResponse(path, rsp)

	if e != nil {
		fmt.Print(e)
	}
	return rsp, e
}

func (c *Client) GetBlockEth(crypto string, blockID string) (*DataBlockEth, error) {
	rsp := &DataBlockEth{}

	var path = crypto + "/dashboards/block/" + blockID
	e := c.loadResponse(path, rsp)

	if e != nil {
		fmt.Print(e)
	}
	return rsp, e
}

func (c *Client) GetBlocksEth(crypto string, blockIDs []string) (*DataBlockEth, error) {
	rsp := &DataBlockEth{}
	var path = crypto + "/dashboards/blocks/" + strings.Join(blockIDs, ",")
	e := c.loadResponse(path, rsp)

	if e != nil {
		fmt.Print(e)
	}
	return rsp, e
}

