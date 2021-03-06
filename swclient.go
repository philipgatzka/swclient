package swclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
	"reflect"
)

// Response represents a response from the Shopware API.
type Response struct {
	Data    json.RawMessage
	Message string
	Success bool
	Total   int
}

// Swclient defines the interface this library exposes.
type Swclient interface {
	Get(id string, o interface{}) error
	GetRaw(resource string, id string) (*Response, error)
	Put(id string, o interface{}) (*Response, error)
	PutRaw(resource string, id string, body io.Reader) (*Response, error)
	Post(o interface{}) (*Response, error)
	PostRaw(resource string, body io.Reader) (*Response, error)
	RequestRaw(method, uri string, body io.Reader) (*http.Response, error)
	Delete(resource string, id ...string) (*Response, error)
}

// swclient holds client and server information.
type swclient struct {
	user    string
	key     string
	baseurl *url.URL
	apiurl  string
	dgc     *digestclient
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
var resources = map[string]string{
	"*address.Address":             "addresses",
	"*article.Article":             "articles",
	"*cache.Cache":                 "caches",
	"*category.Category":           "categories",
	"*country.Country":             "countries",
	"*customer.Customer":           "customers",
	"*manufacturer.Manufacturer":   "manufacturers",
	"*media.Media":                 "media",
	"*order.Order":                 "orders",
	"*shop.Shop":                   "shops",
	"*translation.Translation":     "translations",
	"*variant.Variant":             "variants",
	"*[]address.Address":           "addresses",
	"*[]article.Article":           "articles",
	"*[]cache.Cache":               "caches",
	"*[]category.Category":         "categories",
	"*[]country.Country":           "countries",
	"*[]customer.Customer":         "customers",
	"*[]manufacturer.Manufacturer": "manufacturers",
	"*[]media.Media":               "media",
	"*[]order.Order":               "orders",
	"*[]shop.Shop":                 "shops",
	"*[]translation.Translation":   "translations",
	"*article.Articles":            "articles",
	"*manufacturer.Manufacturers":  "manufacturers",
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
			return cerror{"swclient/swclient.go", "Get() - GetRaw()", err.Error()}
		}

		err = json.Unmarshal(resp.Data, o)
		if err != nil {
			return cerror{"swclient/swclient.go", "Get() - json.Unmarshal()", err.Error()}
		}
	} else {
		return cerror{"swclient/swclient.go", "Get()", reflect.TypeOf(o).String() + " is not a resource of the shopware api!"}
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
	// BUG: This check with reflect.TypeOf(o).String() is ~magic~...
	if res, ok := resources[reflect.TypeOf(o).String()]; ok {
		bts, err := json.Marshal(o)
		if err != nil {
			return nil, cerror{"swclient/swclient.go", "GetSingle()", err.Error()}
		}
		return s.PutRaw(res, id, bytes.NewReader(bts))
	}
	return nil, cerror{"swclient/swclient.go", "PutSingle()", reflect.TypeOf(o).String() + " is not a resource of the shopware api!"}
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
	// BUG: This check with reflect.TypeOf(o).String() is ~magic~...
	if res, ok := resources[reflect.TypeOf(o).String()]; ok {
		bts, err := json.Marshal(o)
		if err != nil {
			return nil, cerror{"swclient/swclient.go", "GetSingle()", err.Error()}
		}
		return s.PostRaw(res, bytes.NewReader(bts))
	}
	return nil, cerror{"swclient/swclient.go", "PutSingle()", reflect.TypeOf(o).String() + " is not a resource of the shopware api!"}
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
			ids := []interface{}{}
			for _, i := range id {
				ids = append(ids, struct {
					Id string `json:"id"`
				}{Id: i})
			}
			bts, err := json.Marshal(ids)
			if err != nil {
				return nil, cerror{"swylient/swclient.go", "Delete()", err.Error()}
			}
			return s.request("DELETE", resource, "", bytes.NewReader(bts))
		}
		return s.request("DELETE", resource, id[0], bytes.NewBufferString(""))
	}
	return nil, cerror{"swclient/swclient.go", "Delete()", "No id given!"}
}

// request executes an http-request.
func (s *swclient) request(method string, resource string, id string, body io.Reader) (*Response, error) {
	// join base-url, resource and id
	s.baseurl.Path = path.Join(s.apiurl, resource, id)

	// execute
	resp, err := s.dgc.request(method, s.baseurl.String(), body, s.user, s.key)
	if err != nil {
		return nil, cerror{"swclient/swclient.go", "request() - s.dgc.request()", err.Error()}
	}
	// read response
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, cerror{"swclient/swclient.go", "request() - ioutil.ReadAll()", err.Error()}
	}
	resp.Body.Close()

	// check if response status is OK
	if !(resp.StatusCode == http.StatusOK || resp.StatusCode == http.StatusCreated) {
		return nil, cerror{"swclient/swclient.go", "request()", fmt.Sprintf("%s:\n%s", resp.Status, string(b))}
	}

	// unmarshal received data into swclient.Response
	data := Response{}
	json.Unmarshal(b, &data)
	return &data, nil
}

// RequestRaw executes a customizable request via the Swclients Digest-HTTP-Client
func (s swclient) RequestRaw(method, uri string, body io.Reader) (*http.Response, error) {
	return s.dgc.request(method, uri, body, s.user, s.key)
}
