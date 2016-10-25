package order

type Shipping struct {
	Id           int                `json:",omitempty"`
	CustomerId   int                `json:",omitempty"`
	CountryId    int                `json:",omitempty"`
	StateId      int                `json:",omitempty"`
	Company      string             `json:",omitempty"`
	Department   string             `json:",omitempty"`
	Salutation   string             `json:",omitempty"`
	Number       string             `json:",omitempty"`
	FirstName    string             `json:",omitempty"`
	LastName     string             `json:",omitempty"`
	Street       string             `json:",omitempty"`
	StreetNumber string             `json:",omitempty"`
	ZipCode      string             `json:",omitempty"`
	City         string             `json:",omitempty"`
	Attribute    *ShippingAttribute `json:",omitempty"`
}
