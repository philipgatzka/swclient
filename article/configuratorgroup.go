package article

type ConfiguratorGroup struct {
	Id          int                  `json:"id,omitempty"`
	Description string               `json:"description,omitempty"`
	Name        string               `json:"name,omitempty"`
	Position    int                  `json:"position,omitempty"`
	Options     []ConfiguratorOption `json:"options,omitempty"`
}
