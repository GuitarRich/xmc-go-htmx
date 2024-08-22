package model

type SiteContext struct {
	params        map[string]string
	preview       bool
	previewData   map[string]interface{}
	draftMode     bool
	locale        string
	locales       []string
	defaultLocale string
}
