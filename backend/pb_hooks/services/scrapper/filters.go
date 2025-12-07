package scrapper

import (
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"

	"pocketbase/pb_hooks/services"
)

func CleanupFilters(url *url.URL, document *goquery.Document) (*goquery.Document, error) {

	// Use services.AllFilters, of type map[string]string
	// to find comma-separated rules using list[url.Hostname()]
	// remove all elements matching these rules
	applyFilters("*", document)

	hostname := strings.TrimPrefix(url.Hostname(), "www.")
	applyFilters(hostname, document)

	return document, nil
}

func applyFilters(key string, document *goquery.Document) {
	if selectors, ok := services.AllFilters[key]; ok && selectors != nil {
		for _, selector := range selectors {
			if selector != "" {
				document.Find(selector).Remove()
			}
		}
	}
}
