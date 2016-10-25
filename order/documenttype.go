package order

type DocumentType struct {
	Id        int    `json:",omitempty"`
	Name      string `json:",omitempty"`
	Template  string `json:",omitempty"`
	Numbers   string `json:",omitempty"`
	Left      int    `json:",omitempty"`
	Right     int    `json:",omitempty"`
	Top       int    `json:",omitempty"`
	Bottom    int    `json:",omitempty"`
	PageBreak int    `json:",omitempty"`
}
