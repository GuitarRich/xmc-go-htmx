package editing

const (
	EditModeChromes  string = "chromes"
	EditModeMetadata        = "metadata"
)

const (
	QUERY_PARAM_EDITING_SECRET           string = "secret"
	QUERY_PARAM_VERCEL_PROTECTION_BYPASS        = "x-vercel-protection-bypass"
	QUERY_PARAM_VERCEL_SET_BYPASS_COOKIE        = "x-vercel-set-bypass-cookie"
)

const DEFAULT_VARIANT string = "_default"
const VARIANT_PREFIX string = "_variantId_"

/**
 * Headers that should be passed along to (Editing Chromes handler) SSR request.
 * Note these are in lowercase format to match expected `IncomingHttpHeaders`.
 */
var EDITING_PASS_THROUGH_HEADERS = [...]string{"authorization", "cookie"}

/**
 * Default allowed origins for editing requests. This is used to enforce CORS, CSP headers.
 */
var EDITING_ALLOWED_ORIGINS = []string{"https://pages.sitecorecloud.io"}

type ResolvePageUrlArgs struct {
	ServerUrl string `json:"serverUrl"`
	ItemPath  string `json:"itemPath"`
}

type ResolvePageUrl func(ResolvePageUrlArgs) string

type EditingRenderMiddlewareConfig struct {
	ResolvePageUrl ResolvePageUrl
}

type MetadataQueryParams struct {
	Secret     string `json:"secret"`
	Sc_Lang    string `json:"sc_lang"`
	Sc_ItemId  string `json:"sc_itemid"`
	Sc_Site    string `json:"sc_site"`
	Route      string `json:"route"`
	Mode       string `json:"mode"`
	Sc_Variant string `json:"sc_variant"`
	Sc_Version string `json:"sc_version"`
}

type MetadataNextApiRequest struct {
	Query MetadataQueryParams
}

type Metadata struct {
	Packages map[string]interface{} `json:"packages"`
}

type EditingMiddlewareConfig struct {
	Components    []string               `json:"components"`
	PagesEditMode string                 `json:"editMode"`
	Packages      map[string]interface{} `json:"packages"`
}

type EditingMetadataPreviewData struct {
	Site       string   `json:"site"`
	ItemId     string   `json:"itemId"`
	Language   string   `json:"language"`
	EditMode   string   `json:"editMode"`
	PageState  string   `json:"pageState"`
	VariantIds []string `json:"variantIds"`
	Version    string   `json:"version"`
	LayoutKind string   `json:"layoutKind"`
}

type EditingRequestHandler struct {
	EditingMetadataPreviewData
}
