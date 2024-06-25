package sitecore

import (
	"fmt"
)

func GetLayoutQuery(itemPath string, language string, siteName string) string {
	return fmt.Sprintf(`
        {
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
        {
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
        {
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
        {
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
