package blockchair

import "testing"

func TestClient_CheckCryptoAddress(t *testing.T) {
	c := New()
	if e := c.ValidateCrypto("ethereum"); e.Error() != ErrSC.Error() {
		t.Fatal("incorrect error: " + e.Error())
	}
}

func TestClient_CheckCryptoAddressEth(t *testing.T) {
	c := New()
	if e := c.ValidateCryptoEth("bitcoin"); e.Error() != ErrSCE.Error() {
		t.Fatal("incorrect error: " + e.Error())
	}
}

func TestClient_CheckCryptoAddressBoth(t *testing.T) {
	c := New()
	if e := c.ValidateCryptoBoth("dollar"); e.Error() != ErrSCG.Error() {
		t.Fatal("incorrect error: " + e.Error())
	}
}

func TestValidateHashEth(t *testing.T) {
	t.Parallel()
	c := New()
	tests := []struct {
		address string
		result  bool
	}{
		// bad hash
		{"", false},
		{"1111111111111111111114oLvT", false},
		{"1111111111111111111114iLvT", false},
		{"0111111111111111111114oLvT2", false},
		{"xpub6DF8uhdarytz3FWdA8TvFSv", false},
		{"xpub3KGPnzYshia2uSSz8BED2kSpx22bbGCkzq", false},
	}
	for _, test := range tests {
		t.Run(test.address, func(t *testing.T) {
			if e := c.ValidateHashEth(test.address); e.Error() != ErrTHW.Error() {
				t.Fatal("incorrect error: " + e.Error())
			}
		})
	}
}

func TestValidateErc20Token(t *testing.T) {
	t.Parallel()
	c := New()
	tests := []struct {
		token  string
		result bool
	}{
		// bad token
		{"", false},
		{"0x411c2474183f1580fc32d09f2149265f786c1663312061dab514cf997c4e1cfd", false},
		{"0xf0e12e5f3933dc91fda83fc6b1f1d7eb63f533994829fdf85f06ed4ba6ed42e0", false},
		{"0111111111111111111114oLvT2", false},
		{"0xasdadasdasda23123dasdasde12d", false},
		{"xpub3KGPnzYshia2uSSz8BED2kSpx22bbGCkzq", false},
	}
	for _, test := range tests {
		t.Run(test.token, func(t *testing.T) {
			if e := c.ValidateErc20Token(test.token); e.Error() != ErrERC.Error() {
				t.Fatal("incorrect error: " + e.Error())
			}
		})
	}
}
