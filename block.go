package blockchair

import (
	"fmt"
	"strings"
)

type DataBlock struct {
	Data       *DataInfo `json:"data"`
	Context    *Context    `json:"context"`
}

type DataBlocks struct {
	Data       []*DataInfo `json:"data"`
	Context    *Context    `json:"context"`
}

type DataInfo struct {
	Block       	*Block   `json:"block"`
	Transactions	[]string `json:"transactions"`
}

type Block struct {
	Id       	  	int    `json:"id"`
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
	Difficulty	  	int	 `json:"difficulty"`
	Chainwork	    string `json:"chainwork"`
	CoinbaseDataHex string `json:"coinbase_data_hex"`
	TransactionCount int `json:"transaction_count"`
	WitnessCount int `json:"witness_count"`
	InputCount int `json:"input_count"`
	OutputCount int `json:"output_count"`
	InputTotal int `json:"input_total"`
	InputTotalUsd int `json:"input_total_usd"`
	OutputTotal int `json:"output_total"`
	OutputTotalUsd int `json:"output_total_usd"`
	FeeTotal int `json:"fee_total"`
	FeeTotalUsd int `json:"fee_total_usd"`
	FeePerKb int `json:"fee_per_kb"`
	FeePerKbUsd int `json:"fee_per_kb_usd"`
	FeePerKwu int `json:"fee_per_kwu"`
	FeePerKwuUsd int `json:"fee_per_kwu_usd"`
	CddTotal int `json:"cdd_total"`
	Generation int `json:"generation"`
	GenerationUsd int `json:"generation_usd"`
	Reward int `json:"reward"`
	RewardUsd int  `json:"reward_usd"`
	GuessedMiner string `json:"guessed_miner"`
}

type Context struct {
	Code       	  	int    `json:"code"`
	Source			string `json:"source"`
	Limit       	int    `json:"limit"`
	Offset       	int    `json:"offset"`
	Results			int    `json:"results"`
	State           int    `json:"state"`
	MarketPriceUsd  int      `json:"market_price_usd"`
	Cache			*Cache  `json:"cache"`
	Api				*Api  `json:"api"`
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

