package swclient

import (
	"net/http"
	"strings"
)

type header struct {
	realm     string
	qop       string
	method    string
	nOnce     string
	opaque    string
	algorithm string
	aOne      string
	aTwo      string
	cNonce    string
	path      string
	nC        int16
	name      string
	key       string
}

// parseParameters saves the values for realm, nOnce, opaque, algorithm and qop
// from a (digest) response header into the local header struct
func (h *header) parseParameters(response http.Response) *header {

	// get the protocol info from the responses auth header
	responseAuthHeader := response.Header.Get("Www-Authenticate")
	// delete the "Digest" keyword from the beginning of the response string
	cleanAuthHeader := strings.Trim(responseAuthHeader, "Digest ")
	// split the response string into a slice
	keyValuePairs := strings.Split(cleanAuthHeader, ", ")

	// auth will hold the all data that was supplied by the response string
	auth := map[string]string{}

	// split pair strings into keys and values and save them in auth[]
	for _, pair := range keyValuePairs {
		tuple := strings.Split(pair, "=")
		key := tuple[0]
		value := strings.Replace(tuple[1], "\"", "", -1) // this just strips tuple[1] from quotation marks
		auth[key] = value
	}

	// assign all values
	h.realm = auth["realm"]
	h.nOnce = auth["nonce"]
	h.opaque = auth["opaque"]
	h.algorithm = auth["algorithm"]
	h.qop = auth["qop"]

	return h
}

func (h *header) isComplete() bool {
	if h.realm == "" {
		return false
	}
	if h.qop == "" {
		return false
	}
	if h.method == "" {
		return false
	}
	if h.nOnce == "" {
		return false
	}
	if h.opaque == "" {
		return false
	}
	if h.algorithm == "" {
		return false
	}
	if h.aOne == "" {
		return false
	}
	if h.aTwo == "" {
		return false
	}
	if h.cNonce == "" {
		return false
	}
	if h.path == "" {
		return false
	}
	if h.nC == 0x0 {
		return false
	}
	if h.name == "" {
		return false
	}
	if h.key == "" {
		return false
	}
	return true
}
