package article

type Detail struct {
	Number              string               `json:"number,omitempty"`
	SupplierNumber      string               `json:"supplierNumber,omitempty"`
	AdditionalText      string               `json:"additionalText,omitempty"`
	Weight              string               `json:"weight,omitempty"`
	Width               string               `json:"width,omitempty"`
	Len                 string               `json:"len,omitempty"`
	Height              string               `json:"height,omitempty"`
	Ean                 string               `json:"ean,omitempty"`
	PurchaseUnit        string               `json:"purchaseUnit,omitempty"`
	DescriptionLong     string               `json:"descriptionLong,omitempty"`
	ReferenceUnit       string               `json:"referenceUnit,omitempty"`
	PackUnit            string               `json:"packUnit,omitempty"`
	ShippingTime        string               `json:"shippingTime,omitempty"`
	Prices              []Price              `json:"prices,omitempty"`
	ConfiguratorOptions []ConfiguratorOption `json:"configuratorOptions,omitempty"`
	Attribute           *Attribute           `json:"attribute,omitempty"`
	Id                  int                  `json:"id,omitempty"`
	ArticleId           int                  `json:"articleId,omitempty"`
	UnitId              int                  `json:"unitId,omitempty"`
	Kind                int                  `json:"kind,omitempty"`
	InStock             int                  `json:"inStock,omitempty"`
	Position            int                  `json:"position,omitempty"`
	MinPurchase         int                  `json:"minPurchase,omitempty"`
	PurchaseSteps       int                  `json:"purchaseSteps,omitempty"`
	MaxPurchase         int                  `json:"maxPurchase,omitempty"`
	ReleaseDate         string               `json:"releaseDate,omitempty"`
	Active              bool                 `json:"active,omitempty"`
	ShippingFree        bool                 `json:"shippingFree,omitempty"`
	PurchasePrice       float64              `json:"purchasePrice,omitempty"`
}
