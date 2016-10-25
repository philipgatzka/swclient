package order

type Document struct {
	Id         int                `json:",omitempty"`
	Date       string             `json:",omitempty"`
	TypeId     int                `json:",omitempty"`
	CustomerId int                `json:",omitempty"`
	OrderId    int                `json:",omitempty"`
	Amount     float64            `json:",omitempty"`
	DocumentId int                `json:",omitempty"`
	Hash       string             `json:",omitempty"`
	Typ        *DocumentType      `json:"type,omitempty"`
	Attribute  *DocumentAttribute `json:",omitempty"`
}
