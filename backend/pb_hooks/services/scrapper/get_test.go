package scrapper

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestFetchHTML_Success(t *testing.T) {
	expected := "<html><head></head><body>Hello, world!</body></html>"
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(expected))
	}))
	defer server.Close()

	document, err := FetchRawHTML(server.URL)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	html, _ := document.Html()
	if html != expected {
		t.Errorf("expected %q, got %q", expected, html)
	}
}

func TestFetchRawHTML_Non200Status(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
	}))
	defer server.Close()

	_, err := FetchRawHTML(server.URL)
	if err == nil {
		t.Fatal("expected error for non-200 status, got nil")
	}
}

func TestFetchHTML_InvalidURL(t *testing.T) {
	_, err := FetchRawHTML("http://[::1]:namedport")
	if err == nil {
		t.Fatal("expected error for invalid URL, got nil")
	}
}
