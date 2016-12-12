package article

type PropertyGroup struct {
	Id         int    `json:"id,omitempty"`
	Name       string `json:"name,omitempty"`
	Position   int    `json:"position,omitempty"`
	Comparable bool   `json:"comparable,omitempty"`
	SortMode   int    `json:"sortMode,omitempty"`
}
