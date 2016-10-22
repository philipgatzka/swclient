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
func New(user string, key string, apiurl string, resource string) *swclient {
	return &swclient{
		user:     user,
		key:      key,
		apiurl:   apiurl,
		resource: resource,
		dgc: &digestclient{
			dgst:  &digest{},
			httpc: &http.Client{},
		},
		hshr: md5.New(),
	}
}

func (s *swclient) Resource(res string) *swclient {
	s.resource = res
	return s
}

func (s swclient) GetById(id int) ([]byte, error) {
	return s.request("GET", strconv.Itoa(id), bytes.NewBufferString(""))
}

func (s swclient) Get() ([]byte, error) {
	return s.request("GET", "", bytes.NewBufferString(""))
}

func (s swclient) PutById(id int, body io.Reader) ([]byte, error) {
	return s.request("PUT", strconv.Itoa(id), body)
}

func (s swclient) PostById(id int, body io.Reader) ([]byte, error) {
	return s.request("POST", strconv.Itoa(id), body)
}

func (s swclient) DelById(id int) ([]byte, error) {
	return s.request("DELETE", strconv.Itoa(id), bytes.NewBufferString(""))
}

// request executes a request of the given method
func (s swclient) request(method string, uri string, body io.Reader) ([]byte, error) {
	fullUri, err := s.constructUri(uri)
	if err != nil {
		return nil, err
	}

	resp, err := s.dgc.request(method, fullUri, body, s.user, s.key, s.hshr)
	if err != nil {
		return nil, err
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, errors.New(resp.Status)
	}

	return b, nil
}

func (s *swclient) constructUri(uri string) (string, error) {
	u, err := url.Parse(s.apiurl)
	if err != nil {
		return "", err
	}

	if len(s.resource) > 0 {
		u.Path = path.Join(u.Path, s.resource, uri)
	} else {
		return "", errors.New("Empty resource attribute! Please specify the desired resource (e.g. 'articles') when calling swclient.New()")
	}

	return u.String(), nil
}
