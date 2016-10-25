package propertygroup

type Option struct {
	Id         int    `json:",omitempty"`
	Name       string `json:",omitempty"`
	Filterable bool   `json:",omitempty"`
}
