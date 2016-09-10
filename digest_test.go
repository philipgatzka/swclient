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

func TestDigestParsedParameters(t *testing.T) {
	d := digest{}
	if d.parsedParameters() {
		t.Error("Expected h.isComplete() to be false")
	}

	d.realm = "somestring"
	d.nOnce = "somestring"
	d.opaque = "somestring"
	d.qop = "somestring"

	if !d.parsedParameters() {
		t.Error("Expected h.isComplete() to be true")
	}

	d.realm = ""

	if d.parsedParameters() {
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
	d := digest{
		nC:    0x0,
		realm: "/is/test",
		qop:   "",
	}
	// expected := ""

	_, err := d.calculateResponse("GET", "http://hello.this/is/test", "user", "key", md5.New())
	if err != nil {
		t.Error(err)
	}

	// FIXME: need to test cNonce generation somehow
	// if got != expected {
	//	t.Error("got", got, "expected", expected)
	//}
}
