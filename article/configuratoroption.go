package article

import "encoding/json"

type ConfiguratorOption struct {
	Id       int    `json:"id,omitempty"`
	GroupId  int    `json:"groupId,omitempty"`
	Name     string `json:"name,omitempty"`
	Position int    `json:"position,omitempty"`
	Group    string `json:"group,omitempty"`
}

// MarshalJSON translates an article into JSON.
// This is necessary, because the Shopware API returns a slightly different object on GETting than it expects on POSTing.
// UPDATE: Had to leave "Name" AND "Option" in group due to a bug, got resolved: https://github.com/shopware/shopware/commit/76a58fc1a43600d8d6be5b3248a044fba84b8526
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
