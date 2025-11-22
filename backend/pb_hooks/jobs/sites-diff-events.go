package jobs

import (
	"pocketbase/pb_hooks/db"
	"pocketbase/pb_hooks/services"
	"pocketbase/pb_hooks/services/scrapper"
	"time"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase"
)

func RegisterDiffDescriptionsJob(app *pocketbase.PocketBase) {
	app.Cron().MustAdd("DiffDescriptionsJob", "*/10 * * * *", func() {
		websites, _ := db.FindCompetitors(app)
		for _, website := range websites {
			resources, _ := db.FindCompetitorResources(app, website.Id)
			if len(resources) == 0 {
				app.Logger().Debug("[DiffDescriptionsJob] no resources for competitor", "competitor", website.Id)
			}
			for _, res := range resources {
				CreateResourceSnapshot(app, res.Id)
			}
		}
	})
}

func CreateResourceSnapshot(app *pocketbase.PocketBase, resourceId string) error {

	res, err := app.FindRecordById("resources", resourceId)
	if err != nil {
		app.Logger().Error("[CreateResourceSnapshot] cannot find resource", "error", err)
		return nil
	}

	url := res.GetString("url")
	html, err := scrapper.FetchHTML(url)
	if err != nil {
		app.Logger().Error("[CreateResourceSnapshot] cannot extract content", "error", err)
		return nil
	}
	newContent, err := scrapper.HTMLToMarkdown(app, html)
	if err != nil {
		app.Logger().Error("[CreateResourceSnapshot] cannot extract content", "error", err)
		return nil
	}

	lastSnapshots, err := app.FindRecordsByFilter(
		"snapshots",
		"resource = {:resource}",
		"-created",
		1,
		0,
		dbx.Params{"resource": resourceId},
	)
	if err != nil {
		app.Logger().Error("[CreateResourceSnapshot] cannot find previous content", "error", err)
		return nil
	}

	diff := ""
	narrative := ""
	lastContent := ""
	snapshot := false

	if len(lastSnapshots) == 1 {
		lastContent = lastSnapshots[0].GetString("content")
		if lastContent != newContent {
			diff = services.DiffDescription(newContent, lastContent)
			if diff != "" {
				narrative, err = services.DiffNarrative(diff)
				if err != nil {
					app.Logger().Error("[CreateResourceSnapshot] cannot get a narrative", "error", err)
					return err
				}
			}
		}
		if narrative != "" {
			app.Logger().Debug("[CreateResourceSnapshot] insert resource snapshot with narrative", "resource", url, "narrative", narrative)
			snapshot = true
		} else {
			app.Logger().Debug("[CreateResourceSnapshot] ignore unchanged resource", "resource", url)
		}
	} else {
		app.Logger().Debug("[CreateResourceSnapshot] first resource snapshot", "resource", url)
		snapshot = true
	}

	if snapshot {
		err = db.CreateSnapshot(app, resourceId, newContent, diff, narrative)
		if err != nil {
			app.Logger().Error("[CreateResourceSnapshot] cannot save diff description", "error", err)
			return err
		}
	}

	res.Set("checked", time.Now().Format(time.RFC3339))
	app.Save(res)

	return nil
}
