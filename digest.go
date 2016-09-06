package swclient

import (
	"errors"
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
func (h *header) parseParameters(response http.Response) map[string]string {

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

	// populate the header
	h.realm = auth["realm"]
	h.nOnce = auth["nonce"]
	h.opaque = auth["opaque"]
	h.algorithm = auth["algorithm"]
	h.qop = auth["qop"]

	return auth
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

// checksums calculates the hashes of aOne and aTwo
func (h *header) checksums(hasher hash.Hash) error {
	// check if name, realm and key have a value
	if h.name == "" || h.realm == "" || h.key == "" {
		return errors.New("name, realm or key missing from header!")
	}
	// join
	aOne := []string{h.name, h.realm, h.key}
	// cheack if method and path have a value
	if h.method == "" || h.path == "" {
		return errors.New("method or path missing from header!")
	}
	// join
	aTwo := []string{h.method, h.path}
	// calculate hash
	hasher.Reset()
	aOneHash, err := hashString(joinWithColon(aOne), hasher)
	if err != nil {
		return err
	}
	// calculate hash
	hasher.Reset()
	aTwoHash, err := hashString(joinWithColon(aTwo), hasher)
	if err != nil {
		return err
	}
	// assign, return
	h.aOne = aOneHash
	h.aTwo = aTwoHash
	return nil
}

// joinWithColon joins a slice of strings into one string separated with colons
func joinWithColon(str []string) string {
	return strings.Join(str, ":")
}
