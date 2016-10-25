package customergroup

import "encoding/json"

type CustomerGroup struct {
	Id                    int                `json:",omitempty"`
	Key                   string             `json:",omitempty"`
	Name                  string             `json:",omitempty"`
	Tax                   bool               `json:",omitempty"`
	TaxInput              bool               `json:",omitempty"`
	Mode                  bool               `json:",omitempty"`
	Discount              int                `json:",omitempty"`
	MinimumOrder          int                `json:",omitempty"`
	MinimumOrderSurcharge int                `json:",omitempty"`
	Discounts             []*json.RawMessage `json:",omitempty"`
}
