package article

type Supplier struct {
	Id              int    `json:"id,omitempty"`
	Name            string `json:"name,omitempty"`
	Image           string `json:"image,omitempty"`
	Link            string `json:"link,omitempty"`
	Description     string `json:"description,omitempty"`
	MetaTitle       string `json:"metaTitle,omitempty"`
	MetaDescription string `json:"metaDescription,omitempty"`
	MetaKeywords    string `json:"metaKeywords,omitempty"`
}
