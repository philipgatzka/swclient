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
)

// TODO: higher-level functions like Put(id int, a article.Article)

// Swclient defines the interface this library exposes
type Swclient interface {
	Resource(string) (*swclient, error)
	Get() (*Response, error)
	GetById(id string) (*Response, error)
	PutById(id string, body io.Reader) (*Response, error)
	PostById(id string, body io.Reader) (*Response, error)
	DelById(id string) (*Response, error)
}

// swclient holds client and server information
type swclient struct {
	user        string
	key         string
	baseurl     *url.URL
	apiEndpoint string
	resource    string
	dgc         *digestclient
	hshr        hash.Hash
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

	u, err := url.Parse(apiurl)
	if err != nil {
		return nil, err
	}

	if len(resource) <= 0 {
		return nil, errors.New("Can't create swclient: no resource specified")
	}
	// initialise and return
	return &swclient{
		user:        user,
		key:         key,
		baseurl:     u,
		apiEndpoint: u.Path,
		resource:    resource,
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
func (s swclient) GetById(id string) (*Response, error) {
	return s.request("GET", id, bytes.NewBufferString(""))
}

// Get gets a list of objects of the type specified in swclient.New() or swclient.Resource() from the shopware-api
func (s swclient) Get() (*Response, error) {
	return s.request("GET", "", bytes.NewBufferString(""))
}

// PutById modifies an object of the type specified in swclient.New() or swclient.Resource() via the shopware-api
func (s swclient) PutById(id string, body io.Reader) (*Response, error) {
	return s.request("PUT", id, body)
}

// PostById creates an object of the type specified in swclient.New() or swclient.Resource() via the shopware-api
func (s swclient) PostById(id string, body io.Reader) (*Response, error) {
	return s.request("POST", id, body)
}

// DelById deletes an object of the type specified in swclient.New() or swclient.Resource() via the shopware-api
func (s swclient) DelById(id string) (*Response, error) {
	return s.request("DELETE", id, bytes.NewBufferString(""))
}

// request executes an http-request of the given method
func (s *swclient) request(method string, id string, body io.Reader) (*Response, error) {
	// join shopware base-url, api-endpoint, resource and id
	s.baseurl.Path = path.Join(s.apiEndpoint, s.resource, id)

	// execute
	resp, err := s.dgc.request(method, s.baseurl.String(), body, s.user, s.key, s.hshr)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, errors.New(resp.Status)
	}
	// read response
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
