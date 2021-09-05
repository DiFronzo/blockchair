package blockchair

import (
	"testing"
)

func TestGetUncle(t *testing.T) {
	tests := []struct {
		currency string
		hash  string
	}{
		{"ethereum", "0xedc7a92c2a8aa140b0afa26db4ce8e05994a67d6fc3d736ddd77210b0ba565bb"},
		{"ethereum", "0x5cd50096dbb856a6d1befa6de8f9c20decb299f375154427d90761dc0b101109"},
	}
	for _, test := range tests {
		t.Run(test.currency, func(t *testing.T) {
			cl, _ := New()
			_, e := cl.GetUncle(test.currency, test.hash)
			if e != nil {
				t.Fatal(e)
			}
		})
	}
}

func TestGetUncles(t *testing.T) {
	tests := []struct {
		currency string
		hashes  []string
	}{
		{"ethereum", []string{"0x6859655d37abc23206e694b4d1b863b1db36f9f0b4272fb75b6885e4e957a90f", "0x30922011012a49f61e83edd60ce169a35b327c86415ed8d4b6c1cb211efabfa8"}},
		{"ethereum", []string{"0x11fa6bd6278453be8678bd3efb3acab836c9830fc56750c1a01717ee8b1cfefb", "0x71fbcd5fb4aa3ec759349b273aa07b66be19b507884263bc20c2735dce03dd9f"}},
	}
	for _, test := range tests {
		t.Run(test.currency, func(t *testing.T) {
			cl, _ := New()
			_, e := cl.GetUncles(test.currency, test.hashes)
			if e != nil {
				t.Fatal(e)
			}
		})
	}
}