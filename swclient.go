package swclient

import (
	"bytes"
	"crypto/md5"
	"encoding/json"
	"errors"
	"hash"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
	"strconv"
)

// TODO: higher-level functions like Put(id int, a article.Article)

// Swclient defines the interface this library exposes
type Swclient interface {
	Resource(string) (*swclient, error)
	Get() (*Response, error)
	GetById(id int) (*Response, error)
	PutById(id int, body io.Reader) (*Response, error)
	PostById(id int, body io.Reader) (*Response, error)
	DelById(id int) (*Response, error)
}

// swclient holds client and server information
type swclient struct {
	user     string
	key      string
	apiurl   string
	resource string
	dgc      *digestclient
	hshr     hash.Hash
}

// Response defines a response from shopware
type Response struct {
	Data    json.RawMessage `json:",omitempty"`
	Success bool            `json:",omitempty"`
}

// New returns an initialised swclient
func New(user string, key string, apiurl string, resource string) (*swclient, error) {
	// check input
	if len(user) <= 0 {
		return nil, errors.New("Can't create swclient: user not specified")
	}

	if len(key) <= 0 {
		return nil, errors.New("Can't create swclient: api-key not specified")
	}

	if len(apiurl) <= 0 {
		return nil, errors.New("Can't create swclient: url to api not specified")
	}

	if len(resource) <= 0 {
		return nil, errors.New("Can't create swclient: no resource specified")
	}
	// initialise and return
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
	}, nil
}

// Resource changes the api-resource swclient points to
func (s *swclient) Resource(resource string) (*swclient, error) {
	if len(resource) <= 0 {
		return nil, errors.New("No resource specified")
	}
	s.resource = resource
	return s, nil
}

// GetById gets an object of the type specified in swclient.New() or swclient.Resource() from the shopware-api
func (s swclient) GetById(id int) (*Response, error) {
	return s.request("GET", strconv.Itoa(id), bytes.NewBufferString(""))
}

// Get gets a list of objects of the type specified in swclient.New() or swclient.Resource() from the shopware-api
func (s swclient) Get() (*Response, error) {
	return s.request("GET", "", bytes.NewBufferString(""))
}

// PutById modifies an object of the type specified in swclient.New() or swclient.Resource() via the shopware-api
func (s swclient) PutById(id int, body io.Reader) (*Response, error) {
	return s.request("PUT", strconv.Itoa(id), body)
}

// PostById creates an object of the type specified in swclient.New() or swclient.Resource() via the shopware-api
func (s swclient) PostById(id int, body io.Reader) (*Response, error) {
	return s.request("POST", strconv.Itoa(id), body)
}

// DelById deletes an object of the type specified in swclient.New() or swclient.Resource() via the shopware-api
func (s swclient) DelById(id int) (*Response, error) {
	return s.request("DELETE", strconv.Itoa(id), bytes.NewBufferString(""))
}

// request executes an http-request of the given method
func (s swclient) request(method string, uri string, body io.Reader) (*Response, error) {
	fullUri, err := s.constructUri(uri)
	if err != nil {
		return nil, err
	}

	resp, err := s.dgc.request(method, fullUri, body, s.user, s.key, s.hshr)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, errors.New(resp.Status)
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()

	// unmarshal received data into swclient.Response
	data := Response{}
	err = json.Unmarshal(b, &data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (s *swclient) constructUri(uri string) (string, error) {
	u, err := url.Parse(s.apiurl)
	if err != nil {
		return "", err
	}
	// join elements
	u.Path = path.Join(u.Path, s.resource, uri)
	// return url as string
	return u.String(), nil
}
