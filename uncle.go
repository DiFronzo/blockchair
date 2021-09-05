package blockchair

import (
	"fmt"
	"log"
	"regexp"
	"strings"
)

type DataUncle struct {
	Data    map[string]UncleInfo    `json:"data"`
	Context ContextUncle `json:"context"`
}

type UncleInfo struct {
	Uncle	Uncle `json:"uncle"`
}

type Uncle struct {
	ParentBlockID    int         `json:"parent_block_id"`
	Index            int         `json:"index"`
	ID               int         `json:"id"`
	Hash             string      `json:"hash"`
	Date             string      `json:"date"`
	Time             string      `json:"time"`
	Size             int         `json:"size"`
	Miner            string      `json:"miner"`
	ExtraDataHex     string      `json:"extra_data_hex"`
	Difficulty       int64       `json:"difficulty"`
	GasUsed          int         `json:"gas_used"`
	GasLimit         int         `json:"gas_limit"`
	BaseFeePerGas    float32	 `json:"base_fee_per_gas"`
	LogsBloom        string      `json:"logs_bloom"`
	MixHash          string      `json:"mix_hash"`
	Nonce            string      `json:"nonce"`
	ReceiptsRoot     string      `json:"receipts_root"`
	Sha3Uncles       string      `json:"sha3_uncles"`
	StateRoot        string      `json:"state_root"`
	TransactionsRoot string      `json:"transactions_root"`
	Generation       string      `json:"generation"`
	GenerationUsd    float64     `json:"generation_usd"`
}

type ContextUncle struct {
	Code           int     `json:"code"`
	Source         string  `json:"source"`
	Results        int     `json:"results"`
	State          int     `json:"state"`
	StateLayer2    int     `json:"state_layer_2"`
	MarketPriceUsd float64 `json:"market_price_usd"`
	Cache          *Cache   `json:"cache"`
	API            *Api     `json:"api"`
	Server         string  `json:"server"`
	Time           float64 `json:"time"`
	RenderTime     float64 `json:"render_time"`
	FullTime       float64 `json:"full_time"`
	RequestCost    float32 `json:"request_cost"`
}

func (c *Client) GetUncle(crypto string, hash string) (*DataUncle, error) {
	if !Contains(GetSupportedCryptoEth(), crypto){
		log.Fatalf("error: %v is not supported",crypto)
	}
	r, _ := regexp.Compile(Hash)
	if !r.MatchString(hash) {
		log.Fatalf("error: %v is not a valid hash",hash)
	}

	rsp := &DataUncle{}

	var path = crypto + "/dashboards/uncle/" + hash
	e := c.loadResponse(path, rsp)

	if e != nil {
		fmt.Print(e)
	}
	return rsp, e
}

func (c *Client) GetUncles(crypto string, hashes []string) (*DataUncle, error) {
	if !Contains(GetSupportedCryptoEth(), crypto){
		log.Fatalf("error: %v is not supported",crypto)
	}
	r, _ := regexp.Compile(Hash)
	for i := range hashes{
		if !r.MatchString(hashes[i]) {
			log.Fatalf("error: %v is not a valid hash",hashes[i])
		}
	}
	rsp := &DataUncle{}
	var path = crypto + "/dashboards/uncles/" + strings.Join(hashes, ",")
	e := c.loadResponse(path, rsp)

	if e != nil {
		fmt.Print(e)
	}
	return rsp, e
}