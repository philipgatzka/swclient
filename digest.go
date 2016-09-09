package swclient

import (
	"fmt"
	"hash"
	"net/http"
	"strings"
	"time"
)

// digest holds all information required for a digest-request
type digest struct {
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

// calculateResponse calculates the response string the server requires
func (d *digest) calculateResponse(method string, uri string, username string, key string, hasher hash.Hash) (string, error) {
	// increment request count
	d.nC += 0x1

	// calculate new cNonce
	cNonce, err := hashNow(hasher)
	if err != nil {
		return "", err
	}
	d.cNonce = cNonce

	// set method
	d.method = method

	// set uri
	d.path = uri

	// set credentials
	d.name = username
	d.key = key

	// calculate aOne
	aOne, err := hashWithColon(hasher, d.name, d.realm, d.key)
	if err != nil {
		return "", err
	}
	// set aOne
	d.aOne = aOne

	// calculate aOne
	aTwo, err := hashWithColon(hasher, d.method, d.path)
	if err != nil {
		return "", err
	}
	// set aTwo
	d.aTwo = aTwo

	// calculate response
	response, err := hashWithColon(hasher, d.aOne, d.nOnce, fmt.Sprintf("%08x", d.nC), d.cNonce, d.qop, d.aTwo)
	if err != nil {
		return "", err
	}

	return response, nil

}

// hashWithColon takes a slice of string, joins its parts into a single string with colons and hashes that
func hashWithColon(hasher hash.Hash, parts ...string) (string, error) {
	hashed, err := hashString(joinWithColon(parts...), hasher)
	if err != nil {
		return "", err
	}
	return hashed, nil
}

// parseParameters gets the values for realm, nOnce, opaque, algorithm and qop from a response header
func parseParameters(response *http.Response) map[string]string {

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

	return auth
}

// hash returns the md5 hash of the supplied string
func hashString(str string, hasher hash.Hash) (string, error) {
	// reset hasher because it could have been used before
	hasher.Reset()

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

// joinWithColon joins a slice of strings into one string separated with colons
func joinWithColon(str ...string) string {
	return strings.Join(str, ":")
}

// isComplete checks if all fields in header are != an empty string
func (d digest) isComplete() bool {
	return !(equalsEmptyString(
		d.realm,
		d.qop,
		d.method,
		d.nOnce,
		d.opaque,
		d.algorithm,
		d.aOne,
		d.aTwo,
		d.cNonce,
		d.path,
		d.name,
		d.key) || d.nC == 0x0)
}

// equalsEmptyString returns true, if any of the provided strings is an empty string
func equalsEmptyString(strings ...string) bool {
	for _, s := range strings {
		if s == "" {
			return true
		}
	}
	return false
}
