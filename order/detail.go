package order

type Detail struct {
	Id            int                   `json:",omitempty"`
	OrderId       int                   `json:",omitempty"`
	ArticleId     int                   `json:",omitempty"`
	TaxId         int                   `json:",omitempty"`
	TaxRate       float64               `json:",omitempty"`
	StatusId      int                   `json:",omitempty"`
	Number        string                `json:",omitempty"`
	ArticleNumber string                `json:",omitempty"`
	Price         float64               `json:",omitempty"`
	Quantity      int                   `json:",omitempty"`
	ArticleName   string                `json:",omitempty"`
	Shipped       int                   `json:",omitempty"`
	ShippedGroup  int                   `json:",omitempty"`
	ReleaseDate   string                `json:",omitempty"`
	Mode          int                   `json:",omitempty"`
	EsdArticle    int                   `json:",omitempty"`
	Config        string                `json:",omitempty"`
	Ean           string                `json:",omitempty"`
	Unit          string                `json:",omitempty"`
	PackUnit      string                `json:",omitempty"`
	Attribute     *OrderDetailAttribute `json:",omitempty"`
}
