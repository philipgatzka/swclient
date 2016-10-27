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
	"reflect"
)

// Swclient defines the interface this library exposes
type Swclient interface {
	GetSingle(id string, o interface{}) error
	GetSingleRaw(resource string, id string) (*Response, error)
	PutSingle(id string, o interface{}) (*Response, error)
	PutSingleRaw(resource string, id string, body io.Reader) (*Response, error)
}

// swclient holds client and server information
type swclient struct {
	user      string
	key       string
	baseurl   *url.URL
	apiurl    string
	resources map[string]string
	dgc       *digestclient
	hshr      hash.Hash
}

// New returns an initialised swclient
func New(user string, key string, apiurl string) (Swclient, error) {
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

	resources := map[string]string{
		"*address.Address":           "addresses",
		"*article.Article":           "articles",
		"*cache.Cache":               "caches",
		"*category.Category":         "categories",
		"*country.Country":           "countries",
		"*customer.Customer":         "customers",
		"*manufacturer.Manufacturer": "manufacturers",
		"*media.Media":               "media",
		"*order.Order":               "orders",
		"*shop.Shop":                 "shops",
		"*translation.Translation":   "translations",
		"*variant.Variant":           "variants",
	}

	// initialise and return
	return &swclient{
		user:    user,
		key:     key,
		baseurl: u,
		apiurl:  u.Path,
		dgc: &digestclient{
			dgst:  &digest{},
			httpc: &http.Client{},
		},
		hshr:      md5.New(),
		resources: resources,
	}, nil
}

// GetSingle
func (s swclient) GetSingle(id string, o interface{}) error {
	if res, ok := s.typeIsResource(o); ok {
		resp, err := s.GetSingleRaw(res, id)
		if err != nil {
			return err
		}

		err = json.Unmarshal(resp.Data, o)
		if err != nil {
			return err
		}
	} else {
		return errors.New("Passed type is not a resource of the shopware api!")
	}
	return nil
}

// GetSinglRaw
func (s swclient) GetSingleRaw(resource string, id string) (*Response, error) {
	return s.request("GET", resource, id, bytes.NewBufferString(""))
}

// PutSingle
func (s swclient) PutSingle(id string, o interface{}) (*Response, error) {
	if res, ok := s.typeIsResource(o); ok {
		bts, err := json.Marshal(o)
		if err != nil {
			return nil, err
		}
		return s.request("PUT", res, id, bytes.NewReader(bts))
	} else {
		return nil, errors.New("Passed type is not a resource of the shopware api!")
	}
}

// PutSingleRaw
func (s swclient) PutSingleRaw(resource string, id string, body io.Reader) (*Response, error) {
	return s.request("PUT", resource, id, body)
}

// typeIsResource checks if the given type has a type->resource mapping (and therefore is a resource of the shopware api)
func (s swclient) typeIsResource(o interface{}) (string, bool) {
	res, isResource := s.resources[reflect.TypeOf(o).String()]
	return res, isResource
}

// request executes an http-request of the given method
func (s *swclient) request(method string, resource string, id string, body io.Reader) (*Response, error) {
	// join shopware base-url, api-endpoint, resource and id
	s.baseurl.Path = path.Join(s.apiurl, resource, id)

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
