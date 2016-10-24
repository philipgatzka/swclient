package article

type Link struct {
	Id        int    `json:",omitempty"`
	ArticleId int    `json:",omitempty"`
	Name      string `json:",omitempty"`
	Link      string `json:",omitempty"`
	Target    string `json:",omitempty"`
}
