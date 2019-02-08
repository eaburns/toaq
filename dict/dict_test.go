package dict

import "testing"

func TestWords(t *testing.T) {
	// This is hardly a test, but let's just make sure that something is there.
	if _, ok := Words["hoaq"]; !ok {
		t.Error(`Words["hoaq"] is missing`)
	}
}
