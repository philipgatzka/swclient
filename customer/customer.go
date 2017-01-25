package customer

import (
	"encoding/json"
)

type Customer struct {
	ID              int          `json:"id,omitempty"`
	PaymentID       int          `json:"paymentId,omitempty"`
	GroupKey        string       `json:"groupKey,omitempty"`
	ShopID          string       `json:"shopId,omitempty"`
	PriceGroupID    int          `json:"priceGroupId,omitempty"`
	EncoderName     string       `json:"encoderName,omitempty"`
	HashPassword    string       `json:"hashPassword,omitempty"`
	RawPassword     string       `json:"rawPassword,omitempty"`
	Active          bool         `json:"active,omitempty"`
	Email           string       `json:"email,omitempty"`
	FirstLogin      string       `json:"firstLogin,omitempty"`
	LastLogin       string       `json:"lastLogin,omitempty"`
	AccountMode     int          `json:"accountMode,omitempty"`
	ConfirmationKey string       `json:"confirmationKey,omitempty"`
	SessionID       string       `json:"sessionId,omitempty"`
	Newsletter      bool         `json:"newsletter,omitempty"`
	Validation      string       `json:"validation,omitempty"`
	Affiliate       bool         `json:"affiliate,omitempty"`
	PaymentPreset   int          `json:"paymentPreset,omitempty"`
	LanguageID      int          `json:"languageId,omitempty"`
	Referer         string       `json:"referer,omitempty"`
	InternalComment string       `json:"internalComment,omitempty"`
	FailedLogins    int          `json:"failedLogins,omitempty"`
	LockedUntil     string       `json:"lockedUntil,omitempty"`
	Attribute       *Attribute   `json:"attribute,omitempty"`
	Billing         *Billing     `json:"billing,omitempty"`
	PaymentData     *PaymentData `json:"paymentData,omitempty"`
	Shipping        *Shipping    `json:"shipping,omitempty"`
	Debit           *Debit       `json:"debit,omitempty"`
}

func New(email, firstname, lastname, salutation, street, streetNumber, city, zipcode string, countryID int) (Customer, error) {
	return Customer{
		Email: email,
		Billing: &Billing{
			FirstName:    firstname,
			LastName:     lastname,
			Salutation:   salutation,
			Street:       street,
			StreetNumber: streetNumber,
			City:         city,
			ZipCode:      zipcode,
			Country:      countryID,
		},
	}, nil
}

type Customers []Customer

// Len implements the sort.Interface
func (c Customers) Len() int {
	return len(c)
}

// Less implements the sort.Interface
func (c Customers) Less(i, j int) bool {
	return c[i].Email < c[j].Email
}

// Swap implements the sort.Interface
func (c Customers) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

// String implements the Stringer interface. Returns the articles Number and Name
func (c Customer) String() string {
	return c.Email + ": " + c.Billing.FirstName + " " + c.Billing.LastName
}

// MarshalJSON translates a customer into JSON.
// This is necessary, because the Shopware API returns a slightly different object on GETting than it expects on POSTing.
func (c Customer) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		ID              int          `json:"id,omitempty"`
		FirstName       string       `json:"firstname,omitempty"`
		LastName        string       `json:"lastname,omitempty"`
		Salutation      string       `json:"salutation"`
		PaymentID       int          `json:"paymentId,omitempty"`
		GroupKey        string       `json:"groupKey,omitempty"`
		ShopID          string       `json:"shopId,omitempty"`
		PriceGroupID    int          `json:"priceGroupId,omitempty"`
		EncoderName     string       `json:"encoderName,omitempty"`
		HashPassword    string       `json:"hashPassword,omitempty"`
		RawPassword     string       `json:"rawPassword,omitempty"`
		Active          bool         `json:"active,omitempty"`
		Email           string       `json:"email,omitempty"`
		FirstLogin      string       `json:"firstLogin,omitempty"`
		LastLogin       string       `json:"lastLogin,omitempty"`
		AccountMode     int          `json:"accountMode,omitempty"`
		ConfirmationKey string       `json:"confirmationKey,omitempty"`
		SessionID       string       `json:"sessionId,omitempty"`
		Newsletter      bool         `json:"newsletter,omitempty"`
		Validation      string       `json:"validation,omitempty"`
		Affiliate       bool         `json:"affiliate,omitempty"`
		PaymentPreset   int          `json:"paymentPreset,omitempty"`
		LanguageID      int          `json:"languageId,omitempty"`
		Referer         string       `json:"referer,omitempty"`
		InternalComment string       `json:"internalComment,omitempty"`
		FailedLogins    int          `json:"failedLogins,omitempty"`
		LockedUntil     string       `json:"lockedUntil,omitempty"`
		Attribute       *Attribute   `json:"attribute,omitempty"`
		Billing         *Billing     `json:"billing,omitempty"`
		PaymentData     *PaymentData `json:"paymentData,omitempty"`
		Shipping        *Shipping    `json:"shipping,omitempty"`
		Debit           *Debit       `json:"debit,omitempty"`
	}{
		ID:              c.ID,
		FirstName:       c.Billing.FirstName,
		LastName:        c.Billing.LastName,
		Salutation:      c.Billing.Salutation,
		PaymentID:       c.PaymentID,
		GroupKey:        c.GroupKey,
		ShopID:          c.ShopID,
		PriceGroupID:    c.PriceGroupID,
		EncoderName:     c.EncoderName,
		HashPassword:    c.HashPassword,
		RawPassword:     c.RawPassword,
		Active:          c.Active,
		Email:           c.Email,
		FirstLogin:      c.FirstLogin,
		LastLogin:       c.LastLogin,
		AccountMode:     c.AccountMode,
		ConfirmationKey: c.ConfirmationKey,
		SessionID:       c.SessionID,
		Newsletter:      c.Newsletter,
		Validation:      c.Validation,
		Affiliate:       c.Affiliate,
		PaymentPreset:   c.PaymentPreset,
		LanguageID:      c.LanguageID,
		Referer:         c.Referer,
		InternalComment: c.InternalComment,
		FailedLogins:    c.FailedLogins,
		LockedUntil:     c.LockedUntil,
		Attribute:       c.Attribute,
		Billing:         c.Billing,
		PaymentData:     c.PaymentData,
		Shipping:        c.Shipping,
		Debit:           c.Debit,
	})
}
