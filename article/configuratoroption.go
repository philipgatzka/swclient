package article

type ConfiguratorOption struct {
	Id       int    `json:"id,omitempty"`
	GroupId  int    `json:"groupId,omitempty"`
	Name     string `json:"name,omitempty"`
	Position int    `json:"position,omitempty"`
	Group    string `json:"group,omitempty"`
}
