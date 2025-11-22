package db

import (
	"context"
	"encoding/json"
	cd "pocketbase/pb_hooks/services/changedetection"
	"time"

	"github.com/google/uuid"
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
)

// CreateWebsiteInput represents the model of a new website.
type CreateWebsiteInput struct {
	Name       string
	URL        string
	ExternalID string
	OwnerID    string
}

// CreateWebsite adds a new row to the competitors collection.
func CreateWebsite(app *pocketbase.PocketBase, model *CreateWebsiteInput) (*core.Record, error) {

	collection, err := app.FindCollectionByNameOrId("competitors")
	if err != nil {
		return nil, err
	}

	record := core.NewRecord(collection)
	record.Set("name", model.Name)
	record.Set("url", model.URL)
	record.Set("external_id", model.ExternalID)
	record.Set("owner", model.OwnerID)

	err = app.Save(record)
	if err != nil {
		app.Logger().Error("[CreateWebsite] cannot insert record", "error", err)
		return nil, err
	}

	app.Logger().Debug("[CreateWebsite] insert record", "id", record.Id)

	return record, nil
}

func GetWebsiteLastSnapshots(app *pocketbase.PocketBase, websiteId string, max int) ([]string, error) {
	website, err := app.FindRecordById("competitors", websiteId)

	if err != nil {
		app.Logger().Error("[GetWebsiteHistory] cannot find website", "error", err)
		return nil, err
	}

	externalId := website.GetString("external_id")
	if externalId == "" {
		app.Logger().Info("[GetWebsiteHistory] website has no external_id")
		return nil, err
	}

	client, _ := cd.CreateChangeDetectionClient()

	resp, err := client.GetWatchHistoryWithResponse(context.Background(), uuid.MustParse(externalId))
	if err != nil {
		app.Logger().Error("[GetWebsiteHistory] failed to get watch history", "error", err)
		return nil, err
	}

	snapshots := cd.WatchHistory{}
	if err := json.Unmarshal(resp.Body, &snapshots); err != nil {
		app.Logger().Error("[GetWebsiteHistory] cannot unmarshal response", "error", err)
		return nil, err
	}

	app.Logger().Debug("[GetWebsiteHistory] got snapshots", "snapshots", snapshots)

	c := 0
	history := []string{}
	for snapshot, _ := range snapshots {
		resp, _ := client.GetWatchSnapshotWithResponse(context.Background(), uuid.MustParse(externalId), snapshot, &cd.GetWatchSnapshotParams{})
		history = append(history, string(resp.Body))

		c = c + 1
		if c >= max {
			break
		}
	}

	return history, nil
}

func FindCompetitorResources(app *pocketbase.PocketBase, competitorId string) ([]*core.Record, error) {
	resources, err := app.FindRecordsByFilter(
		"resources",
		"competitor = {:competitor} && checked <= {:checked}",
		"-created",
		10,
		0,
		dbx.Params{"competitor": competitorId},
		dbx.Params{"checked": time.Now().Add(24 * time.Hour).Format(time.RFC3339)},
	)
	if err != nil {
		app.Logger().Error("[FindCompetitorResources] cannot fetch resources", "error", err)
		return nil, err
	}
	return resources, nil
}

func FindCompetitors(app *pocketbase.PocketBase) ([]*core.Record, error) {
	competitors, err := app.FindAllRecords("competitors")
	if err != nil {
		app.Logger().Error("[FindCompetitors] cannot fetch websites", "error", err)
		return nil, err
	}
	return competitors, nil
}

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
