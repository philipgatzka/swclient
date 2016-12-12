package article

type Link struct {
	Id        int    `json:"id,omitempty"`
	ArticleId int    `json:"articleId,omitempty"`
	Name      string `json:"name,omitempty"`
	Link      string `json:"link,omitempty"`
	Target    string `json:"target,omitempty"`
}
