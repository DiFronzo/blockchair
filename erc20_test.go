package blockchair

import "testing"

func TestGetErc20(t *testing.T) {
	tests := []struct {
		currency string
		token    string
	}{
		{"ethereum", "0x4a73d94683f2c9c2Aaf32ccd5723F3e243D6a654"},
		{"ethereum", "0xdac17f958d2ee523a2206206994597c13d831ec7"},
		{"ethereum", "0x95aD61b0a150d79219dCF64E1E6Cc01f0B64C4cE"},
		{"ethereum", "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2"},
	}
	for _, test := range tests {
		t.Run(test.currency, func(t *testing.T) {
			cl := New()
			cl.APIKey = clientID
			_, e := cl.GetErc20(test.currency, test.token)
			if e != nil {
				t.Fatal(e)
			}
		})
	}
}

func TestGetErc20Holder(t *testing.T) {
	tests := []struct {
		currency string
		token    string
		address  string
	}{
		{"ethereum", "0x4a73d94683f2c9c2Aaf32ccd5723F3e243D6a654", "0x3282791d6fd713f1e94f4bfd565eaa78b3a0599d"},
		{"ethereum", "0x68e14bb5a45b9681327e16e528084b9d962c1a39", "0x3282791d6fd713f1e94f4bfd565eaa78b3a0599d"},
		{"ethereum", "0x7D1AfA7B718fb893dB30A3aBc0Cfc608AaCfeBB0", "0x3282791d6fd713f1e94f4bfd565eaa78b3a0599d"},
	}
	for _, test := range tests {
		t.Run(test.currency, func(t *testing.T) {
			cl := New()
			cl.APIKey = clientID
			_, e := cl.GetErc20Holder(test.currency, test.token, test.address)
			if e != nil {
				t.Fatal(e)
			}
		})
	}
}
