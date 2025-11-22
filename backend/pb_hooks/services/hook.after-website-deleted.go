package services

import (
	"context"
	cd "pocketbase/pb_hooks/services/changedetection"

	"github.com/google/uuid"
	"github.com/oapi-codegen/runtime/types"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
)

func BindAfterWebsiteDeleteSuccessHook(app *pocketbase.PocketBase) {

	app.OnRecordAfterDeleteSuccess("competitors").BindFunc(func(e *core.RecordEvent) error {
		website := e.Record

		externalId := website.GetString("external_id")
		if externalId == "" {
			app.Logger().Info("[OnRecordAfterDeleteSuccess] ignore website without UUID")
			return e.Next()
		}

		app.Logger().Debug("[OnRecordAfterDeleteSuccess] try to delete website watch", "string", externalId)

		err := deleteChangeDetectionWatch(app, uuid.MustParse(externalId))
		if err != nil {
			app.Logger().Error("[OnRecordAfterDeleteSuccess] failed to delete watch", "error", err)
			// return err
		}

		return e.Next()
	})

	app.Logger().Info("BindAfterWebsiteDeleteSuccessHook")
}

func deleteChangeDetectionWatch(app *pocketbase.PocketBase, watchId types.UUID) error {

	app.Logger().Debug("[deleteChangeDetectionWatch] creating watch")

	client, _ := cd.CreateChangeDetectionClient()

	_, err := client.DeleteWatch(context.Background(), watchId)
	if err != nil {
		app.Logger().Error("[deleteChangeDetectionWatch] failed to delete watch", "error", err)
		return err
	}

	return nil
}
