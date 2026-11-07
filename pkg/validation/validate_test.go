package validation

import "testing"

func TestIsValid(t *testing.T) {
	cases := map[string]bool{
		"https://example.com": true,
		"example.com":         false,
		"":                    false,
	}
	for in, want := range cases {
		if got := IsValidURL(in); got != want {
			t.Errorf("isValidURL(%q) = %v, want %v", in, got, want)
		}
	}
}
