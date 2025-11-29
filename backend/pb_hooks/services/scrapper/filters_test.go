package scrapper

import (
	"net/url"
	"strings"
	"testing"

	"pocketbase/pb_hooks/services"

	"github.com/PuerkitoBio/goquery"
)

func TestCleanupAnnoyance(t *testing.T) {
	// Setup mock FiltersList
	services.AllFilters = map[string][]string{
		"example.com": []string{".ad,.popup"},
	}

	// Create a sample HTML document
	html := `
		<html>
			<body>
				<div class="ad">Ad content</div>
				<div class="popup">Popup content</div>
				<div class="content">Main content</div>
			</body>
		</html>
	`
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		t.Fatalf("Failed to create document: %v", err)
	}

	testUrl, _ := url.Parse("https://example.com/page")

	// Call CleanupAnnoyance
	doc, err = CleanupFilters(testUrl, doc)
	if err != nil {
		t.Fatalf("CleanupAnnoyance returned error: %v", err)
	}

	// Check that .ad and .popup elements are removed
	if doc.Find(".ad").Length() != 0 {
		t.Errorf(".ad elements were not removed")
	}
	if doc.Find(".popup").Length() != 0 {
		t.Errorf(".popup elements were not removed")
	}
	// Check that .content element remains
	if doc.Find(".content").Length() != 1 {
		t.Errorf(".content element was incorrectly removed")
	}
}

func TestCleanupAnnoyance_Prefix(t *testing.T) {
	// Setup mock FiltersList
	services.AllFilters = map[string][]string{
		"example.com": []string{".ad,.popup"},
	}

	// Create a sample HTML document
	html := `
		<html>
			<body>
				<div class="ad">Ad content</div>
				<div class="popup">Popup content</div>
				<div class="content">Main content</div>
			</body>
		</html>
	`
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		t.Fatalf("Failed to create document: %v", err)
	}

	testUrl, _ := url.Parse("https://www.example.com/page")

	// Call CleanupAnnoyance
	doc, err = CleanupFilters(testUrl, doc)
	if err != nil {
		t.Fatalf("CleanupAnnoyance returned error: %v", err)
	}

	// Check that .ad and .popup elements are removed
	if doc.Find(".ad").Length() != 0 {
		t.Errorf(".ad elements were not removed")
	}
	if doc.Find(".popup").Length() != 0 {
		t.Errorf(".popup elements were not removed")
	}
	// Check that .content element remains
	if doc.Find(".content").Length() != 1 {
		t.Errorf(".content element was incorrectly removed")
	}
}

func TestCleanupAnnoyance_Unknown(t *testing.T) {
	// Setup mock FiltersList
	services.AllFilters = map[string][]string{
		"example.com": []string{".ad,.popup"},
	}

	// Create a sample HTML document
	html := `
		<html>
			<body>
				<div class="ad">Ad content</div>
				<div class="popup">Popup content</div>
				<div class="content">Main content</div>
			</body>
		</html>
	`
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		t.Fatalf("Failed to create document: %v", err)
	}

	testUrl, _ := url.Parse("https://different.com/page")

	// Call CleanupAnnoyance
	doc, err = CleanupFilters(testUrl, doc)
	if err != nil {
		t.Fatalf("CleanupAnnoyance returned error: %v", err)
	}

	// Check that .ad and .popup elements are removed
	if doc.Find(".ad").Length() == 0 {
		t.Errorf(".ad elements were removed")
	}
	if doc.Find(".popup").Length() == 0 {
		t.Errorf(".popup elements were removed")
	}
	// Check that .content element remains
	if doc.Find(".content").Length() != 1 {
		t.Errorf(".content element was incorrectly removed")
	}
}
