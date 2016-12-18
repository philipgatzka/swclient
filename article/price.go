package article

type Price struct {
	Customergroupkey string         `json:"customerGroupKey,omitempty"`
	Customergroup    *CustomerGroup `json:"customerGroup,omitempty"`
	Articledetailsid int            `json:"articleDetailsId,omitempty"`
	Articleid        int            `json:"articleId,omitempty"`
	Id               int            `json:"id,omitempty"`
	From             int            `json:"from,omitempty"`
	To               string         `json:"to,omitempty"`
	Price            float64        `json:"price"`
	Pseudoprice      float64        `json:"pseudoPrice"`
	Baseprice        float64        `json:"basePrice"`
	Percent          float64        `json:"percent,omitempty"`
}
