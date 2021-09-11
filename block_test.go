package blockchair

import (
	"encoding/json"
	"testing"
)

func TestGetBlock(t *testing.T) {
	tests := []struct {
		currency string
		address  string
	}{
		{"bitcoin", "00000000000000000003c6c4a4b2e5ed0af0a11036eaf1726dd6af8ee9b8d86e"},
		{"bitcoin", "0000000000000000000325cfc4640cf0f2bb0842e139dca758afa4d740ee27c1"},
		{"bitcoin-cash", "0000000000000000045f54cbf86e87a1dc043d74eba40978788dd3c938c0e0e9"},
		{"bitcoin-cash", "000000000000000002fa05038b3972e9f36b7ac4f8ce324348a5ce43813414d5"},
		{"litecoin", "ea55f38d14dbd878f99d9d981630a62a8108ed97df4c3bd94a6d2ae5bb28d57d"},
		{"litecoin", "410182837492d6705e0f2eef50c3d4dc197d240952a317e5ba26afdfb94e9f52"},
		{"bitcoin-sv", "00000000000000000abf435db75d4a52becdaccd4635d03ffbde40a929b584c4"},
		{"bitcoin-sv", "00000000000000000c2a7f12215b338a253d3cdd6dbff8830d969f212570ede4"},
		{"dogecoin", "fadc302877f6660a98ec6e92201fad1034c861225cafb067035be91369359679"},
		{"dogecoin", "c195b67e9874e592800cd0761717d0cd202c1427c3988bbbfe2758bcb9ce9933"},
		{"dash", "0000000000000006acf56c629320671659f9eac03272c81d1b4a2cc882366944"},
		{"dash", "000000000000000c81a6593b47b70be27f2d3069a1a9ee5933b43f713dfedb32"},
		{"groestlcoin", "00000000000001a0587808821600834659a251ab5e157226b04d45bd10b26e91"},
		{"groestlcoin", "0000000000000239315628df2b21a7dbe80e6becced5e51b53688cba5ce6daa5"},
		{"zcash", "00000000023ea958be4fddea5944fa6ac12e3872d95b80f2f157e3cb77417b39"},
		{"zcash", "0000000001453dd79f0fecd492a435ee0390172463da5fd4b46f417913607f50"},
		{"ecash", "000000000000000006fa41c56336bc0590f602564a55471cf97109fd384a4777"},
		{"ecash", "000000000000000001c4a7baf3ed47f614b15809fe0af13582d18b64df8e2e20"},
	}
	for _, test := range tests {
		t.Run(test.currency, func(t *testing.T) {
			cl := New()
			cl.APIKey = clientID
			_, e := cl.GetBlock(test.currency, test.address)
			if e != nil {
				t.Fatal(e)
			}
		})
	}
}

func TestGetBlockEth(t *testing.T) {
	tests := []struct {
		currency string
		address  string
	}{
		{"ethereum", "0x398bb2ed03d3e8bdd93ba782c2c53f6f5dd5eb9c2f75da74b360c46f4a42cfb4"},
		{"ethereum", "0xf0e12e5f3933dc91fda83fc6b1f1d7eb63f533994829fdf85f06ed4ba6ed42e0"},
		{"ethereum", "0x411c2474183f1580fc32d09f2149265f786c1663312061dab514cf997c4e1cfd"},
		{"ethereum", "0x6518d5358ffbf5f641f39d325564dacb749d1b325ac641e6a42e0d277473fda0"},
	}
	for _, test := range tests {
		t.Run(test.currency, func(t *testing.T) {
			cl := New()
			cl.APIKey = clientID
			_, e := cl.GetBlockEth(test.currency, test.address)
			if e != nil {
				t.Fatal(e)
			}
		})
	}
}

func TestGetBlocks(t *testing.T) {
	tests := []struct {
		currency string
		address  []string
	}{
		{"bitcoin", []string{"00000000000000000003c6c4a4b2e5ed0af0a11036eaf1726dd6af8ee9b8d86e", "0000000000000000000325cfc4640cf0f2bb0842e139dca758afa4d740ee27c1"}},
		{"bitcoin-cash", []string{"0000000000000000045f54cbf86e87a1dc043d74eba40978788dd3c938c0e0e9", "000000000000000002fa05038b3972e9f36b7ac4f8ce324348a5ce43813414d5"}},
	}
	for _, test := range tests {
		t.Run(test.currency, func(t *testing.T) {
			cl := New()
			cl.APIKey = clientID
			_, e := cl.GetBlocks(test.currency, test.address)
			if e != nil {
				t.Fatal(e)
			}
		})
	}
}

func TestGetBlocksEth(t *testing.T) {
	tests := []struct {
		currency string
		address  []string
	}{
		{"ethereum", []string{"0xf0e12e5f3933dc91fda83fc6b1f1d7eb63f533994829fdf85f06ed4ba6ed42e0", "0x398bb2ed03d3e8bdd93ba782c2c53f6f5dd5eb9c2f75da74b360c46f4a42cfb4"}},
		{"ethereum", []string{"0x6518d5358ffbf5f641f39d325564dacb749d1b325ac641e6a42e0d277473fda0", "0x411c2474183f1580fc32d09f2149265f786c1663312061dab514cf997c4e1cfd"}},
	}
	for _, test := range tests {
		t.Run(test.currency, func(t *testing.T) {
			cl := New()
			cl.APIKey = clientID
			_, e := cl.GetBlocksEth(test.currency, test.address)
			if e != nil {
				t.Fatal(e)
			}
		})
	}
}

func BenchmarkGetBlockUnmarshal(b *testing.B) {
	cl := New()
	cl.APIKey = clientID
	response, e := cl.GetBlock("bitcoin", "0000000000000000000325cfc4640cf0f2bb0842e139dca758afa4d740ee27c1")
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
