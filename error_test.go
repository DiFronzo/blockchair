package blockchair

import (
	"net/http"
	"testing"
)

func TestNewError(t *testing.T) {
	if NewError(nil, nil, nil, nil) != nil {
		t.Fatal("Wrong error")
	}

	test := "test"
	resp := &http.Response{
		StatusCode: http.StatusOK,
	}
	err := NewError(ErrSC, ErrTHW, nil, nil)
	if err.Response != nil {
		t.Fatal("wrong Error.Response expected nil: ", *err.Response)
	}

	err = NewError(ErrSCE, ErrTHW, resp, &test)
	if *err.Hash != test {
		t.Fatal("wrong Error.Hash", *err.Hash)
	}
	if err.Response.StatusCode != http.StatusOK {
		t.Fatal("wrong Error.Response", *err.Response)
	}
}
