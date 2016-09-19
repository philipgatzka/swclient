package swclient

import (
	"bytes"
	"crypto/md5"
	"net/http"
	"testing"
)

// mockHasher is a struct for testing the hashing functions
type mockHasher struct{}

// Reset is required by hasher interface
func (mockHasher) Reset() {}

// Write is reduired by hasher interface
func (mockHasher) Write(b []byte) (int, error) {
	return 0, nil
}

// Sum is required by hasher interface
func (mockHasher) Sum(b []byte) []byte {
	return []byte("hello, this is mockHasher!")
}

var mockDigest = digest{
	realm:     "test",
	qop:       "test",
	method:    "test",
	nOnce:     "test",
	opaque:    "test",
	algorithm: "test",
	aOne:      "test",
	aTwo:      "test",
	cNonce:    "test",
	path:      "test",
	nC:        0x1,
	name:      "test",
	key:       "test",
}

func TestDigestGenerateRequest(t *testing.T) {
	d := mockDigest
	serverinfo := &http.Response{}

	body := bytes.NewBufferString("test")

	expected, err := http.NewRequest(d.method, d.path, body)
	if err != nil {
		t.Error(err)
	}

	authHeader := `Digest username="test", realm="test", nonce="test", uri="test", cnonce="68656c6c6f2c2074686973206973206d6f636b48617368657221", nc=00000002, qop=test, response="68656c6c6f2c2074686973206973206d6f636b48617368657221", opaque="test"`

	expected.Header.Set("Authorization", authHeader)
	expected.Header.Set("Host", expected.Host)
	expected.Header.Set("Content-Type", "application/json; charset=utf-8")

	got, err := d.generateRequest(d.method, d.path, body, d.name, d.key, serverinfo, mockHasher{})
	if err != nil {
		t.Error(err)
	}

	if expected.Header.Get("Authorization") != got.Header.Get("Authorization") {
		t.Error("got", got.Header.Get("Authorization"), "expected", expected.Header.Get("Authorization"))
	}

	if expected.Header.Get("Host") != got.Header.Get("Host") {
		t.Error("got", got.Header.Get("Host"), "expected", expected.Header.Get("Host"))
	}

	if expected.Header.Get("Content-Type") != got.Header.Get("Content-Type") {
		t.Error("got", got.Header.Get("Content-Type"), "expected", expected.Header.Get("Content-Type"))
	}
}

func TestDigestCalculateResponse(t *testing.T) {
	d := mockDigest
	expected := "68656c6c6f2c2074686973206973206d6f636b48617368657221"

	got, err := d.calculateResponse("GET", "http://hello.this/is/test", "user", "key", mockHasher{})
	if err != nil {
		t.Error(err)
	}

	if got != expected {
		t.Error("got", got, "expected", expected)
	}
}

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

func TestDigestJoinWithColon(t *testing.T) {
	expected := "hello:this:is:test"
	got := joinWithColon("hello", "this", "is", "test")

	if got != expected {
		t.Error("expected", expected, "got", got)
	}
}
