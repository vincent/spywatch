// Package pb_routes provides custom route handlers for the application.
package pb_routes

import (
	"net/http"

	"pocketbase/pb_hooks/db"

	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
)

// RegisterWebsiteHistoryRoute registers the /api/history/:id GET route.
func RegisterWebsiteHistoryRoute(app *pocketbase.PocketBase) {
	app.OnServe().BindFunc(func(se *core.ServeEvent) error {
		se.Router.GET("/api/history/{websiteId}", func(e *core.RequestEvent) error {
			websiteId := e.Request.PathValue("websiteId")

			history, err := db.GetWebsiteLastSnapshots(app, websiteId, 2)

			if err != nil {
				app.Logger().Error("[GetWebsiteHistory] cannot find website history", "error", err)
				return e.JSON(http.StatusNotFound, err)
			}

			return e.JSON(http.StatusOK, history)
		})

		return se.Next()
	})
}
