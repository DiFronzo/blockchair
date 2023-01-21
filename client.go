package blockchair

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const (
	clientVersion = "0.1.4"
	apiRoot       = "https://api.blockchair.com/"
	// Hash used to verify Ethereum hash. TODO: Replace with Deterministic Finite Automaton
	Hash = "^0x[0-9a-f]{64}$"
	// UserAgent request header that lets servers and network peers identify the application of the requesting user agent.
	UserAgent = "blockchair-api-golang-" + clientVersion
)

// Errors it is a set of errors returned when working with the package.
var (
	ErrTHW = errors.New("blockchair: transaction hash is wrong")
	ErrERC = errors.New("blockchair: ERC-20 token is wrong")
	ErrSC  = errors.New("blockchair: the Bitcoin-like cryptocurrency is not supported")
	ErrSCE = errors.New("blockchair: the Ethereum cryptocurrency is not supported")
	ErrSCG = errors.New("blockchair: the cryptocurrency is not supported")
	ErrCGD = errors.New("blockchair: cannot get data on url")
	ErrRPE = errors.New("blockchair: response parsing error")
	ErrIRS = errors.New("blockchair: incorrect response status")
	ErrRLR = errors.New("blockchair: error 402, rate limit reached for free tier")
	ErrMAX = errors.New("blockchair: the maximum number of addresses is 100")
	ErrETH = errors.New("blockchair: can only handle one Ethereum cryptocurrency address")
)

// GetSupportedCrypto List of supported Bitcoin-like crypto.
func GetSupportedCrypto() []string {
	return []string{"bitcoin", "bitcoin-cash", "litecoin", "dogecoin", "dash", "groestlcoin", "zcash", "ecash", "bitcoin/testnet"}
}

// GetSupportedCryptoEth List of supported Ethereum crypto.
func GetSupportedCryptoEth() []string {
	return []string{"ethereum/testnet", "ethereum"}
}

// GetSupportedCryptoMultichain List of supported crypto for multichain address check.
func GetSupportedCryptoMultichain() []string {
	return []string{"bitcoin", "bitcoin-cash", "litecoin", "dash", "groestlcoin", "zcash", "ethereum"}
}

// Client specifies the mechanism by which individual API requests are made.
type Client struct {
	client *http.Client

	APIKey    string // API access key.
	UserAgent string // Optional additional User-Agent fragment.

	RateLimitFunc func(RateLimit) // Func to call after response is returned in LoadResponse.
}

func (c *Client) userAgent() string {
	c.UserAgent = strings.TrimSpace(c.UserAgent)
	if c.UserAgent == "" {
		return UserAgent
	}

	return UserAgent + " " + c.UserAgent
}

// SetRateLimitFunc sets a Client instances' RateLimitFunc.
func SetRateLimitFunc(ratefunc func(rl RateLimit)) func(*Client) {
	return func(c *Client) { c.RateLimitFunc = ratefunc }
}

// RateLimitFunc is rate limiting strategy for the Client instance.
type RateLimitFunc func(RateLimit)

// RateLimit store values from calling Premium API.
type RateLimit struct {
	Limit     int
	Remaining int
	Period    int
}

var defaultRateLimitFunc = func(rl RateLimit) {}

// PercentageLeft returns the ratio of Remaining to Limit as a percentage
func (rl RateLimit) PercentageLeft() int {
	return rl.Remaining * 100 / rl.Limit
}

// WaitTime returns the time.Duration ratio of Period to Limit
func (rl RateLimit) WaitTime() time.Duration {
	return (time.Second * time.Duration(rl.Period)) / time.Duration(rl.Limit)
}

// WaitTimeRemaining returns the time.Duration ratio of Period to Remaining
func (rl RateLimit) WaitTimeRemaining() time.Duration {
	if rl.Remaining < 2 {
		return time.Second * time.Duration(rl.Period)
	}
	return (time.Second * time.Duration(rl.Period)) / time.Duration(rl.Remaining)
}

