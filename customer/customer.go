package customer

type Customer struct {
	Id              int                `json:",omitempty"`
	PaymentId       int                `json:",omitempty"`
	GroupKey        string             `json:",omitempty"`
	ShopId          string             `json:",omitempty"`
	PriceGroupId    int                `json:",omitempty"`
	EncoderName     string             `json:",omitempty"`
	HashPassword    string             `json:",omitempty"`
	Active          bool               `json:",omitempty"`
	Email           string             `json:",omitempty"`
	FirstLogin      string             `json:",omitempty"`
	LastLogin       string             `json:",omitempty"`
	AccountMode     int                `json:",omitempty"`
	ConfirmationKey string             `json:",omitempty"`
	SessionId       string             `json:",omitempty"`
	Newsletter      bool               `json:",omitempty"`
	Validation      string             `json:",omitempty"`
	Affiliate       bool               `json:",omitempty"`
	PaymentPreset   int                `json:",omitempty"`
	LanguageId      int                `json:",omitempty"`
	Referer         string             `json:",omitempty"`
	InternalComment string             `json:",omitempty"`
	FailedLogins    int                `json:",omitempty"`
	LockedUntil     string             `json:",omitempty"`
	Attribute       *CustomerAttribute `json:",omitempty"`
	Billing         *Billing           `json:",omitempty"`
	PaymentData     *[]PaymentData     `json:",omitempty"`
	Shipping        *Shipping          `json:",omitempty"`
	Debit           *Debit             `json:",omitempty"`
}
