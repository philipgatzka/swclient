package article

type ConfiguratorGroup struct {
	Id          int    `json:",omitempty"`
	Description string `json:",omitempty"`
	Name        string `json:",omitempty"`
	Position    int    `json:",omitempty"`
}
