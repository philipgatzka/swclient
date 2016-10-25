package country

type Country struct {
	Id                         int     `json:",omitempty"`
	Name                       string  `json:",omitempty"`
	Iso                        string  `json:",omitempty"`
	IsoName                    string  `json:",omitempty"`
	Position                   int     `json:",omitempty"`
	Description                string  `json:",omitempty"`
	ShippingFree               bool    `json:",omitempty"`
	TaxFree                    bool    `json:",omitempty"`
	TaxFreeUstId               bool    `json:",omitempty"`
	TaxFreeUstIdChecked        bool    `json:",omitempty"`
	Active                     bool    `json:",omitempty"`
	Iso3                       string  `json:",omitempty"`
	DisplayStateInRegistration bool    `json:",omitempty"`
	ForceStateInRegistration   bool    `json:",omitempty"`
	AreaId                     int     `json:",omitempty"`
	States                     []State `json:",omitempty"`
}
