package article

type Price struct {
	Customergroupkey string
	Customergroup    CustomerGroup
	Articledetailsid int
	Articleid        int
	Id               int
	From             int
	To               string
	Price            float64
	Pseudoprice      float64
	Baseprice        float64
	Percent          float64
}
