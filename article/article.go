package article

type Article struct {
	Data struct {
		Id                int
		MainDetailId      int
		SupplierId        int
		TaxId             int
		PriceGroupId      int
		FilterGroupId     int
		ConfiguratorSetId int
		Name              string
		Description       string
		DescriptionLong   string
		Added             string
		Active            bool
		PseudoSales       int
		Highlight         bool
		Keywords          string
		MetaTitle         string
		Changed           string
		PriceGroupActive  bool
		LastStock         bool
		CrossBundleLook   int
		Notification      bool
		Template          string
		Mode              int
		AvailableFrom     string
		AvailableTo       string
		MainDetail        Detail
		Tax               Tax
		PropertyValue     PropertyValue
		Supplier          Supplier
		PropertyGroup     PropertyGroup
		CustomerGroups    []CustomerGroup
		Images            []Image
		ConfiguratorSet   ConfiguratorSet
		Links             []Link
		Downloads         []Download
		Categories        map[int]Category
		Similar           []Similar
		Related           []Related
		Details           []Detail
		Translations      map[int]Translation
	}
	Success bool
}
