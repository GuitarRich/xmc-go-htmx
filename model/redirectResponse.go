package model

type RedirectResponse struct {
	Data struct {
		Site struct {
			SiteInfo struct {
				Name      string     `json:"name"`
				RootPath  string     `json:"rootPath"`
				Redirects []Redirect `json:"redirects"`
			} `json:"siteInfo"`
		} `json:"site"`
	} `json:"data"`
}

type Redirect struct {
	RedirectType           string `json:"redirectType"`
	IsQueryStringPreserved bool   `json:"isQueryStringPreserved"`
	Target                 string `json:"target"`
	Pattern                string `json:"pattern"`
}
