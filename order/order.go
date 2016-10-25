package order

import (
	"encoding/json"
)

type Order struct {
	Id                 int              `json:",omitempty"`
	Number             string           `json:",omitempty"`
	CustomerId         int              `json:",omitempty"`
	PaymentId          int              `json:",omitempty"`
	DispatchId         int              `json:",omitempty"`
	PartnerId          string           `json:",omitempty"`
	ShopId             int              `json:",omitempty"`
	InvoiceAmount      float64          `json:",omitempty"`
	InvoiceAmountNet   float64          `json:",omitempty"`
	InvoiceShipping    float64          `json:",omitempty"`
	InvoiceShippingNet float64          `json:",omitempty"`
	OrderTime          string           `json:",omitempty"`
	TransactionId      string           `json:",omitempty"`
	Comment            string           `json:",omitempty"`
	CustomerComment    string           `json:",omitempty"`
	InternalComment    string           `json:",omitempty"`
	Net                int              `json:",omitempty"`
	TaxFree            int              `json:",omitempty"`
	TemporaryId        string           `json:",omitempty"`
	Referer            string           `json:",omitempty"`
	ClearedDate        string           `json:",omitempty"`
	TrackingCode       string           `json:",omitempty"`
	LanguageIso        string           `json:",omitempty"`
	Currency           string           `json:",omitempty"`
	CurrencyFactor     float64          `json:",omitempty"`
	RemoteAddress      string           `json:",omitempty"`
	DeviceType         string           `json:",omitempty"`
	Details            *[]Detail        `json:",omitempty"`
	Documents          *[]Document      `json:",omitempty"`
	Payment            *Payment         `json:",omitempty"`
	PaymentStatus      *json.RawMessage `json:",omitempty"`
	OrderStatus        *OrderStatus     `json:",omitempty"`
	Customer           *json.RawMessage `json:",omitempty"`
	PaymentInstances   *[]Payment       `json:",omitempty"`
	Billing            *Billing         `json:",omitempty"`
	Shipping           *Shipping        `json:",omitempty"`
	Shop               *Shop            `json:",omitempty"`
	Dispatch           *Dispatch        `json:",omitempty"`
	Attribute          *Attribute       `json:",omitempty"`
	LanguageSubShop    *Shop            `json:",omitempty"`
	PaymentStatusId    int              `json:",omitempty"`
	OrderStatusId      int              `json:",omitempty"`
}
