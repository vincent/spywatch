package scrapper

import (
	"regexp"
	"strings"
	"unicode"

	htmltomarkdown "github.com/JohannesKaufmann/html-to-markdown/v2"
	"github.com/pocketbase/pocketbase"
)

var CONTENT_LINE = regexp.MustCompile(`[^\\/$\n\s•]+`)
var EMPTY_COMMENTS = regexp.MustCompile(`<!--\s*-->`)
var CDATA_CONTENT = regexp.MustCompile(`<!\[CDATA\[.*\]\]>`)

func HtmlContent(app *pocketbase.PocketBase, html string) (string, error) {

	html = HtmlContentPreFilters(html)

	markdown, err := htmltomarkdown.ConvertString(html)
	if err != nil {
		app.Logger().Error("[HTMLToMarkdown] cannot extract content", "error", err)
	}

	markdown = TextContentPostFilters(markdown)

	return markdown, nil
}

func HtmlContentPreFilters(input string) string {
	output := EMPTY_COMMENTS.ReplaceAllString(input, "")
	output = CDATA_CONTENT.ReplaceAllString(output, "")
	return output
}

func TextContentPostFilters(input string) string {
	output := RemoveMeaninglessLines(input)
	output = ReplaceDoubleQuotes(output)
	return output
}

// ReplaceDoubleQuotes
func ReplaceDoubleQuotes(input string) string {
	replacer := strings.NewReplacer(
		"''", `"`,
		"“", `"`,
		"”", `"`,
		"‘", `'`,
		"’", `'`,
	)
	return replacer.Replace(input)
}

// RemoveMeaninglessLines removes lines that only contain punctuation, \, /, bullets, and whitespace
func RemoveMeaninglessLines(input string) string {
	lines := strings.Split(input, "\n")
	var result []string

	for _, line := range lines {
		if !isMeaningless(line) {
			result = append(result, line)
		}
	}

	return strings.Join(result, "\n")
}

// isMeaningless checks if a line contains only punctuation, \, /, bullets, and whitespace
func isMeaningless(line string) bool {
	if len(strings.TrimSpace(line)) == 0 {
		return true
	}

	for _, r := range line {
		if !unicode.IsPunct(r) && !unicode.IsSpace(r) && CONTENT_LINE.MatchString(string(r)) {
			return false
		}
	}

	return true
}
