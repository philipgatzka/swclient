package article

type CustomerGroup struct {
	Id                    int     `json:"id,omitempty"`
	Key                   string  `json:"key,omitempty"`
	Name                  string  `json:"name,omitempty"`
	Tax                   bool    `json:"tax,omitempty"`
	TaxInput              bool    `json:"taxInput,omitempty"`
	Mode                  bool    `json:"mode,omitempty"`
	Discount              float64 `json:"discount,omitempty"`
	MinimumOrder          float64 `json:"minimumOrder,omitempty"`
	MinimumOrderSurcharge float64 `json:"minimumOrderSurcharge,omitempty"`
	BasePrice             float64 `json:"basePrice,omitempty"`
	Percent               float64 `json:"percent,omitempty"`
}
