package variant

type Price struct {
	Customergroupkey string         `json:",omitempty"`
	Customergroup    *CustomerGroup `json:",omitempty"`
	Articledetailsid int            `json:",omitempty"`
	Articleid        int            `json:",omitempty"`
	Id               int            `json:",omitempty"`
	From             int            `json:",omitempty"`
	To               string         `json:",omitempty"`
	Price            float64        `json:",omitempty"`
	Pseudoprice      float64        `json:",omitempty"`
	Baseprice        float64        `json:",omitempty"`
	Percent          float64        `json:",omitempty"`
}
