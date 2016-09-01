package swclient

import (
	"bytes"
	"net/http"
)

func get(uri string) (*http.Response, error) {
	request, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		return nil, err
	}

	client := http.Client{}

	return client.Do(request)
}

func post(uri string, payload string) (*http.Response, error) {
	request, err := http.NewRequest("POST", uri, bytes.NewBuffer([]byte(payload)))
	if err != nil {
		return nil, err
	}

	client := http.Client{}

	return client.Do(request)
}
