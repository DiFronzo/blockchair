package blockchair

// DataErc20 includes full server response to ERC-20 request.
type DataErc20 struct {
	Data    Erc20        `json:"data"`
	Context ContextUncle `json:"context"`
}

// DataErc20Holder includes full server response to ERC-20 holder request.
type DataErc20Holder struct {
	Data    map[string]Erc20HolderInfo `json:"data"`
	Context HolderContext              `json:"context"`
}

// Erc20HolderInfo describes the outer structure of a ERC-20 holder.
type Erc20HolderInfo struct {
	Address     HolderAddress       `json:"address"`
	Transaction []HolderTransaction `json:"transaction,omitempty"`
}

// HolderAddress the structure of one specific Ethereum address.
type HolderAddress struct {
	Balance                   string `json:"balance,omitempty"`
	BalanceApproximate        int    `json:"balance_approximate,omitempty"`
	Received                  string `json:"received,omitempty"`
	ReceivedApproximate       int    `json:"received_approximate,omitempty"`
	Spent                     string `json:"spent,omitempty"`
	SpentApproximate          int    `json:"spent_approximate,omitempty"`
	ReceivingTransactionCount int    `json:"receiving_transaction_count,omitempty"`
	SpendingTransactionCount  int    `json:"spending_transaction_count,omitempty"`
	TransactionCount          int    `json:"transaction_count,omitempty"`
	FirstSeenReceiving        string `json:"first_seen_receiving,omitempty"`
	LastSeenReceiving         string `json:"last_seen_receiving,omitempty"`
	FirstSeenSpending         string `json:"first_seen_spending,omitempty"`
	LastSeenSpending          string `json:"last_seen_spending,omitempty"`
}

// HolderTransaction the structure of one specific Ethereum transaction.
type HolderTransaction struct {
	BlockID          int    `json:"block_id"`
	ID               int    `json:"id"`
	TransactionHash  string `json:"transaction_hash"`
	Time             string `json:"time"`
	TokenAddress     string `json:"token_address"`
	TokenName        string `json:"token_name"`
	TokenSymbol      string `json:"token_symbol"`
	TokenDecimals    int    `json:"token_decimals"`
	Sender           string `json:"sender"`
	Recipient        string `json:"recipient"`
	Value            string `json:"value"`
	ValueApproximate int    `json:"value_approximate"`
}

// Erc20 is the structure of one specific ERC-20 token.
type Erc20 struct {
	Name                    string  `json:"name"`
	Symbol                  string  `json:"symbol"`
	Decimals                int     `json:"decimals"`
	Time                    string  `json:"time"`
	CreatingBlockID         int     `json:"creating_block_id"`
	CreatingTransactionHash string  `json:"creating_transaction_hash"`
	Transactions            int     `json:"transactions"`
	Transactions24H         int     `json:"transactions_24h"`
	Volume24HApproximate    float32 `json:"volume_24h_approximate"`
	Volume24H               string  `json:"volume_24h"`
	Circulation             string  `json:"circulation"`
	CirculationApproximate  float32 `json:"circulation_approximate"`
	MarketPriceUsd          float32 `json:"market_price_usd,omitempty"`
	MarketPriceBtc          float32 `json:"market_price_btc,omitempty"`
	MarketCapUsd            float32 `json:"market_cap_usd,omitempty"`
}

// HolderContext is the structure of context for ERC-20 holder.
type HolderContext struct {
	Code           int     `json:"code"`
	Source         string  `json:"source"`
	Limit          int     `json:"limit"`
	Offset         int     `json:"offset"`
	Results        int     `json:"results"`
	State          int     `json:"state"`
	StateLayer2    int     `json:"state_layer_2"`
	MarketPriceUsd float64 `json:"market_price_usd"`
	Cache          *Cache  `json:"cache"`
	API            *API    `json:"api"`
	Server         string  `json:"server"`
	Time           float32 `json:"time"`
	RenderTime     float32 `json:"render_time"`
	FullTime       float32 `json:"full_time"`
	RequestCost    float32 `json:"request_cost"`
}

// GetErc20 fetch some basic information on an ERC-20 token on Ethereum.
func (c *Client) GetErc20(crypto string, token string) (*DataErc20, error) {
	return c.GetErc20Adv(crypto, token, nil)
}

// GetErc20Adv fetch some basic information on an ERC-20 token on Ethereum with options.
func (c *Client) GetErc20Adv(crypto string, token string, options map[string]string) (resp *DataErc20, e error) {
	if e = c.ValidateCryptoEth(crypto); e != nil {
		return
	}
	if e = c.ValidateErc20Token(token); e != nil {
		return
	}

	resp = &DataErc20{}
	var path = crypto + "/erc-20/" + token + "/stats"
	return resp, c.LoadResponse(path, resp, options)
}

// GetErc20Holder fetch some basic information on an ERC-20 token holder on Ethereum.
func (c *Client) GetErc20Holder(crypto string, token string, address string) (*DataErc20Holder, error) {
	return c.GetErc20HolderAdv(crypto, token, address, nil)
}

// GetErc20HolderAdv fetch some basic information on an ERC-20 token holder on Ethereum with options.
func (c *Client) GetErc20HolderAdv(crypto string, token string, address string, options map[string]string) (resp *DataErc20Holder, e error) {
	if e = c.ValidateCryptoEth(crypto); e != nil {
		return
	}
	if e = c.ValidateErc20Tokens([]string{token, address}); e != nil {
		return
	}

	resp = &DataErc20Holder{}
	var path = crypto + "/erc-20/" + token + "/dashboards/address/" + address
	return resp, c.LoadResponse(path, resp, options)
}
