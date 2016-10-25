package order

type Dispatch struct {
	Id                   int     `json:",omitempty"`
	Name                 string  `json:",omitempty"`
	Typ                  int     `json:",omitempty"`
	Description          string  `json:",omitempty"`
	Comment              string  `json:",omitempty"`
	Active               bool    `json:",omitempty"`
	Position             int     `json:",omitempty"`
	Calculation          int     `json:",omitempty"`
	SurchargeCalculation int     `json:",omitempty"`
	TaxCalculation       int     `json:",omitempty"`
	ShippingFree         float64 `json:",omitempty"`
	MultiShopId          int     `json:",omitempty"`
	CustomerGroupId      int     `json:",omitempty"`
	BindShippingFree     int     `json:",omitempty"`
	BindTimeFrom         int     `json:",omitempty"`
	BindTimeTo           int     `json:",omitempty"`
	BindInStock          int     `json:",omitempty"`
	BindWeekdayFrom      int     `json:",omitempty"`
	BindPriceTo          int     `json:",omitempty"`
	BindSql              string  `json:",omitempty"`
	StatusLink           string  `json:",omitempty"`
	CalculationSql       string  `json:",omitempty"`
}
