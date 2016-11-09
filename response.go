package swclient

import "encoding/json"

// Response defines a response from shopware
type Response struct {
	Data    json.RawMessage
	Success bool
	Total   int
}
