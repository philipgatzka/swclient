package swclient

import (
	"crypto/md5"
	"net/http"
	"testing"
)

func TestDigestParseParameters(t *testing.T) {
	testAuthHeader := `Digest username="user", realm="realm", nonce="nonce", uri="http://testing.org", response="response", opaque="opaque", qop=auth, nc=00000001, cnonce="cnonce", algorithm="md5"`

	testHeader := map[string][]string{}
	testHeader["Www-Authenticate"] = []string{testAuthHeader}
	testResponse := http.Response{Header: testHeader}

	cases := []http.Response{testResponse}

	h := header{}

	for _, c := range cases {
		tuples := h.parseParameters(c)

		if tuples["realm"] == "" {
			t.Error("realm is empty")
		}
		if tuples["qop"] == "" {
			t.Error("qop is empty")
		}
		if tuples["nonce"] == "" {
			t.Error("nOnce is empty")
		}
		if tuples["opaque"] == "" {
			t.Error("opaque is empty")
		}
		if tuples["algorithm"] == "" {
			t.Error("algorithm is empty")
		}
	}
}

func TestDigestIsComplete(t *testing.T) {
	h := header{}
	if h.isComplete() {
		t.Error("Expected h.isComplete() to be false")
	}

	h.realm = "somestring"
	h.algorithm = "somestring"
	h.aOne = "somestring"
	h.aTwo = "somestring"
	h.cNonce = "somestring"
	h.key = "somestring"
	h.method = "somestring"
	h.name = "somestring"
	h.nC = 0x1
	h.nOnce = "somestring"
	h.opaque = "somestring"
	h.path = "somestring"
	h.qop = "somestring"

	if !h.isComplete() {
		t.Error("Expected h.isComplete() to be true")
	}

	h.key = ""

	if h.isComplete() {
		t.Error("Expected h.isComplete() to be false")
	}
}

func TestDigestChecksums(t *testing.T) {
	h := header{}

	h.name = "hello"
	h.realm = "this is"
	h.key = "test"

	h.method = "test"
	h.path = "again"

	err := h.checksums(md5.New())
	if err != nil {
		t.Error(err)
	}

	expectedAOne := "5e6c31ea339e04180ef868d58c9e978b"
	expectedATwo := "b40a1f379821a5dfa8d4ed703bb10bcd"

	if h.aOne != expectedAOne {
		t.Error("aOne was incorrect!")
		t.Log("Got", h.aOne, " instead of", expectedAOne)
	}
	if h.aTwo != expectedATwo {
		t.Error("aTwo was incorrect!")
		t.Log("Got", h.aTwo, "instead of", expectedATwo)
	}
}

func TestDigestJoinWithColon(t *testing.T) {
	expected := "hello:this:is:test"
	got := joinWithColon([]string{"hello", "this", "is", "test"})

	if got != expected {
		t.Error("Expected", expected, "got", got)
	}
}
