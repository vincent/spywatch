package changedetection

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/oapi-codegen/runtime"
	"github.com/oapi-codegen/runtime/types"
)

type CreateWatchResponseBody struct {
	UUID types.UUID `json:"uuid"`
}

func CreateChangeDetectionClient() (*ClientWithResponses, error) {

	serverURL := os.Getenv("CHANGEDETECTION_SERVER_URL")
	apiKey := os.Getenv("CHANGEDETECTION_API_KEY")

	client, err := NewClientWithResponses(
		serverURL,
		WithRequestEditorFn(func(ctx context.Context, req *http.Request) error {
			req.Header.Set("X-Api-Key", apiKey)
			return nil
		}),
	)

	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	return client, nil
}

func IntPtr(pvar int) *int {
	return &pvar
}

func GetWatchSnapshotByString(ctx context.Context, c *Client, uuid types.UUID, timestamp string, params *GetWatchSnapshotParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetWatchSnapshotRequestByString(c.Server, uuid, timestamp, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// NewGetWatchSnapshotRequest generates requests for GetWatchSnapshot
func NewGetWatchSnapshotRequestByString(server string, uuid types.UUID, timestamp string, params *GetWatchSnapshotParams) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "uuid", runtime.ParamLocationPath, uuid)
	if err != nil {
		return nil, err
	}

	var pathParam1 string

	pathParam1, err = runtime.StyleParamWithLocation("simple", false, "timestamp", runtime.ParamLocationPath, timestamp)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/watch/%s/history/%s", pathParam0, pathParam1)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	if params != nil {
		queryValues := queryURL.Query()

		if params.Html != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "html", runtime.ParamLocationQuery, *params.Html); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
				return nil, err
			} else {
				for k, v := range parsed {
					for _, v2 := range v {
						queryValues.Add(k, v2)
					}
				}
			}

		}

		queryURL.RawQuery = queryValues.Encode()
	}

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}
