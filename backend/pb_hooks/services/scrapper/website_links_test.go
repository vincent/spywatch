package scrapper

import (
	"net/url"
	"strings"
	"testing"
)

// Unit tests for categorizeLinks
func TestCategorizeLinks(t *testing.T) {
	baseURL := "https://example.com"
	links := []string{
		"https://example.com/about",
		"https://example.com/contact",
		"https://facebook.com/example",
		"https://instagram.com/example",
		"https://example.com/favicon.ico",
		"https://example.com/#section",
		"https://example.com/rss",
		"https://example.com/file.atom",
		"mailto:info@example.com",
		"https://twitter.com/example",
		"https://example.com/about", // duplicate
	}

	ownLinks, socialLinks, logo := categorizeLinks(links, baseURL)

	// Check ownLinks
	expectedOwn := map[string]bool{
		"https://example.com/about":   true,
		"https://example.com/contact": true,
	}
	if len(ownLinks) != len(expectedOwn) {
		t.Errorf("expected %d ownLinks, got %d", len(expectedOwn), len(ownLinks))
	}
	for _, link := range ownLinks {
		if !expectedOwn[link] {
			t.Errorf("unexpected ownLink: %s", link)
		}
	}

	// Check socialLinks
	expectedSocial := map[string]bool{
		"https://facebook.com/example":  true,
		"https://instagram.com/example": true,
		"https://twitter.com/example":   true,
	}
	if len(socialLinks) != len(expectedSocial) {
		t.Errorf("expected %d socialLinks, got %d", len(expectedSocial), len(socialLinks))
	}
	for _, link := range socialLinks {
		if !expectedSocial[link] {
			t.Errorf("unexpected socialLink: %s", link)
		}
	}

	// Check logo
	expectedLogo := "https://example.com/favicon.ico"
	if logo != expectedLogo {
		t.Errorf("expected logo %s, got %s", expectedLogo, logo)
	}
}

// Unit tests for fetchAndParseMetadata
func TestFetchAndParseMetadata_Invalid(t *testing.T) {
	// This test will only check that the function returns an error for an invalid URL.
	badURL, _ := url.Parse("http://invalid.localhost")
	title, logo, err := fetchAndParseMetadata(badURL)
	if err == nil {
		t.Errorf("expected error for invalid URL, got nil")
	}
	if title != "" {
		t.Errorf("expected empty title for invalid URL, got %s", title)
	}
	if logo != "" {
		t.Errorf("expected empty logo for invalid URL, got %s", logo)
	}
}

func TestFetchAndParseMetadata_Parse(t *testing.T) {
	// This test will only check that the function returns an error for an invalid URL.
	url, _ := url.Parse("https://google.fr")
	title, logo, err := fetchAndParseMetadata(url)
	if err != nil {
		t.Errorf("expected nil error, got %s", err)
	}
	if title != "Google" {
		t.Errorf("expected title, got %s", title)
	}
	if !strings.HasSuffix(logo, "favicon.ico") {
		t.Errorf("expected logo, got %s", logo)
	}

	if !strings.HasPrefix(logo, "http://") {
		t.Errorf("expected absolute logo url, got %s", logo)
	}
}
