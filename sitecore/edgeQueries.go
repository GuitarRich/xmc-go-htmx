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
