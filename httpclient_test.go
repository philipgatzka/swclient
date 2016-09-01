package swclient

import (
	"testing"
)

func TestGet(t *testing.T) {
	_, err := get("https://httpbin.org/digest-auth/auth/user/passwd")
	if err != nil {
		t.Error(err)
	}
}
