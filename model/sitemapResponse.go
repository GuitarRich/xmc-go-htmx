package model

type SitemapResponse struct {
	Data struct {
		Site struct {
			SiteInfo struct {
				Sitemap []string `json:"sitemap"`
			} `json:"siteInfo"`
		} `json:"site"`
	} `json:"data"`
}
