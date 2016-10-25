package order

type OrderStatus struct {
	Id          int    `json:",omitempty"`
	Description string `json:",omitempty"`
	Position    int    `json:",omitempty"`
	Group       string `json:",omitempty"`
	SendMail    bool   `json:",omitempty"`
}
