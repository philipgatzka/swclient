package order

type OrderStatus struct {
	Id          int    `json:",omitempty"`
	Name        string `json:",omitempty"`
	Description string `json:",omitempty"`
	Position    int    `json:",omitempty"`
	Group       string `json:",omitempty"`
	SendMail    int    `json:",omitempty"`
}