// RateLimitStrategySleep sets RateLimitFunc to sleep by WaitTimeRemaining
func (c *Client) RateLimitStrategySleep() {
	c.RateLimitFunc = func(rl RateLimit) {
		remaining := rl.WaitTimeRemaining()
		time.Sleep(remaining)
	}
}

// RateLimitStrategyConcurrent sleeps for WaitTime * parallelism when
// remaining is less than or equal to parallelism.
func (c *Client) RateLimitStrategyConcurrent(parallelism int) {
	c.RateLimitFunc = func(rl RateLimit) {
		if rl.Remaining <= parallelism {
			wait := rl.WaitTime() * time.Duration(parallelism)
			time.Sleep(wait)
		}
	}
}

var countRemaining = 30

// parseRate parses rate related headers from http response.
func parseRate(apikey string) RateLimit {
	var rlVal RateLimit
	// TODO: make it more useful, Blockchair has no rate limit in headers
	// for users with API key we can call the Premium API
	if apikey != "" {
		rlVal.Limit = 1
		rlVal.Remaining = 1
		rlVal.Period = 5
	} else {
		countRemaining = countRemaining - 1
		rlVal.Limit = 30
		rlVal.Remaining = countRemaining
		rlVal.Period = 30
	}

	return rlVal
}

// LoadResponse to send a client request, which is then converted to the passed type.
func (c *Client) LoadResponse(path string, i interface{}, options map[string]string) error {
	fullPath := apiRoot + path

	if c.APIKey != "" && options == nil {
		fullPath += "?key=" + c.APIKey
	}

	if options != nil {
		if c.APIKey != "" {
			options["key"] = c.APIKey
		}
		values := url.Values{}
		for k, v := range options {
			values.Set(k, v)
		}
		fullPath += "?" + (values.Encode())
	}

	req, e := http.NewRequest("GET", fullPath, nil)
	if e != nil {
		return c.err2(ErrCGD, e)
	}

	req.Header.Set("User-Agent", c.userAgent())

	// fmt.Println("Querry... ", fullPath)
	resp, e := c.client.Do(req)

	if e != nil {
		return c.err3(ErrCGD, e, resp)
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(resp.Body)

	rl := parseRate(c.APIKey)
	c.RateLimitFunc(rl)

	err := json.NewDecoder(resp.Body).Decode(&i)
	if err != nil {
		return c.err3(ErrRPE, err, resp)
	}

	if resp.Status[0] != '2' {
		if resp.Status == "402 Payment Required" {
			return c.err3(ErrRLR, e, resp)
		}
		return c.err3(ErrIRS, e, resp)
	}

	return nil
}

// formatRateReset formats d to look like "[rate reset in 2s]" or
// "[rate reset in 87m02s]" for the positive durations. And like "[rate limit was reset 87m02s ago]"
// for the negative cases.
// func formatRateReset(d time.Duration) string {
//	isNegative := d < 0
//	if isNegative {
//		d *= -1
//	}
//	secondsTotal := int(0.5 + d.Seconds())
//	minutes := secondsTotal / 60
//	seconds := secondsTotal - minutes*60

//	var timeString string
//	if minutes > 0 {
//		timeString = fmt.Sprintf("%dm%02ds", minutes, seconds)
//	} else {
//		timeString = fmt.Sprintf("%ds", seconds)
//	}

//	if isNegative {
//		return fmt.Sprintf("[rate limit was reset %v ago]", timeString)
//	}
//	return fmt.Sprintf("[rate reset in %v]", timeString)
// }

// New creates a new client instance the network internet.
func New() *Client {
	return &Client{client: &http.Client{}, RateLimitFunc: defaultRateLimitFunc}
}

// SetClient http client setter.
func (c *Client) SetClient(client *http.Client) {
	if client == nil {
		panic("blockchair: impossible install the client as nil")
	}
	c.client = client
}
