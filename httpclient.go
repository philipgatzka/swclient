package main

import "net/http"

func get(uri string) (*http.Response, error) {
	request, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		return nil, err
	}

	client := http.Client{}

	return client.Do(request)
}
