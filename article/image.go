package article

type Image struct {
	Id              int    `json:"id,omitempty"`
	ArticleId       int    `json:"articleId,omitempty"`
	ArticleDetailId int    `json:"articleDetailId,omitempty"`
	Description     string `json:"description,omitempty"`
	Path            string `json:"path,omitempty"`
	Main            int    `json:"main,omitempty"`
	Position        int    `json:"position,omitempty"`
	Width           int    `json:"width,omitempty"`
	Height          int    `json:"height,omitempty"`
	Relations       string `json:"relations,omitempty"`
	Extension       string `json:"extension,omitempty"`
	ParentId        int    `json:"parentId,omitempty"`
	MediaId         int    `json:"mediaId,omitempty"`
}
