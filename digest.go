package swclient

import (
	"crypto/md5"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
)

// digest holds all information required for digest-authentication.
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

// generateRequest generates a new http.Request which carries all necessary authentication information.
func (d *digest) generateRequest(method string, uri string, body io.Reader, username string, key string, serverinfo *http.Response) (*http.Request, error) {
	// if serverinfo is given
	if serverinfo != nil {
		// parse server info
		auth, err := parseParameters(serverinfo)
		if err != nil {
			return nil, cerror{"swclient/digest.go", "generateRequest()", err.Error()}
		}
		d.realm = auth["realm"]
		d.nOnce = auth["nonce"]
		d.opaque = auth["opaque"]
		d.algorithm = auth["algorithm"]
		d.qop = auth["qop"]
	}
	// calculate response to server challenge
	response, err := d.calculateResponse(method, uri, username, key)
	if err != nil {
		return nil, cerror{"swclient/digest.go", "generateRequest()", err.Error()}
	}
	// construct the digest header string
	authHeader := fmt.Sprintf(
		`Digest username="%s", realm="%s", nonce="%s", uri="%s", cnonce="%s", nc=%08x, qop=%s, response="%s"`,
		d.name, d.realm, d.nOnce, d.path, d.cNonce, d.nC, d.qop, response)
	// if an opaque was provided, add it
	if len(d.opaque) > 0 {
		authHeader = fmt.Sprintf(`%s, opaque="%s"`, authHeader, d.opaque)
	}
	// generate standard request
	request, err := http.NewRequest(method, uri, body)
	if err != nil {
		return nil, cerror{"swclient/digest.go", "generateRequest()", err.Error()}
	}
	// set the authorization, host and content-type headers
	request.Header.Set("Authorization", authHeader)
	request.Header.Set("Host", request.Host)
	request.Header.Set("Content-Type", "application/json; charset=utf-8")
	// return authenticated request
	return request, nil
}

// calculateResponse calculates the response string the server expects.
func (d *digest) calculateResponse(method string, uri string, username string, key string) (string, error) {
	// increment request count
	d.nC += 0x1
	// calculate new cNonce
	cNonce, err := hashRand()
	if err != nil {
		return "", cerror{"swclient/digest.go", "calculateResponse()", err.Error()}
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
	aOne, err := hashWithColon(d.name, d.realm, d.key)
	if err != nil {
		return "", cerror{"swclient/digest.go", "calculateResponse()", err.Error()}
	}
	// set aOne
	d.aOne = aOne
	// calculate aOne
	aTwo, err := hashWithColon(d.method, d.path)
	if err != nil {
		return "", cerror{"swclient/digest.go", "calculateResponse()", err.Error()}
	}
	// set aTwo
	d.aTwo = aTwo
	// calculate response
	response, err := hashWithColon(d.aOne, d.nOnce, fmt.Sprintf("%08x", d.nC), d.cNonce, d.qop, d.aTwo)
	if err != nil {
		return "", cerror{"swclient/digest.go", "calculateResponse()", err.Error()}
	}
	return response, nil
}

// parseParameters gets the values for realm, nOnce, opaque, algorithm and qop from a HTTP response header.
func parseParameters(response *http.Response) (map[string]string, error) {
	// auth will hold the all data that was supplied by the response string
	auth := map[string]string{}

	if len(response.Header.Get("Www-Authenticate")) <= 0 {
		return auth, fmt.Errorf("\n%s, %s: No \"Www-Authenticate\"-field in response header. %s: \"%s\"", "swclient/digest.go", "parseParameters()", response.Status, response.Request.RequestURI)
	}
	// get the protocol info from the responses auth header
	responseAuthHeader := response.Header.Get("Www-Authenticate")

	if !strings.Contains(responseAuthHeader, "Digest") {
		return auth, fmt.Errorf("\n%s, %s: No digest info in response header. %s: \"%s\"", "swclient/digest.go", "parseParameters()", response.Status, response.Request.RequestURI)
	}
	// trim "Digest " from the beginning of the response string
	cleanAuthHeader := strings.Trim(responseAuthHeader, "Digest ")

	if !strings.Contains(cleanAuthHeader, ", ") {
		return auth, fmt.Errorf("\n%s, %s: Response header doesn't contain enough info. %s: \"%s\"", "swclient/digest.go", "parseParameters()", response.Status, response.Request.RequestURI)
	}
	// split the response string into a slice
	keyValuePairs := strings.Split(cleanAuthHeader, ", ")

	if !strings.Contains(keyValuePairs[0], "=") {
		return auth, fmt.Errorf("\n%s, %s: Response header doesn't contain key=value pairs. %s: \"%s\"", "swclient/digest.go", "parseParameters()", response.Status, response.Request.RequestURI)
	}
	// split pair strings into keys and values and save them in auth[]
	for _, pair := range keyValuePairs {
		tuple := strings.Split(pair, "=")
		key := tuple[0]
		value := strings.Replace(tuple[1], "\"", "", -1) // this just strips tuple[1] from quotation marks
		auth[key] = value
	}

	return auth, nil
}

// hashWithColon takes a slice of strings, joins them into a single string separated with colons and hashes that.
func hashWithColon(parts ...string) (string, error) {
	hashed, err := hashString(joinWithColon(parts...))
	if err != nil {
		return "", cerror{"swclient/digest.go", "hashWithColon()", err.Error()}
	}
	return hashed, nil
}

// hash returns the hash of the string passed to it.
func hashString(str string) (string, error) {
	return fmt.Sprintf("%x", md5.Sum([]byte(str))), nil // %x -> hexadecimal
}

// hashRand returns a hashed pseudo-random int.
func hashRand() (string, error) {
	return hashString(strconv.Itoa(rand.Int()))
}

// joinWithColon joins a slice of strings into one string separated with colons.
func joinWithColon(str ...string) string {
	return strings.Join(str, ":")
}
