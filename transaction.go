package blockchair

import (
	"fmt"
	"log"
	"regexp"
	"strings"
)

type DataTransaction struct {
	Data    map[string]TransactionInfo `json:"data"`
	Context *Context                   `json:"context"`
}

type DataTransactionEth struct {
	Data    map[string]TransactionInfoEth `json:"data"`
	Context ContextEth                    `json:"context"`
}

type TransactionInfo struct {
	Transaction Transaction   `json:"transaction"`
	Inputs      []interface{} `json:"inputs"`
	Outputs     []Outputs     `json:"outputs"`
}

type TransactionInfoEth struct {
	Transaction TransactionEth `json:"transaction"`
	Calls       []Calls        `json:"calls"`
}

type TransactionEth struct {
	BlockID              int     `json:"block_id"`
	ID                   int64   `json:"id"`
	Index                int     `json:"index"`
	Hash                 string  `json:"hash"`
	Date                 string  `json:"date"`
	Time                 string  `json:"time"`
	Failed               bool    `json:"failed"`
	Type                 string  `json:"type"`
	Sender               string  `json:"sender"`
	Recipient            string  `json:"recipient"`
	CallCount            int     `json:"call_count"`
	Value                string  `json:"value"`
	ValueUsd             float64 `json:"value_usd"`
	InternalValue        string  `json:"internal_value"`
	InternalValueUsd     float64 `json:"internal_value_usd"`
	Fee                  string  `json:"fee"`
	FeeUsd               float64 `json:"fee_usd"`
	GasUsed              int     `json:"gas_used"`
	GasLimit             int     `json:"gas_limit"`
	GasPrice             int64   `json:"gas_price"`
	InputHex             string  `json:"input_hex"`
	Nonce                int     `json:"nonce"`
	V                    string  `json:"v"`
	R                    string  `json:"r"`
	S                    string  `json:"s"`
	Version              int     `json:"version"`
	EffectiveGasPrice    float32 `json:"effective_gas_price"`
	MaxFeePerGas         float32 `json:"max_fee_per_gas"`
	MaxPriorityFeePerGas float32 `json:"max_priority_fee_per_gas"`
	BaseFeePerGas        float32 `json:"base_fee_per_gas"`
	Burned               string  `json:"burned"`
	Type2718             bool    `json:"type_2718"` //??
}

type Calls struct {
	BlockID         int     `json:"block_id"`
	TransactionID   int64   `json:"transaction_id"`
	TransactionHash string  `json:"transaction_hash"`
	Index           string  `json:"index"`
	Depth           int     `json:"depth"`
	Date            string  `json:"date"`
	Time            string  `json:"time"`
	Failed          bool    `json:"failed"`
	FailReason      string  `json:"fail_reason"`
	Type            string  `json:"type"`
	Sender          string  `json:"sender"`
	Recipient       string  `json:"recipient"`
	ChildCallCount  int     `json:"child_call_count"`
	Value           string  `json:"value"`
	ValueUsd        float64 `json:"value_usd"`
	Transferred     bool    `json:"transferred"`
	InputHex        string  `json:"input_hex"`
	OutputHex       string  `json:"output_hex"`
}

type Transaction struct {
	BlockID        int     `json:"block_id"`
	ID             int     `json:"id"`
	Hash           string  `json:"hash"`
	Date           string  `json:"date"`
	Time           string  `json:"time"`
	Size           int     `json:"size"`
	Weight         int     `json:"weight"`
	Version        int     `json:"version"`
	LockTime       int     `json:"lock_time"`
	IsCoinbase     bool    `json:"is_coinbase"`
	HasWitness     bool    `json:"has_witness"`
	InputCount     int     `json:"input_count"`
	OutputCount    int     `json:"output_count"`
	InputTotal     int     `json:"input_total"`
	InputTotalUsd  float32 `json:"input_total_usd"`
	OutputTotal    int64   `json:"output_total"`
	OutputTotalUsd float32 `json:"output_total_usd"`
	Fee            int     `json:"fee"`
	FeeUsd         float32 `json:"fee_usd"`
	FeePerKb       float32 `json:"fee_per_kb"`
	FeePerKbUsd    float32 `json:"fee_per_kb_usd"`
	FeePerKwu      float32 `json:"fee_per_kwu"`
	FeePerKwuUsd   float32 `json:"fee_per_kwu_usd"`
	CddTotal       float32 `json:"cdd_total"`
	IsRbf          bool    `json:"is_rbf"`
}

