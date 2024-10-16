package model

type EditMode string

const (
	Chromes  EditMode = "chromes"
	Metadata          = "metadata"
)

func (s EditMode) String() string {
	return string(s)
}

type PageSate string

const (
	Preview PageSate = "preview"
	Edit    PageSate = "edit"
	Normal  PageSate = "normal"
)

type Datasource struct {
	Id       string `json:"id"`
	Language string `json:"language"`
	Revision string `json:"revision"`
	Version  string `json:"version"`
}

type MetadataData struct {
	Datasource Datasource `json:"datasource"`
	Title      string     `json:"title"`
	FieldId    string     `json:"fieldId"`
	FieldType  string     `json:"fieldType"`
	RawValue   string     `json:"rawValue"`
}

func (s PageSate) String() string {
	return string(s)
}

type Placeholder struct {
	Components []PlaceholderComponent `json:"components"`
}

type RichTextField struct {
	Value    interface{}  `json:"value"`
	Editable string       `json:"editable"`
	Metadata MetadataData `json:"metadata"`
}

type ImageField struct {
	Src      string       `json:"src"`
	Alt      string       `json:"alt"`
	Width    string       `json:"width"`
	Height   string       `json:"height"`
	Metadata MetadataData `json:"metadata"`
}

type LinkField struct {
	Href        string       `json:"href"`
	Anchor      string       `json:"anchor"`
	Querystring string       `json:"querystring"`
	Text        string       `json:"text"`
	Target      string       `json:"target"`
	Title       string       `json:"title"`
	Class       string       `json:"class"`
	Metadata    MetadataData `json:"metadata"`
}

type ScField struct {
	Value    interface{}  `json:"value"`
	Metadata MetadataData `json:"metadata"`
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
	Fields       interface{}                       `json:"fields"`
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

type Rendered struct {
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
}

type LayoutResponse struct {
	Data struct {
		Layout struct {
			Item struct {
				Rendered Rendered `json:"rendered"`
			} `json:"item"`
		} `json:"layout"`
	} `json:"data"`
}

type NotFoundPageResponse struct {
	Data struct {
		Site struct {
			SiteInfo struct {
				ErrorHandling struct {
					NotFoundPage struct {
						Rendered Rendered `json:"rendered"`
					} `json:"notFoundPage"`
				} `json:"errorHandling"`
			} `json:"siteInfo"`
		} `json:"site"`
	} `json:"data"`
}
