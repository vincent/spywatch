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
				smartLinksResponse := ai.SimpleAIRequestAsJSON("You are a market analyser. here are all links from my competitor website. which links seem interesting to me ? keep similar (but not duplicates) links, ignore any system or technical links, categorize useful social networks links separately. give me only the lists, in json form { socialNetworks: [], ownLinks: [] }. ---- " + strings.Join(rc.SocialNetworks, "\n") + "" + strings.Join(rc.OwnLinks, "\n"))
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
