package model

type SitemapResponse struct {
	Data struct {
		Site struct {
			SiteInfo struct {
				Sitemap      []string `json:"sitemap"`
				NotFoundPage string   `json:"notFoundPage"`
			} `json:"siteInfo"`
		} `json:"site"`
	} `json:"data"`
}
