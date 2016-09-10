package swclient

import (
	"fmt"
	"hash"
	"io"
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

// generateRequest uses the provided information to generate a new http.Request which has all the necessary information
// for digest-authentication
func (d *digest) generateRequest(method string, uri string, body io.Reader, username string, key string, hasher hash.Hash) (*http.Request, error) {
	// generate standard request
	request, err := http.NewRequest(method, uri, body)
	if err != nil {
		return nil, err
	}
	// calculate response to server challenge
	response, err := d.calculateResponse(method, uri, username, key, hasher)
	if err != nil {
		return nil, err
	}
	// construct the digest header string
	authHeader := fmt.Sprintf(
		`Digest username="%s", realm="%s", nonce="%s", uri="%s", cnonce="%s", nc=%08x, qop=%s, response="%s"`,
		d.name, d.realm, d.nOnce, d.path, d.cNonce, d.nC, d.qop, response)
	// if an opaque was provided, add it
	if d.opaque != "" {
		authHeader = fmt.Sprintf(`%s, opaque="%s"`, authHeader, d.opaque)
	}
	// set the authorization, host and content-type headers
	request.Header.Set("Authorization", authHeader)
	request.Header.Set("Host", request.Host)
	request.Header.Set("Content-Type", "application/json; charset=utf-8")
	// return authenticated request
	return request, nil
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

// parsedParameters checks if all fields required have been provided by the server; realm, nOnce, opaque and qop have to be set
func (d digest) parsedParameters() bool {
	return !equalsEmptyString(d.realm, d.nOnce, d.opaque, d.qop)
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
