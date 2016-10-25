package category

import "encoding/json"

type Category struct {
	Id               int                `json:",omitempty"`
	ParentId         int                `json:",omitempty"`
	StreamId         int                `json:",omitempty"`
	Name             string             `json:",omitempty"`
	Position         int                `json:",omitempty"`
	MetaTitle        string             `json:",omitempty"`
	MetaKeywords     string             `json:",omitempty"`
	MetaDescription  string             `json:",omitempty"`
	CmsHeadline      string             `json:",omitempty"`
	CmsText          string             `json:",omitempty"`
	Active           bool               `json:",omitempty"`
	Template         string             `json:",omitempty"`
	ProductBoxLayout string             `json:",omitempty"`
	Blog             bool               `json:",omitempty"`
	Path             string             `json:",omitempty"`
	ShowFilterGroups bool               `json:",omitempty"`
	External         bool               `json:",omitempty"`
	HideFilter       bool               `json:",omitempty"`
	HideTop          bool               `json:",omitempty"`
	Changed          string             `json:",omitempty"`
	Added            string             `json:",omitempty"`
	MediaId          int                `json:",omitempty"`
	Attribute        *[]json.RawMessage `json:",omitempty"`
	Emotions         *[]json.RawMessage `json:",omitempty"`
	Media            *json.RawMessage   `json:",omitempty"`
	CustomerGroups   *[]json.RawMessage `json:",omitempty"`
	ChildrenCount    string             `json:",omitempty"`
	ArticleCount     string             `json:",omitempty"`
}
