package swclient

import (
	"net/http"
)

type httpclient struct {
	d digest
	c http.Client
	// TODO: serverinfo as func?
}

func (h *httpclient) call() {

}
