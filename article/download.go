package article

type Download struct {
	Id        int    `json:",omitempty"`
	ArticleId int    `json:",omitempty"`
	Name      string `json:",omitempty"`
	File      string `json:",omitempty"`
	Size      int    `json:",omitempty"`
}
