package swclient

import (
	"bytes"
	"crypto/md5"
	"errors"
	"hash"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
	"strconv"
)

// swclient wraps the package and holds client and server information
type swclient struct {
	user     string
	key      string
	apiurl   string
	resource string
	dgc      *digestclient
	hshr     hash.Hash
}

// New returns an initialised swclient
func New(user string, key string, apiurl string) *swclient {
	return &swclient{
		user:   user,
		key:    key,
		apiurl: apiurl,
		dgc: &digestclient{
			dgst:  &digest{},
			httpc: &http.Client{},
		},
		hshr: md5.New(),
	}
}

// Resource sets the resource attribute of swclient to res
// res will be appended to the apiurl before the next request
func (s *swclient) Resource(res string) *swclient {
	s.resource = res
	return s
}

// TODO: make all request methods members of a resource struct or find another way of ensuring a resource is given before request functions can be called
func (s swclient) GetById(id int) (string, error) {
	return s.request("GET", strconv.Itoa(id), bytes.NewBufferString(""))
}

func (s swclient) Get() (string, error) {
	return s.request("GET", "", bytes.NewBufferString(""))
}

func (s swclient) PutById(id int, body io.Reader) (string, error) {
	return s.request("PUT", strconv.Itoa(id), body)
}

func (s swclient) PostById(id int, body io.Reader) (string, error) {
	return s.request("POST", strconv.Itoa(id), body)
}

func (s swclient) DelById(id int) (string, error) {
	return s.request("DELETE", strconv.Itoa(id), bytes.NewBufferString(""))
}

func (s swclient) request(method string, uri string, body io.Reader) (string, error) {
	fullUri, err := s.constructUri(uri)
	if err != nil {
		return "", err
	}
	resp, err := s.dgc.request(method, fullUri, body, s.user, s.key, s.hshr)
	if err != nil {
		return "", err
	}
	resstr, err := responseString(resp)
	if err != nil {
		return "", err
	}
	return resstr, nil
}

func (s *swclient) constructUri(uri string) (string, error) {
	u, err := url.Parse(s.apiurl)
	if err != nil {
		return "", err
	}

	if len(s.resource) > 0 {
		u.Path = path.Join(u.Path, s.resource, uri)
		s.resource = ""
	} else {
		u.Path = path.Join(u.Path, uri)
	}

	return u.String(), nil
}

func responseString(resp *http.Response) (string, error) {
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	if resp.StatusCode != 200 {
		return "", errors.New(resp.Status)
	}

	return string(b), nil
}