type Inputs struct {
	BlockID                 int     `json:"block_id"`
	TransactionID           int     `json:"transaction_id"`
	Index                   int     `json:"index"`
	TransactionHash         string  `json:"transaction_hash"`
	Date                    string  `json:"date"`
	Time                    string  `json:"time"`
	Value                   int64   `json:"value"`
	ValueUsd                float32 `json:"value_usd"`
	Recipient               string  `json:"recipient"`
	Type                    string  `json:"type"`
	ScriptHex               string  `json:"script_hex"`
	IsFromCoinbase          bool    `json:"is_from_coinbase"`
	IsSpendable             bool    `json:"is_spendable"`
	IsSpent                 bool    `json:"is_spent"`
	SpendingBlockID         int     `json:"spending_block_id"`
	SpendingTransactionID   int     `json:"spending_transaction_id"`
	SpendingIndex           int     `json:"spending_index"`
	SpendingTransactionHash string  `json:"spending_transaction_hash"`
	SpendingDate            string  `json:"spending_date"`
	SpendingTime            string  `json:"spending_time"`
	SpendingValueUsd        float32 `json:"spending_value_usd"`
	SpendingSequence        int64   `json:"spending_sequence"`
	SpendingSignatureHex    string  `json:"spending_signature_hex"`
	SpendingWitness         string  `json:"spending_witness"`
	Lifespan                int     `json:"lifespan"`
	Cdd                     float32 `json:"cdd"`
}

type Outputs struct {
	BlockID                 int     `json:"block_id"`
	TransactionID           int     `json:"transaction_id"`
	Index                   int     `json:"index"`
	TransactionHash         string  `json:"transaction_hash"`
	Date                    string  `json:"date"`
	Time                    string  `json:"time"`
	Value                   int     `json:"value"`
	ValueUsd                float32 `json:"value_usd"`
	Recipient               string  `json:"recipient"`
	Type                    string  `json:"type"`
	ScriptHex               string  `json:"script_hex"`
	IsFromCoinbase          bool    `json:"is_from_coinbase"`
	IsSpendable             bool    `json:"is_spendable"`
	IsSpent                 bool    `json:"is_spent"`
	SpendingBlockID         int     `json:"spending_block_id"`
	SpendingTransactionID   int     `json:"spending_transaction_id"`
	SpendingIndex           int     `json:"spending_index"`
	SpendingTransactionHash string  `json:"spending_transaction_hash"`
	SpendingDate            string  `json:"spending_date"`
	SpendingTime            string  `json:"spending_time"`
	SpendingValueUsd        float32 `json:"spending_value_usd"`
	SpendingSequence        int64   `json:"spending_sequence"`
	SpendingSignatureHex    string  `json:"spending_signature_hex"`
	SpendingWitness         string  `json:"spending_witness"`
	Lifespan                int     `json:"lifespan"`
	Cdd                     float32 `json:"cdd"`
}

type ContextEth struct {
	Code           int     `json:"code"`
	Source         string  `json:"source"`
	Results        int     `json:"results"`
	State          int     `json:"state"`
	StateLayer2    int     `json:"state_layer_2"`
	MarketPriceUsd float64 `json:"market_price_usd"`
	Cache          *Cache  `json:"cache"`
	API            *Api    `json:"api"`
	Server         string  `json:"server"`
	Time           float64 `json:"time"`
	RenderTime     float64 `json:"render_time"`
	FullTime       float64 `json:"full_time"`
	RequestCost    float32 `json:"request_cost"`
}

func (c *Client) GetTransaction(crypto string, TxID string) (*DataTransaction, error) {
	if !Contains(GetSupportedCrypto(), crypto) {
		log.Fatalf("error: %v is not supported", crypto)
	}
	rsp := &DataTransaction{}
	var path = crypto + "/dashboards/transaction/" + TxID
	e := c.loadResponse(path, rsp)

	if e != nil {
		fmt.Print(e)
	}
	return rsp, e
}

func (c *Client) GetTransactionEth(crypto string, TxID string) (*DataTransactionEth, error) {
	if !Contains(GetSupportedCryptoEth(), crypto) {
		log.Fatalf("error: %v is not supported", crypto)
	}
	r, _ := regexp.Compile(Hash)
	if !r.MatchString(TxID) {
		log.Fatalf("error: %v is not a valid hash", TxID)
	}

	rsp := &DataTransactionEth{}
	var path = crypto + "/dashboards/transaction/" + TxID
	e := c.loadResponse(path, rsp)

	if e != nil {
		fmt.Print(e)
	}
	return rsp, e
}

func (c *Client) GetTransactions(crypto string, TxIDs []string) (*DataTransaction, error) {
	if !Contains(GetSupportedCrypto(), crypto) {
		log.Fatalf("error: %v is not supported", crypto)
	}
	rsp := &DataTransaction{}
	var path = crypto + "/dashboards/transactions/" + strings.Join(TxIDs, ",")
	e := c.loadResponse(path, rsp)

	if e != nil {
		fmt.Print(e)
	}
	return rsp, e
}

func (c *Client) GetTransactionsEth(crypto string, TxIDs []string) (*DataTransactionEth, error) {
	if !Contains(GetSupportedCryptoEth(), crypto) {
		log.Fatalf("error: %v is not supported", crypto)
	}
	r, _ := regexp.Compile(Hash)
	for i := range TxIDs {
		if !r.MatchString(TxIDs[i]) {
			log.Fatalf("error: %v is not a valid hash", TxIDs[i])
		}
	}

	rsp := &DataTransactionEth{}
	var path = crypto + "/dashboards/transactions/" + strings.Join(TxIDs, ",")
	e := c.loadResponse(path, rsp)

	if e != nil {
		fmt.Print(e)
	}
	return rsp, e
}
