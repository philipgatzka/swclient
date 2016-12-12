package article

type ConfiguratorSet struct {
	Id     int                 `json:"id,omitempty"`
	Name   string              `json:"name,omitempty"`
	Public bool                `json:"public,omitempty"`
	Typ    int                 `json:"type,omitempty"`
	Groups []ConfiguratorGroup `json:"groups,omitempty"`
}
