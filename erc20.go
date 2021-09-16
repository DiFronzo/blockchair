package blockchair

// DataErc20 includes full server response to ERC-20 request.
type DataErc20 struct {
	Data    Erc20        `json:"data"`
	Context ContextUncle `json:"context"`
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

// GetErc20 fetch some basic information on an ERC-20 token on Ethereum.
func (c *Client) GetErc20(crypto string, token string) (*DataErc20, error) {
	return c.GetErc20Adv(crypto, token, nil)
}

// GetErc20Adv fetch some basic information on an ERC-20 token on Ethereum. with options.
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
