package swclient

import (
	"bytes"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

var dgst = digestclient{
	dgst:  &digest{},
	httpc: &http.Client{},
}

var mockResponseBody = `{"data":[{"id":1}]}`

func TestExec(t *testing.T) {

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if r.ContentLength == 0 {
			_, err := w.Write([]byte(mockResponseBody))
			if err != nil {
				t.Error(err)
			}
		} else if r.Method == "DELETE" {
			_, err := w.Write([]byte("deleted!"))
			if err != nil {
				t.Error(err)
			}
		} else {
			bts, err := ioutil.ReadAll(r.Body)
			if err != nil {
				t.Error(err)
			}
			_, err = w.Write(bts)
			if err != nil {
				t.Error(err)
			}
		}
	}))
	defer ts.Close()

	err := request("GET", ts.URL, bytes.NewBufferString(""), mockResponseBody)
	if err != nil {
		t.Error(err)
	}

	err = request("POST", ts.URL, bytes.NewBufferString(mockResponseBody+" post"), mockResponseBody+" post")
	if err != nil {
		t.Error(err)
	}

	err = request("PUT", ts.URL, bytes.NewBufferString(mockResponseBody+" put"), mockResponseBody+" put")
	if err != nil {
		t.Error(err)
	}

	err = request("DELETE", ts.URL, bytes.NewBufferString("deleted!"), "deleted!")
	if err != nil {
		t.Error(err)
	}
}

func request(method string, uri string, body io.Reader, expected string) error {
	req, err := http.NewRequest(method, uri, body)
	if err != nil {
		return err
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
