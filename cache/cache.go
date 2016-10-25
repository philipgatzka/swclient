package cache

type Cache struct {
	Dir       string `json:",omitempty"`
	Size      string `json:",omitempty"`
	Files     int    `json:",omitempty"`
	FreeSpace string `json:",omitempty"`
	Name      string `json:",omitempty"`
	Backend   string `json:",omitempty"`
	Id        string `json:",omitempty"`
}
