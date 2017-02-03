package manufacturer

type Manufacturer struct {
	ID          int    `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	Image       string `json:"image,omitempty"`
	Link        string `json:"link,omitempty"`
	Description string `json:"description,omitempty"`
}

type Manufacturers []Manufacturer

// Len implements the sort.Interface
func (m Manufacturers) Len() int {
	return len(m)
}

// Less implements the sort.Interface
func (m Manufacturers) Less(i, j int) bool {
	return m[i].Name < m[j].Name
}

// Swap implements the sort.Interface
func (m Manufacturers) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

// String implements the Stringer interface. Returns the articles Number and Name
func (m Manufacturer) String() string {
	return m.Name
}
