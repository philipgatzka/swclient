package swclient

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

// digest holds all information required for digest-authentication
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

// generateRequest uses the provided information to generate a new http.Request which has all the necessary information for digest-authentication
func (d *digest) generateRequest(method string, uri string, body io.Reader, username string, key string, serverinfo *http.Response, hshr hasher) (*http.Request, error) {
	// if serverinfo is given
	if serverinfo != nil {
		// parse server info
		auth, err := parseParameters(serverinfo)
		if err != nil {
			return nil, swerror{"swclient/digest.go", "generateRequest()", err.Error()}
		}
		d.realm = auth["realm"]
		d.nOnce = auth["nonce"]
		d.opaque = auth["opaque"]
		d.algorithm = auth["algorithm"]
		d.qop = auth["qop"]
	}
	// calculate response to server challenge
	response, err := d.calculateResponse(method, uri, username, key, hshr)
	if err != nil {
		return nil, swerror{"swclient/digest.go", "generateRequest()", err.Error()}
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
		return nil, swerror{"swclient/digest.go", "generateRequest()", err.Error()}
	}
	// set the authorization, host and content-type headers
	request.Header.Set("Authorization", authHeader)
	request.Header.Set("Host", request.Host)
	request.Header.Set("Content-Type", "application/json; charset=utf-8")
	// return authenticated request
	return request, nil
}

// calculateResponse calculates the response string the server requires
func (d *digest) calculateResponse(method string, uri string, username string, key string, hshr hasher) (string, error) {
	// increment request count
	d.nC += 0x1
	// calculate new cNonce
	cNonce, err := hashNow(hshr)
	if err != nil {
		return "", swerror{"swclient/digest.go", "calculateResponse()", err.Error()}
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
	aOne, err := hashWithColon(hshr, d.name, d.realm, d.key)
	if err != nil {
		return "", swerror{"swclient/digest.go", "calculateResponse()", err.Error()}
	}
	// set aOne
	d.aOne = aOne
	// calculate aOne
	aTwo, err := hashWithColon(hshr, d.method, d.path)
	if err != nil {
		return "", swerror{"swclient/digest.go", "calculateResponse()", err.Error()}
	}
	// set aTwo
	d.aTwo = aTwo
	// calculate response
	response, err := hashWithColon(hshr, d.aOne, d.nOnce, fmt.Sprintf("%08x", d.nC), d.cNonce, d.qop, d.aTwo)
	if err != nil {
		return "", swerror{"swclient/digest.go", "calculateResponse()", err.Error()}
	}
	return response, nil
}

// parseParameters gets the values for realm, nOnce, opaque, algorithm and qop from a response header
func parseParameters(response *http.Response) (map[string]string, error) {
	// auth will hold the all data that was supplied by the response string
	auth := map[string]string{}

	if response.Header.Get("Www-Authenticate") != "" {
		// get the protocol info from the responses auth header
		responseAuthHeader := response.Header.Get("Www-Authenticate")

		if strings.Contains(responseAuthHeader, "Digest") {
			// trim "Digest " from the beginning of the response string
			cleanAuthHeader := strings.Trim(responseAuthHeader, "Digest ")

			if strings.Contains(cleanAuthHeader, ", ") {
				// split the response string into a slice
				keyValuePairs := strings.Split(cleanAuthHeader, ", ")

				if strings.Contains(keyValuePairs[0], "=") {

					// split pair strings into keys and values and save them in auth[]
					for _, pair := range keyValuePairs {
						tuple := strings.Split(pair, "=")
						key := tuple[0]
						value := strings.Replace(tuple[1], "\"", "", -1) // this just strips tuple[1] from quotation marks
						auth[key] = value
					}
				} else {
					return auth, fmt.Errorf("\n%s, %s: Response header doesn't contain key=value pairs. %s: \"%s\"", "swclient/digest.go", "parseParameters()", response.Status, response.Request.RequestURI)
				}
			} else {
				return auth, fmt.Errorf("\n%s, %s: Response header doesn't contain enough info. %s: \"%s\"", "swclient/digest.go", "parseParameters()", response.Status, response.Request.RequestURI)
			}
		} else {
			return auth, fmt.Errorf("\n%s, %s: No digest info in response header. %s: \"%s\"", "swclient/digest.go", "parseParameters()", response.Status, response.Request.RequestURI)
		}
	} else {
		return auth, fmt.Errorf("\n%s, %s: No \"Www-Authenticate\"-field in response header. %s: \"%s\"", "swclient/digest.go", "parseParameters()", response.Status, response.Request.RequestURI)
	}

	return auth, nil
}

// hashWithColon takes a slice of string, joins its parts into a single string with colons and hashes that
func hashWithColon(hshr hasher, parts ...string) (string, error) {
	hashed, err := hashString(joinWithColon(parts...), hshr)
	if err != nil {
		return "", swerror{"swclient/digest.go", "hashWithColon()", err.Error()}
	}
	return hashed, nil
}

// hash returns the md5 hash of the supplied string
func hashString(str string, hshr hasher) (string, error) {
	// reset hasher because it could have been used before
	hshr.Reset()

	_, err := hshr.Write([]byte(str))
	if err != nil {
		return "", swerror{"swclient/digest.go", "hashString()", err.Error()}
	}
	return fmt.Sprintf("%x", hshr.Sum(nil)), nil // %x -> base 16
}

// hashNow returns the hashed system time at the time of execution
func hashNow(hshr hasher) (string, error) {
	return hashString(time.Now().String(), hshr)
}

// joinWithColon joins a slice of strings into one string separated with colons
func joinWithColon(str ...string) string {
	return strings.Join(str, ":")
}
