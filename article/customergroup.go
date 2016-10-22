package article

type CustomerGroup struct {
	Id                    int
	Key                   string
	Name                  string
	Tax                   bool
	TaxInput              bool
	Mode                  bool
	Discount              float64
	MinimumOrder          float64
	MinimumOrderSurcharge float64
	BasePrice             float64
	Percent               float64
}
