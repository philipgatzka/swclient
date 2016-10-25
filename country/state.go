package country

type State struct {
	Id        int    `json:",omitempty"`
	CountryId int    `json:",omitempty"`
	Position  int    `json:",omitempty"`
	Name      string `json:",omitempty"`
	ShortCode string `json:",omitempty"`
	Active    bool   `json:",omitempty"`
}
