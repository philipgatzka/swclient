package article

import (
	"encoding/json"
)

type ConfiguratorOption struct {
	Id       int    `json:"id,omitempty"`
	GroupId  int    `json:"groupId,omitempty"`
	Name     string `json:"name,omitempty"`
	Position int    `json:"position,omitempty"`
	Group    string `json:"group,omitempty"`
}

// MarshalJSON translates an article into JSON.
// This is necessary, because the Shopware API returns a slightly different object on GETting than it expects on POSTing.
func (o ConfiguratorOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Id       int    `json:"id,omitempty"`
		GroupId  int    `json:"groupId,omitempty"`
		Option   string `json:"option,omitempty"`
		Name     string `json:"name,omitempty"`
		Position int    `json:"position,omitempty"`
		Group    string `json:"group,omitempty"`
	}{
		Id:       o.Id,
		GroupId:  o.GroupId,
		Option:   o.Name,
		Name:     o.Name,
		Position: o.Position,
		Group:    o.Group,
	})
}
