package blockchair

import (
	"strconv"
	"testing"
)

func TestGetOmni(t *testing.T) {
	tests := []struct {
		currency int64
	}{
		{31},
		{544},
		{1},
	}
	for _, test := range tests {
		t.Run(strconv.FormatInt(test.currency, 10), func(t *testing.T) {
			cl := New()
			cl.APIKey = clientID
			_, e := cl.GetOmni(test.currency)
			if e != nil {
				t.Fatal(e)
			}
		})
	}
}
