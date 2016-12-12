package article

type PropertyValue struct {
	ValueNumeric float64 `json:"valueNumeric,omitempty"`
	Position     int     `json:"position,omitempty"`
	OptionId     int     `json:"optionId,omitempty"`
	Id           int     `json:"id,omitempty"`
	Value        string  `json:"value,omitempty"`
}
