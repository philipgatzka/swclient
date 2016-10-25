package order

type Payment struct {
	Id            int    `json:",omitempty"`
	FirstName     string `json:",omitempty"`
	LastName      string `json:",omitempty"`
	Address       string `json:",omitempty"`
	ZipCode       string `json:",omitempty"`
	City          string `json:",omitempty"`
	BankName      string `json:",omitempty"`
	BankCode      string `json:",omitempty"`
	AccountNumber string `json:",omitempty"`
	AccountHolder string `json:",omitempty"`
	Bic           string `json:",omitempty"`
	Iban          string `json:",omitempty"`
	Amount        string `json:",omitempty"`
	CreatedAt     string `json:",omitempty"`
}
