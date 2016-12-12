package article

type Tax struct {
	Id   int    `json:"id,omitempty"`
	Tax  string `json:"tax,omitempty"`
	Name string `json:"name,omitempty"`
}
