package supplier

type Supplier struct {
	ID          int    `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	Image       string `json:"image,omitempty"`
	Link        string `json:"link,omitempty"`
	Description string `json:"description,omitempty"`
}

type Suppliers []Supplier

// Len implements the sort.Interface
func (s Suppliers) Len() int {
	return len(s)
}

// Less implements the sort.Interface
func (s Suppliers) Less(i, j int) bool {
	return s[i].Name < s[j].Name
}

// Swap implements the sort.Interface
func (s Suppliers) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

// String implements the Stringer interface. Returns the articles Number and Name
func (s Supplier) String() string {
	return s.Name
}
