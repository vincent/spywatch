package db

import (
	"time"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
)

var MAX_NARRATIVE_LENGTH = 100000
var MAX_CONTENT_LENGTH = 100000
var MAX_DIFF_LENGTH = 100000

func CreateSnapshot(app *pocketbase.PocketBase, resourceId string, content string, diff string, narrative string) error {

	collection, err := app.FindCollectionByNameOrId("snapshots")
	if err != nil {
		return err
	}

	if len(narrative) > MAX_NARRATIVE_LENGTH {
		narrative = narrative[:MAX_NARRATIVE_LENGTH]
	}
	if len(content) > MAX_CONTENT_LENGTH {
		content = content[:MAX_CONTENT_LENGTH]
	}
	if len(diff) > MAX_DIFF_LENGTH {
		diff = diff[:MAX_DIFF_LENGTH]
	}

	record := core.NewRecord(collection)
	record.Set("narrative", narrative)
	record.Set("resource", resourceId)
	record.Set("content", content)
	record.Set("diff", diff)

	return app.Save(record)
}

type ReleaseItem struct {
	WorkspaceName string `db:"workspaceName"`
	EntityName    string `db:"entityName"`
	EntityURL     string `db:"entityURL"`
	ResourceURL   string `db:"resourceURL"`
	SnapshotID    string `db:"id"`
	Created       string `db:"created"`
	Narrative     string `db:"narrative"`
}

func FindLastSnapshots(app *pocketbase.PocketBase, entityIds []string, since time.Time) ([]*ReleaseItem, error) {
	snapshots := []*ReleaseItem{}

	for _, eid := range entityIds {
		items := []*ReleaseItem{}
		err := app.DB().
			Select("snapshots.id", "snapshots.created", "snapshots.narrative").
			AndSelect("resources.url as resourceURL").
			AndSelect("bodies.name as entityURL", "bodies.name as entityName", "bodies.workspace as workspaceName").
			From("snapshots").
			InnerJoin("resources", dbx.NewExp("snapshots.resource = resources.id")).
			InnerJoin("bodies", dbx.NewExp("resources.body = bodies.id")).
			Where(dbx.NewExp("resources.url != ''")).
			Where(dbx.NewExp("snapshots.narrative != ''")).
			Where(dbx.NewExp("bodies.id = {:eid}", dbx.Params{"eid": eid})).
			Where(dbx.NewExp("snapshots.created > {:since}", dbx.Params{"since": since.Format(time.RFC3339)})).
			Limit(10).
			All(&items)

		if err != nil {
			app.Logger().Error("[FindLastSnapshots] cannot fetch snapshots", "error", err)
			return nil, err
		}
		snapshots = append(snapshots, items...)
	}

	return snapshots, nil
}

func BuildRelease(app *pocketbase.PocketBase, entityIds []string) ([]*ReleaseItem, error) {
	since := time.Now().Add(-30 * 24 * time.Hour)
	snapshots := []*ReleaseItem{}

	if len(entityIds) == 0 {
		return snapshots, nil
	}

	snapshots, err := FindLastSnapshots(app, entityIds, since)
	if err != nil {
		app.Logger().Error("[SendReleasesJob] cannot load snapshots for entities", "entities", entityIds, "error", err)
		return snapshots, err
	}

	return snapshots, nil
}
