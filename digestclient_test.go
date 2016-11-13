package swclient

import (
	"bytes"
	"crypto/md5"
	"errors"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

var mockResponseBody = `{"data":[{"id":1, "name": "Test"}], "success": true, "total": 1}`

func TestRequest(t *testing.T) {
	ts, err := testDigestServer("401")
	if err != nil {
		t.Error(err)
	}
	defer ts.Close()

	dgst := digestclient{
		dgst:  &digest{},
		httpc: &http.Client{},
	}

	_, err = dgst.request("GET", ts.URL, bytes.NewBufferString(""), "user", "key", md5.New())
	if err != nil {
		t.Error(err)
	}

	resp, err := dgst.request("GET", ts.URL, bytes.NewBufferString(""), "user", "key", md5.New())
	if err != nil {
		t.Error(err)
	}

	h := resp.Request.Header

	auth := h.Get("Authorization")

	if auth == "" {
		t.Error("digestclient didn't generate an authorised request!")
	} else if !strings.Contains(auth, `opaque="thisisatest"`) {
		t.Error("digestclient didn't parse server info on first try!", h)
	}
}

func TestExec(t *testing.T) {

	ts, err := testExecServer()
	if err != nil {
		t.Error(err)
	}
	defer ts.Close()

	err = request("GET", ts.URL, bytes.NewBufferString(mockResponseBody+"get"), mockResponseBody+"get")
	if err != nil {
		t.Error(err)
	}

	err = request("POST", ts.URL, bytes.NewBufferString(mockResponseBody+"post"), mockResponseBody+"post")
	if err != nil {
		t.Error(err)
	}

	err = request("PUT", ts.URL, bytes.NewBufferString(mockResponseBody+"put"), mockResponseBody+"put")
	if err != nil {
		t.Error(err)
	}

	err = request("DELETE", ts.URL, bytes.NewBufferString(mockResponseBody+"delete"), mockResponseBody+"delete")
	if err != nil {
		t.Error(err)
	}
}

func testDigestServer(status string) (*httptest.Server, error) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		bts, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Fatal(err)
		}
		r.Body.Close()
		w.Header().Set("Status", status)
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Www-Authenticate", `Digest realm="Shopware REST-API", domain="/", nonce="nonce", opaque="thisisatest", algorithm="MD5", qop="auth"`)
		w.Write(bts)
	}))
	return ts, nil
}

func testExecServer() (*httptest.Server, error) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		bts, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Fatal(err)
		}
		r.Body.Close()
		w.Write(bts)
	}))
	return ts, nil
}

func request(method string, uri string, body io.Reader, expected string) error {
	req, err := http.NewRequest(method, uri, body)
	if err != nil {
		return err
	}

	dgst := digestclient{
		dgst:  &digest{},
		httpc: &http.Client{},
	}

	resp, err := dgst.exec(req)
	defer resp.Body.Close()
	if err != nil {
		return err
	}

	bts, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if string(bts) != expected {
		return errors.New("expected " + expected + " got " + string(bts))
	}
	return nil
}
