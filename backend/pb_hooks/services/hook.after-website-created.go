package services

import (
	"context"
	"encoding/json"
	"net/http"
	cd "pocketbase/pb_hooks/services/changedetection"

	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
)

func BindAfterWebsiteCreateSuccessHook(app *pocketbase.PocketBase) {

	app.OnRecordAfterCreateSuccess("competitors").BindFunc(func(e *core.RecordEvent) error {
		website := e.Record

		paused := true
		notificationMuted := true
		fetchBackend := cd.CreateWatchFetchBackendHtmlRequests
		method := cd.CreateWatchMethodGET

		checkInterval := struct {
			Days    *int `json:"days,omitempty"`
			Hours   *int `json:"hours,omitempty"`
			Minutes *int `json:"minutes,omitempty"`
			Seconds *int `json:"seconds,omitempty"`
			Weeks   *int `json:"weeks,omitempty"`
		}{
			Hours: intPtr(1),
		}

		name := website.GetString("name")
		url := website.GetString("url")

		watch, err := createChangeDetectionWatch(app, cd.CreateWatchJSONRequestBody{
			Url:               url,
			Title:             &name,
			Paused:            &paused,
			NotificationMuted: &notificationMuted,
			FetchBackend:      &fetchBackend,
			Method:            &method,
			TimeBetweenCheck:  &checkInterval,
		})

		if err != nil {
			app.Logger().Error("[OnRecordAfterCreateSuccess] Failed to create watch", "error", err)
			return err
		} else {
			app.Logger().Info("[OnRecordAfterCreateSuccess] Created watch", "result", watch.UUID)
		}

		website.Set("external_id", watch.UUID)
		err = app.Save(website)

		if err != nil {
			app.Logger().Error("[OnRecordAfterCreateSuccess] Failed to save watch external_id", "error", err)
			return err
		}

		return e.Next()
	})

	app.Logger().Info("BindAfterWebsiteCreateSuccessHook")
}

func createChangeDetectionWatch(app *pocketbase.PocketBase, watchRequest cd.CreateWatchJSONRequestBody) (*cd.CreateWatchResponseBody, error) {

	app.Logger().Debug("[createChangeDetectionWatch] creating watch")

	client, _ := cd.CreateChangeDetectionClient()

	resp, err := client.CreateWatchWithResponse(context.Background(), watchRequest)
	if err != nil {
		app.Logger().Error("[createChangeDetectionWatch] failed to create watch", "error", err)
		return nil, err
	}

	switch resp.StatusCode() {
	case http.StatusCreated, http.StatusOK:
		body := string(resp.Body)
		var responseModel = cd.CreateWatchResponseBody{}
		if err := json.Unmarshal(([]byte)(body), &responseModel); err != nil {
			app.Logger().Error("[createChangeDetectionWatch] cannot unmarshal response", "error", err)
		}
		return &responseModel, err
	default:
		app.Logger().Error("[createChangeDetectionWatch] unexpected status code", "code", resp.StatusCode())
	}

	return nil, err
}

func intPtr(i int) *int {
	return &i
}
