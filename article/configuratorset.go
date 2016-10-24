package article

type ConfiguratorSet struct {
	Id     int                 `json:",omitempty"`
	Name   string              `json:",omitempty"`
	Public bool                `json:",omitempty"`
	Typ    int                 `json:"type,omitempty"`
	Groups []ConfiguratorGroup `json:",omitempty"`
}
