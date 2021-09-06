package blockchair

import (
	"encoding/json"
	"testing"
)

func TestGetTransaction(t *testing.T) {
	tests := []struct {
		currency string
		tx  string
	}{
		{"bitcoin", "ed232ffd13184a8ff682364d20d25575492fbc9f8904343308e2b68d71feda21"},
		{"bitcoin", "ef4e4240cdca472910ba4e6d77102320cb866d378964f3e663d69fc5fbfc0cd9"},
		{"bitcoin-cash", "263d4198ef627d582b8b96e0d4f58f8eea77e9e2393539afd30e696af1500c8e"},
		{"bitcoin-cash", "96fb4ea992bd6bf086d7ac8cee07ee6c261aaf65570b6bc64288b57debd0e13c"},
		{"litecoin", "dc4b9ce0713971202270ae8240ec206f34d1289dfba5b96dcf12bb98c4dc96d1"},
		{"litecoin", "5e8ccf8c5869e66d498c348aabeb564d330e846a806acbea00715e9005b3fc8d"},
		{"bitcoin-sv", "1fcb691ace1894b7f1dd556b0ba3070bd6b799795a2192d0dc22ec223c945b9a"},
		{"bitcoin-sv", "124ce1001e77bceabeb098fce32c27bc41b9ee9b6723a445ecddf8927549641b"},
		{"dogecoin", "863556186e0b0687a5ab329927fbd40bb558b3e232b9d80616d406da91e53822"},
		{"dogecoin", "0696eecd9f12940b08e9ac565fcdd7a58c513043457c51652bdbd051f30a45cb"},
		{"dash", "b1624961540cc2317bbd4b2e1ec01f3e80f5c06ca76cacdebd99a3eeae76f63a"},
		{"dash", "cf31db4b858758a508f1267181b6d4455658d6717e835bace2c566557cfe0a40"},
		{"groestlcoin", "42dcc9cf3b3321c370ec5100c9e0eba0cde46cba019c42a103d7643974e736db"},
		{"groestlcoin", "b9d5ccda8d3d597d53d1987df57e88fe9c51605c92e48c0fb932b1e5eb0ee1c8"},
		{"zcash", "f802ff6b839d2b84312935380ffa96fdf2029208e9a28c17dfa3f347f85497a4"},
		{"zcash", "22b782138b32eff09afdcddbaa31826696d14c3d093d30e0068781fe4a45af14"},
		{"ecash", "2e426ca796d6a4695ddaf6823671efd8f212685fd9d99e38db69ed8799658a12"},
		{"ecash", "19fbd9269149b0496390df5891458b09cc1147d35a3d60818e97bed3368c6a01"},
	}
	for _, test := range tests {
		t.Run(test.currency, func(t *testing.T) {
			cl, _ := New()
			_, e := cl.GetTransaction(test.currency, test.tx)
			if e != nil {
				t.Fatal(e)
			}
		})
	}
}

func TestGetTransactions(t *testing.T) {
	tests := []struct {
		currency string
		tx  []string
	}{
		{"bitcoin", []string{"ed232ffd13184a8ff682364d20d25575492fbc9f8904343308e2b68d71feda21", "ef4e4240cdca472910ba4e6d77102320cb866d378964f3e663d69fc5fbfc0cd9"}},
		{"bitcoin-cash", []string{"263d4198ef627d582b8b96e0d4f58f8eea77e9e2393539afd30e696af1500c8e", "96fb4ea992bd6bf086d7ac8cee07ee6c261aaf65570b6bc64288b57debd0e13c"}},
	}
	for _, test := range tests {
		t.Run(test.currency, func(t *testing.T) {
			cl, _ := New()
			_, e := cl.GetTransactions(test.currency, test.tx)
			if e != nil {
				t.Fatal(e)
			}
		})
	}
}

func TestGetTransactionEth(t *testing.T) {
	tests := []struct {
		currency string
		tx  string
	}{
		{"ethereum", "0xb8e130ee4f809e495c91e2bbc071e7f18ba7ec5a15b33b7ea6243c9ab8f89bee"},
		{"ethereum", "0x9c6a8ca391990a5da03f200a62e8338020b8783d94b21ba8beccbf81f6e6d1f3"},
		{"ethereum", "0x3e6db7f8a4f73114da2281a070e02bfa11412fb11fa512ab2e075d722498e932"},
		{"ethereum", "0x9d04edc6f05b8366613fa484ad16976b8b23aae73083d60dae81f95624619f12"},
	}
	for _, test := range tests {
		t.Run(test.currency, func(t *testing.T) {
			cl, _ := New()
			_, e := cl.GetTransactionEth(test.currency, test.tx)
			if e != nil {
				t.Fatal(e)
			}
		})
	}
}

func TestGetTransactionsEth(t *testing.T) {
	tests := []struct {
		currency string
		tx  []string
	}{
		{"ethereum", []string{"0xb8e130ee4f809e495c91e2bbc071e7f18ba7ec5a15b33b7ea6243c9ab8f89bee", "0x9c6a8ca391990a5da03f200a62e8338020b8783d94b21ba8beccbf81f6e6d1f3"}},
		{"ethereum", []string{"0x3e6db7f8a4f73114da2281a070e02bfa11412fb11fa512ab2e075d722498e932", "0x9d04edc6f05b8366613fa484ad16976b8b23aae73083d60dae81f95624619f12"}},
	}
	for _, test := range tests {
		t.Run(test.currency, func(t *testing.T) {
			cl, _ := New()
			_, e := cl.GetTransactionsEth(test.currency, test.tx)
			if e != nil {
				t.Fatal(e)
			}
		})
	}
}

func BenchmarkGetTransactionUnmarshal(b *testing.B) {
	cl, _ := New()
	response, e := cl.GetTransaction("bitcoin","ed232ffd13184a8ff682364d20d25575492fbc9f8904343308e2b68d71feda21")
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