package swclient

import (
	"net/http"
	"testing"
)

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

func TestDigestParseParameters(t *testing.T) {

	testHeader := map[string][]string{}
	testHeader["Www-Authenticate"] = []string{`Digest username="user", realm="realm", nonce="nonce", uri="u/r/i", response="response", opaque="opaque", qop=auth, nc=00000001, cnonce="cnonce", algorithm="md5"`}
	testResponse := http.Response{Header: testHeader}

	cases := []http.Response{testResponse}

	for _, c := range cases {
		h := header{}
		h.parseParameters(c)

		if h.realm == "" {
			t.Error("realm is empty")
		}
		if h.qop == "" {
			t.Error("qop is empty")
		}
		if h.nOnce == "" {
			t.Error("nOnce is empty")
		}
		if h.opaque == "" {
			t.Error("opaque is empty")
		}
		if h.algorithm == "" {
			t.Error("algorithm is empty")
		}
	}
}

func TestHash(t *testing.T) {
	h, err := hash("test")
	if err != nil {
		t.Error(err)
	}

	if h != "098f6bcd4621d373cade4e832627b4f6" {
		t.Error("hash() calculated a wrong hash")
	}
}

func TestHashNow(t *testing.T) {
	_, err := hashNow()
	if err != nil {
		t.Error(err)
	}
}
