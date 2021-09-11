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
