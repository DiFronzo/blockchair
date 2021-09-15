package blockchair

import "strconv"

// DataOmni includes full server response to the Omni Layer property request.
type DataOmni struct {
	Data    OmniInfo    `json:"data"`
	Context ContextOmni `json:"context"`
}

// OmniInfo describes the outer structure of the Omni Layer property.
type OmniInfo struct {
	ID                      int     `json:"id"`
	Name                    string  `json:"name"`
	Category                string  `json:"category"`
	Subcategory             string  `json:"subcategory"`
	Description             string  `json:"description"`
	URL                     string  `json:"url"`
	IsDivisible             bool    `json:"is_divisible"`
	Issuer                  string  `json:"issuer"`
	CreationTransactionHash string  `json:"creation_transaction_hash,omitempty"`
	CreationTime            string  `json:"creation_time,omitempty"`
	CreationBlockID         int     `json:"creation_block_id"`
	IsIssuanceFixed         bool    `json:"is_issuance_fixed"`
	IsIssuanceManaged       bool    `json:"is_issuance_managed"`
	Circulation             float32 `json:"circulation"`
	Ecosystem               int     `json:"ecosystem"`
}

// ContextOmni for omni.
type ContextOmni struct {
	Code           int     `json:"code"`
	Source         string  `json:"source"`
	Results        int     `json:"results"`
	State          int     `json:"state"`
	MarketPriceUsd int     `json:"market_price_usd"`
	Cache          *Cache  `json:"cache"`
	API            *API    `json:"api"`
	Server         string  `json:"server"`
	Time           float32 `json:"time"`
	RenderTime     float32 `json:"render_time"`
	FullTime       float32 `json:"full_time"`
	RequestCost    float32 `json:"request_cost"`
}

// GetOmni fetch some basic information on an Omni Layer (Bitcoin) property (token)-
func (c *Client) GetOmni(prorertyID int64) (*DataOmni, error) {
	return c.GetOmniAdv(prorertyID, nil)
}

// GetOmniAdv fetch some basic information on an Omni Layer (Bitcoin) property (token) with options.
func (c *Client) GetOmniAdv(prorertyID int64, options map[string]string) (resp *DataOmni, e error) {

	resp = &DataOmni{}
	var path = "bitcoin" + "/omni/dashboards/property/" + strconv.FormatInt(prorertyID, 10)
	return resp, c.LoadResponse(path, resp, options)
}
