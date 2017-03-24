package article

import "encoding/json"

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
	Similar           []Similar           `json:"similar,omitempty"`
	Related           []Related           `json:"related,omitempty"`
	Details           []Detail            `json:"details,omitempty"`
	Translations      map[int]Translation `json:"translations,omitempty"`
	MainNumber        string              `json:"mainNumber,omitempty"`
}

// New creates a new article with all data required by the shopware API
func New(name, supplier, number, customerGroupKey string, taxID, category int, price float64, active bool) (Article, error) {
	return Article{
		Active: active,
		TaxId:  taxID,
		Name:   name,
		Supplier: &Supplier{
			Name: supplier,
		},
		MainDetail: &Detail{
			Active: active,
			Number: number,
			Prices: []Price{
				{
					Price:            price,
					Customergroupkey: customerGroupKey,
				},
			},
		},
		Categories: []Category{
			{
				Id: category,
			},
		},
	}, nil
}

type Articles []Article

// Len implements the sort.Interface
func (a Articles) Len() int {
	return len(a)
}

// Less implements the sort.Interface
func (a Articles) Less(i, j int) bool {
	return a[i].MainDetail.Number < a[j].MainDetail.Number
}

// Swap implements the sort.Interface
func (a Articles) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

// String implements the Stringer interface. Returns the articles Number and Name
func (a Article) String() string {
	return a.MainDetail.Number + ": " + a.Name
}

// MarshalJSON translates an article into JSON.
// This is necessary, because the Shopware API returns a slightly different object on GETting than it expects on POSTing.
func (a Article) MarshalJSON() ([]byte, error) {
	tmp := struct {
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
		Supplier          string              `json:"supplier,omitempty"`
		PropertyGroup     *PropertyGroup      `json:"propertyGroup,omitempty"`
		CustomerGroups    []CustomerGroup     `json:"customerGroups,omitempty"`
		Images            []Image             `json:"images,omitempty"`
		ConfiguratorSet   *ConfiguratorSet    `json:"configuratorSet,omitempty"`
		Links             []Link              `json:"links,omitempty"`
		Downloads         []Download          `json:"downloads,omitempty"`
		Categories        []Category          `json:"categories,omitempty"`
		Similar           []Similar           `json:"similar,omitempty"`
		Related           []Related           `json:"related,omitempty"`
		Variants          []Detail            `json:"variants,omitempty"`
		Translations      map[int]Translation `json:"translations,omitempty"`
	}{
		Id:                a.Id,
		MainDetailId:      a.MainDetailId,
		SupplierId:        a.SupplierId,
		TaxId:             a.TaxId,
		PriceGroupId:      a.PriceGroupId,
		FilterGroupId:     a.FilterGroupId,
		ConfiguratorSetId: a.ConfiguratorSetId,
		Name:              a.Name,
		Description:       a.Description,
		DescriptionLong:   a.DescriptionLong,
		Added:             a.Added,
		Active:            a.Active,
		PseudoSales:       a.PseudoSales,
		Highlight:         a.Highlight,
		Keywords:          a.Keywords,
		MetaTitle:         a.MetaTitle,
		Changed:           a.Changed,
		PriceGroupActive:  a.PriceGroupActive,
		LastStock:         a.LastStock,
		CrossBundleLook:   a.CrossBundleLook,
		Notification:      a.Notification,
		Template:          a.Template,
		Mode:              a.Mode,
		AvailableFrom:     a.AvailableFrom,
		AvailableTo:       a.AvailableTo,
		MainDetail:        a.MainDetail,
		Tax:               a.Tax,
		PropertyValue:     a.PropertyValue,
		Supplier:          "",
		PropertyGroup:     a.PropertyGroup,
		CustomerGroups:    a.CustomerGroups,
		Images:            a.Images,
		ConfiguratorSet:   a.ConfiguratorSet,
		Links:             a.Links,
		Downloads:         a.Downloads,
		Categories:        a.Categories,
		Similar:           a.Similar,
		Related:           a.Related,
		Variants:          a.Details,
		Translations:      a.Translations,
	}

	if a.Supplier != nil {
		tmp.Supplier = a.Supplier.Name
	}

	return json.Marshal(tmp)
}
