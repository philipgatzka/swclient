package swclient

import (
	"io"
	"net/http"
	"fmt"
)

type digestclient struct {
	dgst       *digest
	httpc      *http.Client
	serverinfo *http.Response
}

// request executes an http-request of the given method
func (d *digestclient) request(method string, uri string, body io.Reader, username string, key string, hshr hasher) (*http.Response, error) {
	// if we need to (re-)authenticate
	if d.serverinfo == nil || d.serverinfo.StatusCode != 200 {
		// probe server
		serverinfo, err := d.httpc.Get(uri)
		if err != nil {
			return nil, fmt.Errorf("%s, %s: %s", "swclient/digestclient.go", "request()", err)
		}
		// generate new request
		req, err := d.dgst.generateRequest(method, uri, body, username, key, serverinfo, hshr)
		if err != nil {
			return nil, fmt.Errorf("%s, %s: %s", "swclient/digestclient.go", "request()", err)
		}
		return d.exec(req)
	} else {
		// generate new request
		req, err := d.dgst.generateRequest(method, uri, body, username, key, nil, hshr)
		if err != nil {
			return nil, fmt.Errorf("%s, %s: %s", "swclient/digestclient.go", "request()", err)
		}
		return d.exec(req)
	}

}

// exec takes an *http.Request and executes it via the digestclients http.Client
func (d digestclient) exec(req *http.Request) (*http.Response, error) {
	// execute, return
	resp, err := d.httpc.Do(req)
	if err != nil {
		return nil, fmt.Errorf("%s, %s: %s", "swclient/digestclient.go", "exec()", err)
	}
	return resp, nil
}
