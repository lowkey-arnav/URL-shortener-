package store

import "testing"

func TestCacheKey(t *testing.T) { // 		what it says that if i give a code it shall return the key in the following format
	if got := cacheKey("abc123"); got != "url:abc123" {
		t.Errorf("cacheKey = %q, want %q", got, "url:abc123")
	}
}
