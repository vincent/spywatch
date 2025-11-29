// Package pb_routes provides custom route handlers for the application.
package pb_routes

import (
	"encoding/json"
	"net/http"
	"strings"

	"pocketbase/pb_hooks/services/ai"
	"pocketbase/pb_hooks/services/scrapper"

	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
)

var EXTRACT_LINKS_PROMPT = `
You are a content analyser. 

GOAL:
Here are all links from a website I am interested in. 
Which links seem interesting to me ?
Give me only the lists, in json form { socialNetworks: [], ownLinks: [] }.

RECOMMANDATIONS:
Keep similar (but not duplicates) links. 
Ignore any system or technical links. 
Categorize useful social networks links separately.
---
`

// competitorPreviewRequest represents the expected request body.
type competitorPreviewRequest struct {
	Url string `json:"url"`
}

// RegisterCompetitorPreviewRoute registers the /api/preview/:id POST route.
func RegisterCompetitorPreviewRoute(app *pocketbase.PocketBase) {
	app.OnServe().BindFunc(func(se *core.ServeEvent) error {
		se.Router.POST("/api/preview", func(e *core.RequestEvent) error {

			// Parse request body
			var form competitorPreviewRequest
			if err := e.BindBody(&form); err != nil {
				return e.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
			}

			rc, err := scrapper.GetLinksFromWebsite(form.Url)
			if err != nil {
				app.Logger().Error("[PostCompetitorPreview] cannot get website links", "error", err)
				return e.JSON(http.StatusInternalServerError, "cannot get website links")
			}

			if false {
				smartLinksResponse := ai.SimpleAIRequestAsJSON(EXTRACT_LINKS_PROMPT + strings.Join(rc.SocialNetworks, "\n") + "\n" + strings.Join(rc.OwnLinks, "\n"))
				app.Logger().Info("[PostCompetitorPreview] smartLinks", "smartLinks", smartLinksResponse)

				if err := json.Unmarshal(([]byte)(smartLinksResponse), &rc); err != nil {
					app.Logger().Error("[PostCompetitorPreview] cannot get smart links", "error", err)
				}
			}

			return e.JSON(http.StatusOK, rc)
		})

		return se.Next()
	})
}
