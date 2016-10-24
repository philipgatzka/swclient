package article

type PropertyValue struct {
	ValueNumeric float64 `json:",omitempty"`
	Position     int     `json:",omitempty"`
	OptionId     int     `json:",omitempty"`
	Id           int     `json:",omitempty"`
	Value        string  `json:",omitempty"`
}
