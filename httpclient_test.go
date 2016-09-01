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

	goodCases := []string{"http://httpbin.org"}

	for _, c := range goodCases {
		_, err := get(c)
		if err != nil {
			t.Error("get() shouldn't have failed")
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

	goodCases := []string{"http://httpbin.org/post"}
	for _, c := range goodCases {
		_, err := post(c, "q=test")
		if err != nil {
			t.Error("post() shouldn't have failed")
		}
	}
}
