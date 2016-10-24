package article

type Detail struct {
	Number              string               `json:",omitempty"`
	SupplierNumber      string               `json:",omitempty"`
	AdditionalText      string               `json:",omitempty"`
	Weight              string               `json:",omitempty"`
	Width               string               `json:",omitempty"`
	Len                 string               `json:",omitempty"`
	Height              string               `json:",omitempty"`
	Ean                 string               `json:",omitempty"`
	PurchaseUnit        string               `json:",omitempty"`
	DescriptionLong     string               `json:",omitempty"`
	ReferenceUnit       string               `json:",omitempty"`
	PackUnit            string               `json:",omitempty"`
	ShippingTime        string               `json:",omitempty"`
	Prices              []Price              `json:",omitempty"`
	ConfiguratorOptions []ConfiguratorOption `json:",omitempty"`
	Attribute           *Attribute           `json:",omitempty"`
	Id                  int                  `json:",omitempty"`
	ArticleId           int                  `json:",omitempty"`
	UnitId              int                  `json:",omitempty"`
	Kind                int                  `json:",omitempty"`
	InStock             int                  `json:",omitempty"`
	Position            int                  `json:",omitempty"`
	MinPurchase         int                  `json:",omitempty"`
	PurchaseSteps       int                  `json:",omitempty"`
	MaxPurchase         int                  `json:",omitempty"`
	ReleaseDate         string               `json:",omitempty"`
	Active              bool                 `json:",omitempty"`
	ShippingFree        bool                 `json:",omitempty"`
}
