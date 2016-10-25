package manufacturers

import "encoding/json"

type Manufacturers struct {
	Id              int                `json:",omitempty"`
	Name            string             `json:",omitempty"`
	Image           string             `json:",omitempty"`
	Link            string             `json:",omitempty"`
	Description     string             `json:",omitempty"`
	MetaTitle       string             `json:",omitempty"`
	MetaKeywords    string             `json:",omitempty"`
	MetaDescription string             `json:",omitempty"`
	Attribute       *[]json.RawMessage `json:",omitempty"`
}
