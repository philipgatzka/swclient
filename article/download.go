package article

type Download struct {
	Id        int    `json:"id,omitempty"`
	ArticleId int    `json:"articleId,omitempty"`
	Name      string `json:"name,omitempty"`
	File      string `json:"file,omitempty"`
	Size      int    `json:"size,omitempty"`
}
