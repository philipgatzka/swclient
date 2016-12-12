package swclient

import (
	"bytes"
	"crypto/md5"
	"encoding/json"
	"fmt"
	"hash"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
	"reflect"
)

// Swclient defines the interface this library exposes.
type Swclient interface {
	Get(id string, o interface{}) error
	GetRaw(resource string, id string) (*Response, error)
	Put(id string, o interface{}) (*Response, error)
	PutRaw(resource string, id string, body io.Reader) (*Response, error)
	Post(o interface{}) (*Response, error)
	PostRaw(resource string, body io.Reader) (*Response, error)
	Delete(resource string, id ...string) (*Response, error)
}

// swclient holds client and server information.
type swclient struct {
	user    string
	key     string
	baseurl *url.URL
	apiurl  string
	dgc     *digestclient
	hshr    hash.Hash
}

// cerror provides custom error messages.
type cerror struct {
	file     string
	function string
	message  string
}

func (s cerror) Error() string {
	return fmt.Sprintf("\n%s, %s: %s", s.file, s.function, s.message)
}

// resources maps typenames to Shopware API-resources/endpoints.
// TODO: get rid of hardcoded limit
var resources = map[string]string{
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
	// TODO: value below gets converted to: articles%3Flimit=999999, 404 Not found response
	"*[]address.Address":           "addresses?limit=999999",
	"*[]article.Article":           "articles?limit=999999",
	"*[]cache.Cache":               "caches?limit=999999",
	"*[]category.Category":         "categories?limit=999999",
	"*[]country.Country":           "countries?limit=999999",
	"*[]customer.Customer":         "customers?limit=999999",
	"*[]manufacturer.Manufacturer": "manufacturers?limit=999999",
	"*[]media.Media":               "media?limit=999999",
	"*[]order.Order":               "orders?limit=999999",
	"*[]shop.Shop":                 "shops?limit=999999",
	"*[]translation.Translation":   "translations?limit=999999",
}

// New returns an initialised swclient.
func New(user string, key string, apiurl string) (Swclient, error) {
	// check input
	if len(user) <= 0 {
		return nil, cerror{"swclient/swclient.go", "New()", "Can't create swclient: user not specified"}
	}

	if len(key) <= 0 {
		return nil, cerror{"swclient/swclient.go", "New()", "Can't create swclient: key not specified"}
	}

	if len(apiurl) <= 0 {
		return nil, cerror{"swclient/swclient.go", "New()", "Can't create swclient: api-url not specified"}
	}

	u, err := url.Parse(apiurl)
	if err != nil {
		return nil, cerror{"swclient/swclient.go", "New()", err.Error()}
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
		hshr: md5.New(),
	}, nil
}

// Get fetches a resource or multiple resources from a shop.
// Resource selection is done by passing a pointer to a struct of one of the types present in swclient.resources.
// Data returned from the shop is unmarshaled into the passed struct.
// Examples:
// 	a := article.Article{}
// 	s.Get("4", &a)	// single
// 	b := []article.Article{}
// 	s.Get("", &b)	// list
func (s swclient) Get(id string, o interface{}) error {
	// BUG(philipgatzka): This check with reflect.TypeOf(o).String() is ~magic~...
	if res, ok := resources[reflect.TypeOf(o).String()]; ok {
		resp, err := s.GetRaw(res, id)
		if err != nil {
			return cerror{"swclient/swclient.go", "GetSingle()", err.Error()}
		}

		err = json.Unmarshal(resp.Data, o)
		if err != nil {
			return cerror{"swclient/swclient.go", "GetSingle()", err.Error()}
		}
	} else {
		return cerror{"swclient/swclient.go", "GetSingle()", reflect.TypeOf(o).String() + " is not a resource of the shopware api!"}
	}
	return nil
}

// GetRaw fetches a resource or multiple resources from a shop.
// Examples:
//	s.GetRaw("articles", "6")	// single
//	s.GetRaw("articles", "")	// list
func (s swclient) GetRaw(resource string, id string) (*Response, error) {
	return s.request("GET", resource, id, bytes.NewBufferString(""))
}

