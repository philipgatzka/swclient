package article

type Data struct {
	Id                int                 `json:",omitempty"`
	MainDetailId      int                 `json:",omitempty"`
	SupplierId        int                 `json:",omitempty"`
	TaxId             int                 `json:",omitempty"`
	PriceGroupId      int                 `json:",omitempty"`
	FilterGroupId     int                 `json:",omitempty"`
	ConfiguratorSetId int                 `json:",omitempty"`
	Name              string              `json:",omitempty"`
	Description       string              `json:",omitempty"`
	DescriptionLong   string              `json:",omitempty"`
	Added             string              `json:",omitempty"`
	Active            bool                `json:",omitempty"`
	PseudoSales       int                 `json:",omitempty"`
	Highlight         bool                `json:",omitempty"`
	Keywords          string              `json:",omitempty"`
	MetaTitle         string              `json:",omitempty"`
	Changed           string              `json:",omitempty"`
	PriceGroupActive  bool                `json:",omitempty"`
	LastStock         bool                `json:",omitempty"`
	CrossBundleLook   int                 `json:",omitempty"`
	Notification      bool                `json:",omitempty"`
	Template          string              `json:",omitempty"`
	Mode              int                 `json:",omitempty"`
	AvailableFrom     string              `json:",omitempty"`
	AvailableTo       string              `json:",omitempty"`
	MainDetail        *Detail             `json:",omitempty"`
	Tax               *Tax                `json:",omitempty"`
	PropertyValue     *PropertyValue      `json:",omitempty"`
	Supplier          *Supplier           `json:",omitempty"`
	PropertyGroup     *PropertyGroup      `json:",omitempty"`
	CustomerGroups    []CustomerGroup     `json:",omitempty"`
	Images            []Image             `json:",omitempty"`
	ConfiguratorSet   *ConfiguratorSet    `json:",omitempty"`
	Links             []Link              `json:",omitempty"`
	Downloads         []Download          `json:",omitempty"`
	Categories        map[int]Category    `json:",omitempty"`
	Similar           *[]Similar          `json:",omitempty"`
	Related           *[]Related          `json:",omitempty"`
	Details           *[]Detail           `json:",omitempty"`
	Translations      map[int]Translation `json:",omitempty"`
}
