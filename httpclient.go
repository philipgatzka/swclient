package swclient

import (
	"bytes"
	"io"
	"net/http"
)

type httpclient struct {
	d            *digest
	c            *http.Client
	lastResponse *http.Response
}

// request executes an http-request of the given method
func (h *httpclient) request(method string, uri string, body io.Reader, username string, key string, hshr hasher) (*http.Response, error) {
	// if we need to (re-)authenticate
	if h.lastResponse == nil || h.lastResponse.StatusCode != 200 {
		// probe server
		resp, err := h.c.Get(uri)
		if err != nil {
			return nil, err
		}
		// save info
		h.lastResponse = resp
	}
	// generate new request
	req, err := h.d.generateRequest(method, uri, body, username, key, h.lastResponse, hshr)
	if err != nil {
		return nil, err
	}
	// execute, return
	resp, err := h.c.Do(req)
	if err != nil {
		return nil, err
	}
	h.lastResponse = resp
	return resp, nil
}

// get executes a get request
func (h httpclient) get(uri string, username string, key string, hshr hasher) (*http.Response, error) {
	return h.request("GET", uri, bytes.NewBufferString(""), username, key, hshr)
}

// put executes a put request
func (h httpclient) put(uri string, body io.Reader, username string, key string, hshr hasher) (*http.Response, error) {
	return h.request("PUT", uri, body, username, key, hshr)
}

// post executes a post request
func (h httpclient) post(uri string, body io.Reader, username string, key string, hshr hasher) (*http.Response, error) {
	return h.request("POST", uri, body, username, key, hshr)
}

// del executes a delete request
func (h httpclient) del(uri string, username string, key string, hshr hasher) (*http.Response, error) {
	return h.request("DELETE", uri, bytes.NewBufferString(""), username, key, hshr)
}
