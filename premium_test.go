package blockchair

import "testing"

func TestClient_GetUsage(t *testing.T) {
	cl := New()
	cl.APIKey = clientID
	_, e := cl.GetUsage()
	if e != nil {
		t.Fatal(e)
	}
}
