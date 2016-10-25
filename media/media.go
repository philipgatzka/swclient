package media

type Media struct {
	Id          int    `json:",omitempty"`
	AlbumId     int    `json:",omitempty"`
	Name        string `json:",omitempty"`
	Description string `json:",omitempty"`
	Path        string `json:",omitempty"`
	Typ         string `json:"type,omitempty"`
	Extension   string `json:",omitempty"`
	UserId      int    `json:",omitempty"`
	Created     string `json:",omitempty"`
	FileSize    int    `json:",omitempty"`
}
