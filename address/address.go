package address

type Address struct {
	Id                     int      `json:",omitempty"`
	Company                string   `json:",omitempty"`
	Department             string   `json:",omitempty"`
	Salutation             string   `json:",omitempty"`
	Firstname              string   `json:",omitempty"`
	Lastname               string   `json:",omitempty"`
	Street                 string   `json:",omitempty"`
	Zipcode                string   `json:",omitempty"`
	City                   string   `json:",omitempty"`
	Phone                  string   `json:",omitempty"`
	VatId                  string   `json:",omitempty"`
	AdditionalAddressLine1 string   `json:",omitempty"`
	AdditionalAddressLine2 string   `json:",omitempty"`
	Country                int      `json:",omitempty"`
	State                  int      `json:",omitempty"`
	Attribute              []string `json:",omitempty"`
}
