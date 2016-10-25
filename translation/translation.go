package translation

import "encoding/json"

type Translation struct {
	typ    string             `json:"type,omitempty"`
	data   *[]json.RawMessage `json:",omitempty"`
	key    int                `json:",omitempty"`
	shopId int                `json:",omitempty"`
	shop   *Shop              `json:",omitempty"`
}
