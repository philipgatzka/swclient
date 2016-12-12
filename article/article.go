package article

type Article struct {
	Id                int                 `json:"id,omitempty"`
	MainDetailId      int                 `json:"mainDetailId,omitempty"`
	SupplierId        int                 `json:"supplierId,omitempty"`
	TaxId             int                 `json:"taxId,omitempty"`
	PriceGroupId      int                 `json:"priceGroupId,omitempty"`
	FilterGroupId     int                 `json:"filterGroupId,omitempty"`
	ConfiguratorSetId int                 `json:"configuratorSetId,omitempty"`
	Name              string              `json:"name,omitempty"`
	Description       string              `json:"description,omitempty"`
	DescriptionLong   string              `json:"descriptionLong,omitempty"`
	Added             string              `json:"added,omitempty"`
	Active            bool                `json:"active,omitempty"`
	PseudoSales       int                 `json:"pseudoSales,omitempty"`
	Highlight         bool                `json:"highlight,omitempty"`
	Keywords          string              `json:"keywords,omitempty"`
	MetaTitle         string              `json:"metaTitle,omitempty"`
	Changed           string              `json:"changed,omitempty"`
	PriceGroupActive  bool                `json:"priceGroupActive,omitempty"`
	LastStock         bool                `json:"lastStock,omitempty"`
	CrossBundleLook   int                 `json:"crossBundleLook,omitempty"`
	Notification      bool                `json:"notification,omitempty"`
	Template          string              `json:"template,omitempty"`
	Mode              int                 `json:"mode,omitempty"`
	AvailableFrom     string              `json:"availableFrom,omitempty"`
	AvailableTo       string              `json:"availableTo,omitempty"`
	MainDetail        *Detail             `json:"mainDetail,omitempty"`
	Tax               *Tax                `json:"tax,omitempty"`
	PropertyValue     *PropertyValue      `json:"propertyValue,omitempty"`
	Supplier          *Supplier           `json:"supplier,omitempty"`
	PropertyGroup     *PropertyGroup      `json:"propertyGroup,omitempty"`
	CustomerGroups    []CustomerGroup     `json:"customerGroups,omitempty"`
	Images            []Image             `json:"images,omitempty"`
	ConfiguratorSet   *ConfiguratorSet    `json:"configuratorSet,omitempty"`
	Links             []Link              `json:"links,omitempty"`
	Downloads         []Download          `json:"downloads,omitempty"`
	Categories        []Category          `json:"categories,omitempty"`
	Similar           *[]Similar          `json:"similar,omitempty"`
	Related           *[]Related          `json:"related,omitempty"`
	Details           *[]Detail           `json:"details,omitempty"`
	Translations      map[int]Translation `json:"translations,omitempty"`
	MainNumber        string              `json:"mainNumber,omitempty"`
}
