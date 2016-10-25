package shop

type Currency struct {
	Id             int     `json:",omitempty"`
	Currency       string  `json:",omitempty"`
	Name           string  `json:",omitempty"`
	Def            int     `json:"default,omitempty"`
	Factor         float64 `json:",omitempty"`
	Symbol         string  `json:",omitempty"`
	SymbolPosition int     `json:",omitempty"`
	Position       int     `json:",omitempty"`
}
