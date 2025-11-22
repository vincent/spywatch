package scrapper

import (
	htmltomarkdown "github.com/JohannesKaufmann/html-to-markdown/v2"
	"github.com/pocketbase/pocketbase"
)

func HTMLToMarkdown(app *pocketbase.PocketBase, input string) (string, error) {
	markdown, err := htmltomarkdown.ConvertString(input)
	if err != nil {
		app.Logger().Error("[HTMLToMarkdown] cannot extract content", "error", err)
	}
	return markdown, nil
}
