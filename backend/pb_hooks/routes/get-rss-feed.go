package pb_routes

import (
	"net/http"
	"pocketbase/pb_hooks/db"
	"pocketbase/pb_hooks/services"

	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
)

// RegisterReleaseRoute registers the /api/config GET route.
func RegisterReleaseRoute(app *pocketbase.PocketBase) {
	app.OnServe().BindFunc(func(se *core.ServeEvent) error {
		se.Router.GET("/api/rss/{apiKey}/{bodyId}", func(e *core.RequestEvent) error {
			user, err := db.FindUserByApiKey(app, e.Request.PathValue("apiKey"))
			if err != nil {
				return e.XML(http.StatusForbidden, "Unrecognized api key")
			}
			app.Logger().Debug("[RssReleaseRoute] generating an rss feed for user within ws", "user", user.Id, "ws", user.Get("workspaces"))
			feed, err := services.GetEntityRssFeed(app, user, e.Request.PathValue("bodyId"))
			if err != nil {
				return e.XML(http.StatusInternalServerError, "Unknown feed source")
			}
			return e.XML(http.StatusOK, feed.FeedXml())
		})
		return se.Next()
	})
}
