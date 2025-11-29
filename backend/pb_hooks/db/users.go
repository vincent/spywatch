package db

import (
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
)

func FindUserByApiKey(app *pocketbase.PocketBase, apiKey string) (*core.Record, error) {
	return app.FindFirstRecordByData("users", "api_key", apiKey)
}

func InSlice(field string, items []string) string {
	exp := "(1=0"
	for i := 0; i < len(items); i++ {
		exp = exp + " || " + field + "='" + items[i] + "'"
	}
	exp = exp + ")"
	return exp
}
