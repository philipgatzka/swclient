package swclient

import (
	"bytes"
	"crypto/md5"
	"encoding/json"
	"hash"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
	"reflect"
	"fmt"
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

// swerror provides better error messages
type swerror struct {
	file string
	function string
	message string
}

func (s swerror) Error() string{
	return fmt.Sprintf("\n%s, %s: %s", s.file, s.function, s.message)
}

// New returns an initialised swclient
func New(user string, key string, apiurl string) (Swclient, error) {
	// check input
	if len(user) <= 0 {
		return nil, swerror{"swclient/swclient.go", "New()", "Can't create swclient: user not specified"}
	}

	if len(key) <= 0 {
		return nil, swerror{"swclient/swclient.go", "New()", "Can't create swclient: key not specified"}
	}

	if len(apiurl) <= 0 {
		return nil, swerror{"swclient/swclient.go", "New()", "Can't create swclient: api-url not specified"}
	}

	u, err := url.Parse(apiurl)
	if err != nil {
		return nil, swerror{"swclient/swclient.go", "New()", err.Error()}
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
	if res, ok := s.resources[reflect.TypeOf(o).String()]; ok {
		resp, err := s.GetSingleRaw(res, id)
		if err != nil {
			return swerror{"swclient/swclient.go", "GetSingle()", err.Error()}
		}

		err = json.Unmarshal(resp.Data, o)
		if err != nil {
			return swerror{"swclient/swclient.go", "GetSingle()", err.Error()}
		}
	} else {
		return swerror{"swclient/swclient.go", "GetSingle()", "Passed type is not a resource of the shopware api!"}
	}
	return nil
}

// GetSinglRaw
func (s swclient) GetSingleRaw(resource string, id string) (*Response, error) {
	return s.request("GET", resource, id, bytes.NewBufferString(""))
}

// PutSingle
func (s swclient) PutSingle(id string, o interface{}) (*Response, error) {
	if res, ok := s.resources[reflect.TypeOf(o).String()]; ok {
		bts, err := json.Marshal(o)
		if err != nil {
			return nil, swerror{"swclient/swclient.go", "GetSingle()", err.Error()}
		}
		return s.request("PUT", res, id, bytes.NewReader(bts))
	} else {
		return nil, swerror{"swclient/swclient.go", "PutSingle()", "Passed type is not a resource of the shopware api!"}
	}
}

// PutSingleRaw
func (s swclient) PutSingleRaw(resource string, id string, body io.Reader) (*Response, error) {
	return s.request("PUT", resource, id, body)
}

// request executes an http-request of the given method
func (s *swclient) request(method string, resource string, id string, body io.Reader) (*Response, error) {
	// join shopware base-url, api-endpoint, resource and id
	s.baseurl.Path = path.Join(s.apiurl, resource, id)

	// execute
	resp, err := s.dgc.request(method, s.baseurl.String(), body, s.user, s.key, s.hshr)
	if err != nil {
		return nil, swerror{"swclient/swclient.go", "request()", err.Error()}
	}

	if resp.StatusCode != 200 {
		return nil, swerror{"swclient/swclient.go", "request()", resp.Status}
	}
	// read response
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, swerror{"swclient/swclient.go", "request()", err.Error()}
	}
	resp.Body.Close()

	// unmarshal received data into swclient.Response
	data := Response{}
	err = json.Unmarshal(b, &data)
	if err != nil {
		return nil, swerror{"swclient/swclient.go", "request()", err.Error()}
	}

	return &data, nil
}
