package db

import (
	"os"
	"strconv"
	"time"

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

// CreateWebsite adds a new row to the bodies collection.
func CreateWebsite(app *pocketbase.PocketBase, model *CreateWebsiteInput) (*core.Record, error) {

	collection, err := app.FindCollectionByNameOrId("bodies")
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

func FindUncheckedResources(app *pocketbase.PocketBase, entityId string) ([]*core.Record, error) {
	backStr := os.Getenv("SPYWATCH_REFETCH_HOURS")
	back, err := strconv.Atoi(backStr)
	if err != nil || back <= 0 {
		back = 24
	}
	resources, err := app.FindRecordsByFilter(
		"resources",
		"url != '' && body = {:body} && checked < {:checked}",
		"created",
		10,
		0,
		dbx.Params{"body": entityId},
		dbx.Params{"checked": time.Now().Add(-1 * time.Duration(back) * time.Hour).Format(time.RFC3339)},
	)
	if err != nil {
		app.Logger().Error("[FindBodiesResources] cannot fetch resources", "error", err)
		return nil, err
	}
	return resources, nil
}

func FindUpdatedResourcesByEntities(app *pocketbase.PocketBase, entityIds []string) ([]*core.Record, error) {
	resources, err := app.FindRecordsByFilter(
		"resources",
		"url != '' && body.id ?= {:ws} && updated >= {:updated}",
		"updated",
		10,
		0,
		dbx.Params{"ws": entityIds},
		dbx.Params{"updated": time.Now().Add(-1 * 30 * time.Hour).Format(time.RFC3339)},
	)
	if err != nil {
		app.Logger().Error("[FindUpdatedResourcesByEntities] cannot fetch resources", "error", err)
		return nil, err
	}
	return resources, nil
}

func FindBodies(app *pocketbase.PocketBase) ([]*core.Record, error) {
	bodies, err := app.FindAllRecords("bodies")
	if err != nil {
		app.Logger().Error("[FindBodies] cannot fetch websites", "error", err)
		return nil, err
	}
	return bodies, nil
}

type WorkspaceForRelease struct {
	Body          string `db:"body"`
	WorkspaceID   string `db:"workspaceID"`
	WorkspaceName string `db:"workspaceName"`
}

func FindWorkspacesForReleases(app *pocketbase.PocketBase) ([]WorkspaceForRelease, error) {
	workspaces := []WorkspaceForRelease{}
	err := app.DB().
		Select("bodies.id as body", "bodies.workspace as workspaceID", "workspaces.name as workspaceName").
		Distinct(true).
		From("bodies").
		InnerJoin("workspaces", dbx.NewExp("workspaces.id = bodies.workspace")).
		Where(dbx.NewExp("workspace != ''")).
		All(&workspaces)
	return workspaces, err
}
