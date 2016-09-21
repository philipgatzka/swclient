package swclient

import (
	"crypto/md5"
	"errors"
	"hash"
	"io"
	"io/ioutil"
	"net/http"
)

type swclient struct {
	user    string
	key     string
	shopurl string
	httpc   *httpclient
	hshr    hash.Hash
}

func New(user string, key string, shopurl string) *swclient {
	return &swclient{
		user:    user,
		key:     key,
		shopurl: shopurl,
		httpc: &httpclient{
			dgst:  &digest{},
			httpc: &http.Client{},
		},
		hshr: md5.New(),
	}
}

func (s swclient) Get(uri string) (string, error) {
	resp, err := s.httpc.get(s.constructUri(uri), s.user, s.key, s.hshr)
	if err != nil {
		return "", err
	}

	return responseString(resp)
}

func (s swclient) Put(uri string, body io.Reader) (string, error) {
	resp, err := s.httpc.put(s.constructUri(uri), body, s.user, s.key, s.hshr)
	if err != nil {
		return "", err
	}

	return responseString(resp)
}

func (s swclient) Post(uri string, body io.Reader) (string, error) {
	resp, err := s.httpc.post(s.constructUri(uri), body, s.user, s.key, s.hshr)
	if err != nil {
		return "", err
	}

	return responseString(resp)
}

func (s swclient) Del(uri string) (string, error) {
	resp, err := s.httpc.del(s.constructUri(uri), s.user, s.key, s.hshr)
	if err != nil {
		return "", err
	}

	return responseString(resp)
}

func (s swclient) constructUri(uri string) string {
	return s.shopurl + uri
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
