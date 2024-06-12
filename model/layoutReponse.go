package model

type Placeholder struct {
	Components []PlaceholderComponent `json:"components"`
}

type ScField struct {
	Value interface{} `json:"value"`
}

type PlaceholderComponent struct {
	UID           string `json:"uid"`
	ComponentName string `json:"componentName"`
	DataSource    string `json:"dataSource"`
	Params        struct {
		GridParameters       string `json:"GridParameters"`
		FieldNames           string `json:"FieldNames"`
		Styles               string `json:"Styles"`
		DynamicPlaceholderID string `json:"DynamicPlaceholderId"`
		Sig                  string `json:"sig"`
		Ph                   string `json:"ph"`
	} `json:"params"`
	Fields       map[string]ScField                `json:"fields"`
	Placeholders map[string][]PlaceholderComponent `json:"placeholders"`
}

type RouteData struct {
	Name         string                            `json:"name"`
	DisplayName  string                            `json:"displayName"`
	Fields       map[string]interface{}            `json:"fields"`
	DatabaseName string                            `json:"databaseName"`
	DeviceID     string                            `json:"deviceId"`
	ItemID       string                            `json:"itemId"`
	ItemLanguage string                            `json:"itemLanguage"`
	ItemVersion  int                               `json:"itemVersion"`
	LayoutID     string                            `json:"layoutId"`
	TemplateID   string                            `json:"templateId"`
	TemplateName string                            `json:"templateName"`
	Placeholders map[string][]PlaceholderComponent `json:"placeholders"`
}

type LayoutResponse struct {
	Data struct {
		Layout struct {
			Item struct {
				Rendered struct {
					Sitecore struct {
						Context struct {
							PageEditing bool `json:"pageEditing"`
							Site        struct {
								Name string `json:"name"`
							} `json:"site"`
							PageState string `json:"pageState"`
							EditMode  string `json:"editMode"`
							Language  string `json:"language"`
							ItemPath  string `json:"itemPath"`
						} `json:"context"`
						Route RouteData `json:"route"`
					} `json:"sitecore"`
				} `json:"rendered"`
			} `json:"item"`
		} `json:"layout"`
	} `json:"data"`
}
