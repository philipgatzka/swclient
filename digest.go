package swclient

import (
	"fmt"
	"hash"
	"net/http"
	"strings"
	"time"
)

// header holds all information required for a digest-request
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
	// trim "Digest " from the beginning of the response string
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
	h.qop = auth["qop"]
	h.algorithm = auth["algorithm"]

	return h
}

// hash returns the md5 hash of the supplied string
func hashString(str string, hasher hash.Hash) (string, error) {
	_, err := hasher.Write([]byte(str))
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", hasher.Sum(nil)), nil // %x renders the string in base 16
}

// hashNow returns the hashed system time at the time of execution
func hashNow(hasher hash.Hash) (string, error) {
	return hashString(time.Now().String(), hasher)
}

// isComplete checks if all fields in header are != an empty string
func (h header) isComplete() bool {
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
