package customer

type Debit struct {
	Id            int    `json:",omitempty"`
	CustomerId    int    `json:",omitempty"`
	Account       string `json:",omitempty"`
	BankCode      string `json:",omitempty"`
	BankName      string `json:",omitempty"`
	AccountHolder string `json:",omitempty"`
}
