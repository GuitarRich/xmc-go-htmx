package sitecore

import (
	"fmt"
)

func GetLayoutQuery(itemPath string, language string, siteName string) string {
	return fmt.Sprintf(`
        query LayoutQuery {
          layout(site: "%s", routePath: "%s", language: "%s") {
            item {
              rendered
            }
          }
        }
    `, siteName, itemPath, language)
}

func GetSitemapQuery(siteName string) string {
	return fmt.Sprintf(`
        query {
            site {
                siteInfo(site: "%s") {
                    sitemap
                }
            }
        }
    `, siteName)
}

func GetNotFoundPageQuery(siteName string, language string) string {
	return fmt.Sprintf(`
        query {
            site {
                siteInfo(site: "%s") {
                    errorHandling(language: "%s") {
                        notFoundPage {
                            rendered
                        }
                    }
                }
            }
        }
    `, siteName, language)
}

func GetRedirectsForSiteQuery(siteName string) string {
	return fmt.Sprintf(`
        query {
            site { 
                siteInfo(site: "%s") {
                    name
                    rootPath
                    redirects {
                        redirectType
                        isQueryStringPreserved
                        target
                        pattern
                    }
                }
            }
        }`, siteName)
}

func GetEditingDataQuery() string {
	return `
query EditingQuery($siteName: String!, $itemId: String!, $language: String!, $version: String, $after: String, $pageSize: Int = 10) {
	item(path: $itemId, language: $language, version: $version) {
	    rendered
	}
	site {
	    siteInfo(site: $siteName) {
	        dictionary(language: $language, first: $pageSize, after: $after) {
	            results {
					key
	                value
	            }
	            pageInfo {
	              endCursor
	              hasNext
	            }
	        }
	    }
	}
}`
}
