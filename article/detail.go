package article

type Detail struct {
	Number              string
	SupplierNumber      string
	AdditionalText      string
	Weight              string
	Width               string
	Len                 string
	Height              string
	Ean                 string
	PurchaseUnit        string
	DescriptionLong     string
	ReferenceUnit       string
	PackUnit            string
	ShippingTime        string
	Prices              []Price
	ConfiguratorOptions []ConfiguratorOption
	Attribute           Attribute
	Id                  int
	ArticleId           int
	UnitId              int
	Kind                int
	InStock             int
	Position            int
	MinPurchase         int
	PurchaseSteps       int
	MaxPurchase         int
	ReleaseDate         string
	Active              bool
	ShippingFree        bool
}
