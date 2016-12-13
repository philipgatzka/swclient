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

// request executes an http-request.
func (d *digestclient) request(method string, uri string, body io.Reader, username string, key string) (*http.Response, error) {
	// if we need to (re-)authenticate
	if d.serverinfo == nil || d.serverinfo.StatusCode != 200 {
		// probe server
		serverinfo, err := d.httpc.Get(uri)
		if err != nil {
			return nil, cerror{"swclient/digestclient.go", "request()", err.Error()}
		}
		// generate new request
		req, err := d.dgst.generateRequest(method, uri, body, username, key, serverinfo)
		if err != nil {
			return nil, cerror{"swclient/digestclient.go", "request()", err.Error()}
		}
		return d.exec(req)
	}
	// generate new request
	req, err := d.dgst.generateRequest(method, uri, body, username, key, nil)
	if err != nil {
		return nil, cerror{"swclient/digestclient.go", "request()", err.Error()}
	}
	return d.exec(req)
}

// exec takes an *http.Request and executes it.
func (d digestclient) exec(req *http.Request) (*http.Response, error) {
	// execute, return
	resp, err := d.httpc.Do(req)
	if err != nil {
		return nil, cerror{"swclient/digestclient.go", "exec()", err.Error()}
	}
	return resp, nil
}