// Put updates a shop resource.
// Resource selection is done by passing a pointer to a struct of one of the types present in swclient.resources.
// Example:
// 	a := article.Article{
// 		Name: "New name"
// 		MainDetail: &article.Detail{
// 			InStock: 78,
// 			Prices: []article.Price{
// 				{
// 					Price: 123.456,
// 				},
// 			},
// 		},
// 	}
// 	s.Put("4", &a)	// single
func (s swclient) Put(id string, o interface{}) (*Response, error) {
	// BUG(philipgatzka): This check with reflect.TypeOf(o).String() is ~magic~...
	if res, ok := resources[reflect.TypeOf(o).String()]; ok {
		bts, err := json.Marshal(o)
		if err != nil {
			return nil, cerror{"swclient/swclient.go", "GetSingle()", err.Error()}
		}
		return s.PutRaw(res, id, bytes.NewReader(bts))
	} else {
		return nil, cerror{"swclient/swclient.go", "PutSingle()", reflect.TypeOf(o).String() + " is not a resource of the shopware api!"}
	}
}

// PutRaw updates a shop resource.
// Example:
//	resp, err := s.PutRaw("articles", "6", bytes.NewBufferString("{Name:"New name"}"))
func (s swclient) PutRaw(resource string, id string, body io.Reader) (*Response, error) {
	return s.request("PUT", resource, id, body)
}

// Post creates a shop resource.
// Resource selection is done by passing a pointer to a struct of one of the types present in swclient.resources.
// Example:
// 	a := article.Article{
// 		Name: "The name"
// 		MainDetail: &article.Detail{
// 			InStock: 78,
// 			Prices: []article.Price{
// 				{
// 					Price: 123.456,
// 				},
// 			},
// 		},
// 	}
// 	s.Post("4", &a)	// single
func (s swclient) Post(o interface{}) (*Response, error) {
	// BUG(philipgatzka): This check with reflect.TypeOf(o).String() is ~magic~...
	if res, ok := resources[reflect.TypeOf(o).String()]; ok {
		bts, err := json.Marshal(o)
		if err != nil {
			return nil, cerror{"swclient/swclient.go", "GetSingle()", err.Error()}
		}
		return s.PostRaw(res, bytes.NewReader(bts))
	} else {
		return nil, cerror{"swclient/swclient.go", "PutSingle()", reflect.TypeOf(o).String() + " is not a resource of the shopware api!"}
	}
}

// PostRaw updates a shop resource.
// Example:
//	resp, err := s.PostRaw("articles", bytes.NewBufferString("{Name:"The name"}"))
func (s swclient) PostRaw(resource string, body io.Reader) (*Response, error) {
	return s.request("POST", resource, "", body)
}

// Delete one or more shop resources.
// Example:
//	resp, err := s.Delete("articles", "6")
func (s swclient) Delete(resource string, id ...string) (*Response, error) {
	if len(id) > 0 {
		if len(id) > 1 {
			bts, err := json.Marshal(id)
			if err != nil {
				return nil, cerror{"swylient/swclient.go", "Delete()", err.Error()}
			}
			return s.request("DELETE", resource, "", bytes.NewReader(bts))
		} else {
			return s.request("DELETE", resource, id[0], bytes.NewBufferString(""))
		}
	}
	return nil, cerror{"swclient/swclient.go", "Delete()", "No id given!"}
}

// request executes an http-request.
func (s *swclient) request(method string, resource string, id string, body io.Reader) (*Response, error) {
	// join base-url, resource and id
	s.baseurl.Path = path.Join(s.apiurl, resource, id)

	// execute
	resp, err := s.dgc.request(method, s.baseurl.String(), body, s.user, s.key, s.hshr)
	if err != nil {
		return nil, cerror{"swclient/swclient.go", "request()", err.Error()}
	}

	if !(resp.StatusCode == http.StatusOK || resp.StatusCode == http.StatusCreated) {
		return nil, cerror{"swclient/swclient.go", "request()", resp.Status}
	}
	// read response
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, cerror{"swclient/swclient.go", "request()", err.Error()}
	}
	resp.Body.Close()

	// unmarshal received data into swclient.Response
	data := Response{}
	err = json.Unmarshal(b, &data)
	if err != nil {
		return nil, cerror{"swclient/swclient.go", "request()", err.Error()}
	}

	return &data, nil
}
