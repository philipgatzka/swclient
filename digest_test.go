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
	testResponse := &http.Response{Header: testHeader}

	cases := []*http.Response{testResponse}

	for _, c := range cases {
		tuples := parseParameters(c)

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

func TestDigestJoinWithColon(t *testing.T) {
	expected := "hello:this:is:test"
	got := joinWithColon("hello", "this", "is", "test")

	if got != expected {
		t.Error("expected", expected, "got", got)
	}
}

func TestDigestIsComplete(t *testing.T) {
	d := digest{}
	if d.isComplete() {
		t.Error("Expected h.isComplete() to be false")
	}

	d.realm = "somestring"
	d.algorithm = "somestring"
	d.aOne = "somestring"
	d.aTwo = "somestring"
	d.cNonce = "somestring"
	d.key = "somestring"
	d.method = "somestring"
	d.name = "somestring"
	d.nC = 0x1
	d.nOnce = "somestring"
	d.opaque = "somestring"
	d.path = "somestring"
	d.qop = "somestring"

	if !d.isComplete() {
		t.Error("Expected h.isComplete() to be true")
	}

	d.key = ""

	if d.isComplete() {
		t.Error("Expected h.isComplete() to be false")
	}
}

func TestDigestHashWithColon(t *testing.T) {
	expected := "c1832f88c38ea538adc290536b3f7fcf"

	got, err := hashWithColon(md5.New(), "hello", "this", "is", "test")
	if err != nil {
		t.Error(err)
	}

	if got != expected {
		t.Error("got", got, "expected", expected)
	}
}

func TestDigestCalculateResponse(t *testing.T) {
	// TODO: finish
}
