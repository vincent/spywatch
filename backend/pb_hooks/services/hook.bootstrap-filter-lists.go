package services

import (
	"path"

	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
)

func BindBootstrapFiltersLists(app *pocketbase.PocketBase, hooksDir string) {
	app.OnBootstrap().BindFunc(func(e *core.BootstrapEvent) error {
		if err := e.Next(); err != nil {
			return err
		}

		LoadFiltersList(app, path.Join(hooksDir, BuiltinFilters))

		ykUrl := "https://raw.githubusercontent.com/yokoffing/filterlists/refs/heads/main/annoyance_list.txt"
		ykName := "yokoffing_annoyance_list.txt"
		ykFilters := path.Join(FiltersDir(), ykName)
		DownloadFilters(app, ykUrl, ykFilters)
		LoadFiltersList(app, ykFilters)

		return nil
	})
}
