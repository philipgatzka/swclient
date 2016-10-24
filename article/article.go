package article

type Article struct {
	Data    Data `json:",omitempty"`
	Success bool `json:",omitempty"`
}
