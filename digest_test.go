package swclient

import (
	"bytes"
	"crypto/md5"
	"fmt"
	"net/http"
	"testing"
)

// mockHasher is a struct for testing the hashing functions.
type mockHasher struct{}

// Reset is required by hasher interface.
func (mockHasher) Reset() {}

// Write is required by hasher interface.
func (mockHasher) Write(b []byte) (int, error) {
	return 0, nil
}

// Sum is required by hasher interface.
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

var mockUsername = "test"
var mockRealm = "test"
var mockNonce = "test"
var mockUri = "test"
var mockCnonce = "68656c6c6f2c2074686973206973206d6f636b48617368657221"
var mockNc = "00000002"
var mockResponse = "68656c6c6f2c2074686973206973206d6f636b48617368657221"

var mockData = map[string]string{
	"username": mockUsername,
	"realm":    mockRealm,
	"nonce":    mockNonce,
	"uri":      mockUri,
	"cnonce":   mockCnonce,
	"nc":       mockNc,
	"response": mockResponse,
}

var mockResponseHeader = fmt.Sprintf(`Digest username="%s", realm="%s", nonce="%s", uri="%s", cnonce="%s", nc=%s, qop=, response="%s"`, mockUsername, mockRealm, mockNonce, mockUri, mockCnonce, mockNc, mockResponse)

var mockServerResponse = &http.Response{
	Header: map[string][]string{
		"Www-Authenticate": []string{
			mockResponseHeader,
		},
	},
}

func TestDigestGenerateRequest(t *testing.T) {
	d := mockDigest

	// define what is expected
	expected, err := http.NewRequest(d.method, d.path, bytes.NewBufferString("test"))
	if err != nil {
		t.Error(err)
	}
	expected.Header.Set("Authorization", mockResponseHeader)
	expected.Header.Set("Host", expected.Host)
	expected.Header.Set("Content-Type", "application/json; charset=utf-8")

	// see what is returned
	got, err := d.generateRequest(d.method, d.path, bytes.NewBufferString("test"), d.name, d.key, mockServerResponse, mockHasher{})
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
	got, err := mockDigest.calculateResponse("GET", "http://hello.this/is/irrelevant", "someUser", "someKey", mockHasher{})
	if err != nil {
		t.Error(err)
	}

	if got != mockResponse {
		t.Error("got", got, "expected", mockResponse)
	}
}

func TestDigestParseParameters(t *testing.T) {
	tuples, err := parseParameters(mockServerResponse)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	for key, val := range tuples {
		if val != mockData[key] {
			t.Error("expected", mockData[key], "got", val)
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
