package swclient

import (
	"testing"
)

func TestGet(t *testing.T) {

	badCases := []string{"", "0.0.0.0", "http:///", "http", "://", "//////", "httpbin.org"}

	for _, c := range badCases {
		_, err := get(c)
		if err == nil {
			t.Error("get() should have failed due to uri being malformed")
		}
	}
}

func TestPost(t *testing.T) {

	badCases := []string{"", "0.0.0.0", "http:///", "http", "://", "//////", "httpbin.org"}

	for _, c := range badCases {
		_, err := post(c, "q=test")
		if err == nil {
			t.Error("post() should have failed due to uri being malformed")
		}
	}
}
