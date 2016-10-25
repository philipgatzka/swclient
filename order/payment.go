package order

import "encoding/json"

type Payment struct {
	Id                    int              `json:",omitempty"`
	Name                  string           `json:",omitempty"`
	Description           string           `json:",omitempty"`
	Template              string           `json:",omitempty"`
	Class                 string           `json:",omitempty"`
	Table                 string           `json:",omitempty"`
	Hide                  bool             `json:",omitempty"`
	AdditionalDescription string           `json:",omitempty"`
	DebitPercent          int              `json:",omitempty"`
	Surcharge             int              `json:",omitempty"`
	SurchargeString       string           `json:",omitempty"`
	Position              int              `json:",omitempty"`
	Active                bool             `json:",omitempty"`
	EsdActive             bool             `json:",omitempty"`
	MobileInactive        bool             `json:",omitempty"`
	EmbedIFrame           string           `json:",omitempty"`
	HideProspect          int              `json:",omitempty"`
	Action                string           `json:",omitempty"`
	PluginId              *json.RawMessage `json:",omitempty"`
	Source                *json.RawMessage `json:",omitempty"`
	Attribute             *Attribute       `json:",omitempty"`
}
