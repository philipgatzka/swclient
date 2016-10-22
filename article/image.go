package article

type Image struct {
	Id              int
	ArticleId       int
	ArticleDetailId int
	Description     string
	Path            string
	Main            int
	Position        int
	Width           int
	Height          int
	Relations       string
	Extension       string
	ParentId        int
	MediaId         int
}
