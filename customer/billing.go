package customer

// TODO
type Billing struct {
	id           int
	customerId   int
	countryId    int
	stateId      int
	company      string
	department   string
	salutation   string
	number       string
	firstName    string
	lastName     string
	street       string
	streetNumber string
	zipCode      string
	city         string
	phone        string
	fax          string
	vatId        string
	birthday     string
	attribute    *BillingAttribute
}
