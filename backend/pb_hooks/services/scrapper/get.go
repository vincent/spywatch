package scrapper

import (
	"errors"
	"net/http"
	"net/url"
	"time"

	"github.com/PuerkitoBio/goquery"
)

func FetchHTML(target string) (string, error) {
	url, err := url.Parse(target)
	if err != nil {
		return "", err
	}

	document, err := FetchRawHTML(target)
	if err != nil {
		return "", err
	}

	document, _ = CleanupFilters(url, document)

	return document.Html()
}

func FetchRawHTML(url string) (*goquery.Document, error) {
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (compatible; FetchHTML/1.0)")

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("failed to fetch HTML: status " + resp.Status)
	}

	return goquery.NewDocumentFromReader(resp.Body)
}
