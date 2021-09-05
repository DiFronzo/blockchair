package blockchair

import (
	"fmt"
	"strings"
)

type DataTransaction struct {
	Data	map[string]TransactionInfo	`json:"data"`
	Context    *Context    `json:"context"`
}

type TransactionInfo struct {
	Transaction		Transaction	`json:"transaction"`
	Inputs	[]interface{}	`json:"inputs"`
	Outputs []Outputs	`json:"outputs"`

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
	InputTotalUsd  float32     `json:"input_total_usd"`
	OutputTotal    int64   `json:"output_total"`
	OutputTotalUsd float32 `json:"output_total_usd"`
	Fee            int     `json:"fee"`
	FeeUsd         float32     `json:"fee_usd"`
	FeePerKb       float32     `json:"fee_per_kb"`
	FeePerKbUsd    float32     `json:"fee_per_kb_usd"`
	FeePerKwu      float32     `json:"fee_per_kwu"`
	FeePerKwuUsd   float32     `json:"fee_per_kwu_usd"`
	CddTotal       float32     `json:"cdd_total"`
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

func (c *Client) GetTransaction(crypto string, TxID string) (*DataTransaction, error) {
	rsp := &DataTransaction{}
	var path = crypto + "/dashboards/transaction/" + TxID
	e := c.loadResponse(path, rsp)

	if e != nil {
		fmt.Print(e)
	}
	return rsp, e
}

func (c *Client) GetTransactions(crypto string, TxIDs []string) (*DataTransaction, error) {
	rsp := &DataTransaction{}
	var path = crypto + "/dashboards/transactions/" + strings.Join(TxIDs, ",")
	e := c.loadResponse(path, rsp)

	if e != nil {
		fmt.Print(e)
	}
	return rsp, e
}