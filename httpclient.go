package swclient

import (
	"bytes"
	"net/http"
)

type httpGetter interface {
	Get(uri string) (*http.Response, error)
}

func Get(uri string) (*http.Response, error) {
	request, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		return nil, err
	}

	return client.Do(request)
}

func Post(uri string, payload string) (*http.Response, error) {
	request, err := http.NewRequest("POST", uri, bytes.NewBuffer([]byte(payload)))
	if err != nil {
		return nil, err
	}

	return client.Do(request)
}
