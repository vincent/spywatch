package services

import (
	"pocketbase/pb_hooks/db"

	"github.com/gorilla/feeds"
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
)

func GetEntityRssFeed(app *pocketbase.PocketBase, authRecord *core.Record, entityId string) (*feeds.RssFeed, error) {
	body, err := app.FindFirstRecordByFilter(
		"bodies",
		"id = {:id} && ("+db.InSlice("workspace.id", authRecord.GetStringSlice("workspaces"))+")",
		dbx.Params{"id": entityId},
	)
	if err != nil {
		app.Logger().Error("[getEntityFeed] unknown entity", "entityId", entityId, "error", err)
		return nil, err
	}
	items, err := app.FindRecordsByFilter(
		"snapshots",
		"resource.body = {:eid} && narrative != ''",
		"-created",
		10,
		0,
		dbx.Params{"eid": entityId},
	)
	if err != nil {
		app.Logger().Error("[getEntityFeed] cannot query entity items", "entityId", entityId, "error", err)
		return nil, err
	}

	app.Logger().Debug("[getEntityFeed] got items for entity", "entity", body.Get("name"), "count", len(items))

	feed := &feeds.Feed{
		Title:       "SpyWatch: " + body.GetString("name"),
		Link:        &feeds.Link{Href: "http://spywatch/api/rss/" + body.Id},
		Description: "SpyWatch feed of " + body.GetString("name"),
		Author:      &feeds.Author{Name: "SpyWatch"},
		Created:     body.GetDateTime("last_check").Time(),
	}

	var feedItems []*feeds.Item
	for i := 0; i < len(items); i++ {
		item := items[i]
		feedItems = append(feedItems,
			&feeds.Item{
				Id:          item.Id,
				Title:       "Updated on " + body.GetString("created"),
				Link:        &feeds.Link{Href: "feeds item link"},
				Description: item.GetString("narrative"),
				Created:     body.GetDateTime("created").Time(),
			})
	}
	feed.Items = feedItems

	return (&feeds.Rss{Feed: feed}).RssFeed(), nil
}
