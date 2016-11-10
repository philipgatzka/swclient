package swclient

import "encoding/json"

// Response represents a response from the Shopware API.
type Response struct {
	Data    json.RawMessage
	Success bool
	Total   int
}
