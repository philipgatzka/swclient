package variant

type CustomerGroup struct {
	Id                    int     `json:",omitempty"`
	Key                   string  `json:",omitempty"`
	Name                  string  `json:",omitempty"`
	Tax                   bool    `json:",omitempty"`
	TaxInput              bool    `json:",omitempty"`
	Mode                  bool    `json:",omitempty"`
	Discount              float64 `json:",omitempty"`
	MinimumOrder          float64 `json:",omitempty"`
	MinimumOrderSurcharge float64 `json:",omitempty"`
	BasePrice             float64 `json:",omitempty"`
	Percent               float64 `json:",omitempty"`
}
