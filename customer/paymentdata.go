package customer

type PaymentData struct {
	Id             int    `json:",omitempty"`
	PaymentMeanId  int    `json:",omitempty"`
	UseBillingData string `json:",omitempty"`
	BankName       string `json:",omitempty"`
	Bic            string `json:",omitempty"`
	Iban           string `json:",omitempty"`
	AccountNumber  string `json:",omitempty"`
	BankCode       string `json:",omitempty"`
	AccountHolder  string `json:",omitempty"`
	CreatedAt      string `json:",omitempty"`
}
