package scrapper

import (
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"

	"pocketbase/pb_hooks/services"
)

func CleanupFilters(url *url.URL, document *goquery.Document) (*goquery.Document, error) {

	// Use services.AnnoyanceList, of type map[string]string
	// to find comma-separated rules using list[url.Hostname()]
	// remove all elements matching these rules
	hostname := strings.TrimPrefix(url.Hostname(), "www.")
	if selectors, ok := services.AllFilters[hostname]; ok && selectors != nil {
		for _, selector := range selectors {
			if selector != "" {
				document.Find(selector).Remove()
			}
		}
	}

	return document, nil
}
