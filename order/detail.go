package order

type Detail struct {
	id            int
	orderId       string
	articleId     int
	taxId         int
	taxRate       float64
	statusId      int
	number        string
	articleNumber string
	price         float64
	quantity      int
	articleName   string
	shipped       bool
	shippedGroup  int
	releaseDate   string
	mode          int
	esdArticle    int
	config        string
	ean           string
	unit          string
	packUnit      string
	attribute     *OrderDetailAttribute
}
