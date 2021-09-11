package blockchair

import (
	"encoding/json"
	"testing"
)

func TestGetAddress(t *testing.T) {
	tests := []struct {
		currency string
		address  string
	}{
		{"bitcoin", "34xp4vRoCGJym3xR7yCVPFHoCNxv4Twseo"},
		{"bitcoin", "bc1qgdjqv0av3q56jvd82tkdjpy7gdp9ut8tlqmgrpmv24sq90ecnvqqjwvw97"},
		{"bitcoin-cash", "qram0qvz9xsuqutqauufr36a07xtz3ssvue2w43c8m"},
		{"bitcoin-cash", "pq47a3s9exn9zt64l6f66an48cj0eptekq3vk6udg0"},
		{"litecoin", "MKLyUgmiWF6SPZJVzm1hbiiNgpTxgnRiDE"},
		{"litecoin", "MD4Q9gLQ7Cv1ZCqwFi33SezVFGomja8r9f"},
		{"bitcoin-sv", "s-3412d0b7305a8173d53edc596012dc1d"},
		{"bitcoin-sv", "s-166ee714859656ae26bc1c041d808804"},
		{"dogecoin", "DPEzPFx1YAg2AndcYXD9ouPiNT5izSgeHL"},
		{"dogecoin", "D9a1Ah7mUNAJwNqHPER4DN9zNLoqcYFDZW"},
		{"dash", "Xty4Q4B1CCm1qA4sMFkmczZqCtftFJuEse"},
		{"dash", "XpxyaeV8yABKekJMEEB8jhvybxHxHDANeV"},
		{"groestlcoin", "grs1q7ur04yyce3gp0vwz9897dmhem0mcgeh3hpgqad"},
		{"groestlcoin", "FnQZTzpY3c7BTrQ2SPDcFftxWjJRwhFDXQ"},
		{"zcash", "t1Mda6nGyo4RJCxYeyaUnKK9Pawn87tryZm"},
		{"zcash", "t1cgx4kHg49vYEp3MPWV3NeUPg8fpmuH5zR"},
		{"ecash", "ecash:qqv4ruhsv7psuaep04dkq5tpp35v22x9mv5akz9m9q"},
		{"ecash", "ecash:qqvcj5lutpms84wx6cxr8hxzt8z7d885luzhre8dsc"},
	}
	for _, test := range tests {
		t.Run(test.currency, func(t *testing.T) {
			cl := New()
			cl.APIKey = clientID
			_, e := cl.GetAddress(test.currency, test.address)
			if e != nil {
				t.Fatal(e)
			}
		})
	}
}

func TestGetAddresses(t *testing.T) {
	tests := []struct {
		currency string
		address  []string
	}{
		{"bitcoin", []string{"34xp4vRoCGJym3xR7yCVPFHoCNxv4Twseo", "bc1qgdjqv0av3q56jvd82tkdjpy7gdp9ut8tlqmgrpmv24sq90ecnvqqjwvw97"}},
		{"bitcoin-cash", []string{"qram0qvz9xsuqutqauufr36a07xtz3ssvue2w43c8m", "pq47a3s9exn9zt64l6f66an48cj0eptekq3vk6udg0"}},
	}
	for _, test := range tests {
		t.Run(test.currency, func(t *testing.T) {
			cl := New()
			cl.APIKey = clientID
			_, e := cl.GetAddresses(test.currency, test.address)
			if e != nil {
				t.Fatal(e)
			}
		})
	}
}

func TestGetAddressEth(t *testing.T) {
	tests := []struct {
		currency string
		address  string
	}{
		{"ethereum", "0x3282791d6fd713f1e94f4bfd565eaa78b3a0599d"},
		{"ethereum", "0x9b22a80D5c7B3374a05b446081f97d0A34079e7F"},
		{"ethereum", "0x3282791d6fd713f1e94f4bfd565eaa78b3a0599d"},
		{"ethereum", "0xea674fdde714fd979de3edf0f56aa9716b898ec8"},
	}
	for _, test := range tests {
		t.Run(test.currency, func(t *testing.T) {
			cl := New()
			cl.APIKey = clientID
			_, e := cl.GetAddressEth(test.currency, test.address)
			if e != nil {
				t.Fatal(e)
			}
		})
	}
}

func BenchmarkGetAddressUnmarshal(b *testing.B) {
	cl := New()
	cl.APIKey = clientID
	response, e := cl.GetAddress("bitcoin", "bc1qgdjqv0av3q56jvd82tkdjpy7gdp9ut8tlqmgrpmv24sq90ecnvqqjwvw97")
	if e != nil {
		b.Fatal(e)
	}

	bytes, e := json.Marshal(response)
	if e != nil {
		b.Fatal(e)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		e := json.Unmarshal(bytes, response)
		if e != nil {
			b.Fatal(e)
		}
	}
	b.StopTimer()
}
