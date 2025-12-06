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

type User struct {
	Id    string `db:"id"`
	Name  string `db:"name"`
	Email string `db:"email"`
}

func FindUsersByWorkspace(app *pocketbase.PocketBase, workspace string) ([]User, error) {
	users := []User{}
	err := app.DB().
		NewQuery(`SELECT users.id, users.name, users.email FROM users LEFT JOIN json_each(users.workspaces) je WHERE je.value = '` + workspace + `'`).
		All(&users)
	return users, err
}
