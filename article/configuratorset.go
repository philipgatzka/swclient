package article

type ConfiguratorSet struct {
	Id     int
	Name   string
	Public bool
	Typ    int `json:"type"`
	Groups []ConfiguratorGroup
}
