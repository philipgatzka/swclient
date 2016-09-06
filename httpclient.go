package swclient

import (
	"bytes"
	"net/http"
)

func get(client *http.Client, uri string) (*http.Response, error) {
	request, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		return nil, err
	}

	return client.Do(request)
}

func post(client *http.Client, uri string, payload string) (*http.Response, error) {
	request, err := http.NewRequest("POST", uri, bytes.NewBuffer([]byte(payload)))
	if err != nil {
		return nil, err
	}

	return client.Do(request)
}
