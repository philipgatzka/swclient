package shop

type Shop struct {
	Id             int       `json:",omitempty"`
	MainId         int       `json:",omitempty"`
	CategoryId     int       `json:",omitempty"`
	Name           string    `json:",omitempty"`
	Title          string    `json:",omitempty"`
	Position       int       `json:",omitempty"`
	Host           string    `json:",omitempty"`
	BasePath       string    `json:",omitempty"`
	BaseUrl        string    `json:",omitempty"`
	Hosts          string    `json:",omitempty"`
	Secure         bool      `json:",omitempty"`
	AlwaysSecure   bool      `json:",omitempty"`
	SecureHost     string    `json:",omitempty"`
	SecureBasePath string    `json:",omitempty"`
	Def            bool      `json:"default,omitempty"`
	Active         bool      `json:",omitempty"`
	CustomerScope  bool      `json:",omitempty"`
	Currency       *Currency `json:",omitempty"`
}
