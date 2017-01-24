package customer

type Billing struct {
	ID           int               `json:"id"`
	CustomerID   int               `json:"customerId"`
	Country      int               `json:"country"`
	CountryID    int               `json:"countryId"`
	StateID      int               `json:"stateId"`
	Company      string            `json:"company"`
	Department   string            `json:"department"`
	Salutation   string            `json:"salutation"`
	Number       string            `json:"number"`
	FirstName    string            `json:"firstName"`
	LastName     string            `json:"lastName"`
	Street       string            `json:"street"`
	StreetNumber string            `json:"streetNumber"`
	ZipCode      string            `json:"zipCode"`
	City         string            `json:"city"`
	Phone        string            `json:"phone"`
	Fax          string            `json:"fax"`
	VatID        string            `json:"vatId"`
	Birthday     string            `json:"birthday"`
	Attribute    *BillingAttribute `json:"attribute"`
}
