package swclient

import (
	"net/http"
)

// get executes an http-get request using digest authentication
func get(uri string, user string, key string, h *header) (http.Response, error) {
	// get new auth params if the header is incomplete
	if !h.isComplete() {
		response, err := http.Get(uri)
		if err != nil {
			return nil, err
		}
		h.parseParameters(response)
	}

	// TODO: come up with a good structure for this

	return nil, nil
}
