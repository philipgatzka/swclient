package translation

type Locale struct {
	Id        int    `json:",omitempty"`
	Locale    string `json:",omitempty"`
	Language  string `json:",omitempty"`
	Territory string `json:",omitempty"`
}
