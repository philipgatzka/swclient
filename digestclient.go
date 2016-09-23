package swclient

import (
	"io"
	"net/http"
)

type digestclient struct {
	dgst       *digest
	httpc      *http.Client
	serverinfo *http.Response
}

// request executes an http-request of the given method
func (h *digestclient) request(method string, uri string, body io.Reader, username string, key string, hshr hasher) (*http.Response, error) {
	// if we need to (re-)authenticate
	if h.serverinfo == nil || h.serverinfo.StatusCode != 200 {
		// probe server
		serverinfo, err := h.httpc.Get(uri)
		if err != nil {
			return nil, err
		}
		// generate new request
		req, err := h.dgst.generateRequest(method, uri, body, username, key, serverinfo, hshr)
		if err != nil {
			return nil, err
		}
		return h.exec(req)
	} else {
		// generate new request
		req, err := h.dgst.generateRequest(method, uri, body, username, key, nil, hshr)
		if err != nil {
			return nil, err
		}
		return h.exec(req)
	}

}

// exec takes an *http.Request and executes it via the digestclients http.Client
func (h digestclient) exec(req *http.Request) (*http.Response, error) {
	// execute, return
	resp, err := h.httpc.Do(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
