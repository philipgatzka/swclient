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
	Similar           *[]Similar          `json:"similar,omitempty"`
	Related           *[]Related          `json:"related,omitempty"`
	Details           *[]Detail           `json:"details,omitempty"`
	Translations      map[int]Translation `json:"translations,omitempty"`
}

// MarshalJSON translates an article into JSON.
// This is necessary, because the Shopware API returns a slightly different object on GETting than it expects on POSTing.
func (a Article) MarshalJSON() ([]byte, error) {
	temp := map[string]interface{}{}
	temp["id"] = a.Id
	temp["mainDetailId"] = a.MainDetailId
	temp["supplierId"] = a.SupplierId
	temp["taxId"] = a.TaxId
	temp["priceGroupId"] = a.PriceGroupId
	temp["filterGroupId"] = a.FilterGroupId
	temp["configuratorSetId"] = a.ConfiguratorSetId
	temp["name"] = a.Name
	temp["description"] = a.Description
	temp["descriptionLong"] = a.DescriptionLong
	temp["added"] = a.Added
	temp["active"] = a.Active
	temp["pseudoSales"] = a.PseudoSales
	temp["highlight"] = a.Highlight
	temp["keywords"] = a.Keywords
	temp["metaTitle"] = a.MetaTitle
	temp["changed"] = a.Changed
	temp["priceGroupActive"] = a.PriceGroupActive
	temp["lastStock"] = a.LastStock
	temp["crossBundleLook"] = a.CrossBundleLook
	temp["notification"] = a.Notification
	temp["template"] = a.Template
	temp["mode"] = a.Mode
	temp["availableFrom"] = a.AvailableFrom
	temp["availableTo"] = a.AvailableTo
	temp["mainDetail"] = a.MainDetail
	temp["tax"] = a.Tax
	temp["propertyValue"] = a.PropertyValue
	temp["supplier"] = a.Supplier.Name
	temp["propertyGroup"] = a.PropertyGroup
	temp["customerGroups"] = a.CustomerGroups
	temp["images"] = a.Images
	temp["configuratorSet"] = a.ConfiguratorSet
	temp["links"] = a.Links
	temp["downloads"] = a.Downloads
	temp["categories"] = a.Categories
	temp["similar"] = a.Similar
	temp["related"] = a.Related
	temp["variants"] = a.Details
	temp["translations"] = a.Translations
	return json.Marshal(temp)
}
