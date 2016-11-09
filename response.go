package swclient

import "encoding/json"

/*
	Response represents a response from shopware
*/
type Response struct {
	Data    json.RawMessage
	Success bool
	Total   int
}
