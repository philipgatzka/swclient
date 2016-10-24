package article

type Image struct {
	Id              int    `json:",omitempty"`
	ArticleId       int    `json:",omitempty"`
	ArticleDetailId int    `json:",omitempty"`
	Description     string `json:",omitempty"`
	Path            string `json:",omitempty"`
	Main            int    `json:",omitempty"`
	Position        int    `json:",omitempty"`
	Width           int    `json:",omitempty"`
	Height          int    `json:",omitempty"`
	Relations       string `json:",omitempty"`
	Extension       string `json:",omitempty"`
	ParentId        int    `json:",omitempty"`
	MediaId         int    `json:",omitempty"`
}
