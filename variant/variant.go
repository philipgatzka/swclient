package variant

type Variant struct {
	Id                  int                   `json:",omitempty"`
	ArticleId           int                   `json:",omitempty"`
	UnitId              int                   `json:",omitempty"`
	Number              string                `json:",omitempty"`
	SupplierNumber      string                `json:",omitempty"`
	Kind                int                   `json:",omitempty"`
	AdditionalText      string                `json:",omitempty"`
	Active              bool                  `json:",omitempty"`
	InStock             int                   `json:",omitempty"`
	StockMin            int                   `json:",omitempty"`
	Weight              string                `json:",omitempty"`
	Len                 string                `json:",omitempty"`
	Height              string                `json:",omitempty"`
	Ean                 string                `json:",omitempty"`
	Position            int                   `json:",omitempty"`
	MinPurchase         int                   `json:",omitempty"`
	PurchaseSteps       int                   `json:",omitempty"`
	MaxPurchase         int                   `json:",omitempty"`
	PurchaseUnit        string                `json:",omitempty"`
	ShippingFree        bool                  `json:",omitempty"`
	ReleaseDate         string                `json:",omitempty"`
	ShippingTime        string                `json:",omitempty"`
	Prices              *[]Price              `json:",omitempty"`
	Attribute           *Attribute            `json:",omitempty"`
	ConfiguratorOptions *[]ConfiguratorOption `json:",omitempty"`
}
