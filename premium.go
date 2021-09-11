package blockchair

// DataPremium includes full server response to premium stats request.
type DataPremium struct {
	Data    Premium        `json:"data"`
	Context ContextPremium `json:"context"`
}

// Premium describes the structure of the premium stats.
type Premium struct {
	ValidUntil            string `json:"valid_until"`
	MaxRequestsPerDay     uint32 `json:"max_requests_per_day,omitempty"`
	MaxRequestsInParallel uint32 `json:"max_requests_in_parallel,omitempty"`
	RequestsToday         uint32 `json:"requests_today"`
}

// ContextPremium the structure of context for premium stats.
type ContextPremium struct {
	Code        int     `json:"code"`
	Cache       *Cache  `json:"cache"`
	API         *API    `json:"api"`
	Time        float32 `json:"time.omitempty"`
	RenderTime  float64 `json:"render_time"`
	FullTime    float64 `json:"full_time"`
	RequestCost int     `json:"request_cost"`
}

// GetUsage get the API key usage.
// Needs c.APIKey to be set.
func (c *Client) GetUsage() (resp *DataPremium, e error) {

	resp = &DataPremium{}
	var path = "premium/stats"
	return resp, c.LoadResponse(path, resp, nil)
}
