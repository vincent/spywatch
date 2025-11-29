package services

import (
	"io"
	"net/http"
	"os"
	"path"
	"strconv"
	"strings"

	"github.com/pocketbase/pocketbase"
)

var AllFilters = make(map[string][]string)

var BuiltinFilters = "filters_builtin.txt"

func FiltersDir() string {
	dir := strings.TrimSpace(os.Getenv("SPYWATCH_FILTERS_DIR"))
	if dir != "" {
		return dir
	}
	return path.Join(os.TempDir(), "spywatch", "filters")
}

func DownloadFilters(app *pocketbase.PocketBase, src string, dest string) error {
	if _, err := os.Stat(dest); os.IsNotExist(err) {
		// Ensure the directory exists
		if err := os.MkdirAll(FiltersDir(), 0755); err != nil {
			app.Logger().Error("[DownloadFilters] failed to create filtersDir: " + err.Error())
			return err
		}

		// Download the file
		resp, err := http.Get(src)
		if err != nil {
			app.Logger().Error("[DownloadFilters] failed to download filters: " + err.Error())
			return err
		}
		defer resp.Body.Close()

		out, err := os.Create(dest)
		if err != nil {
			app.Logger().Error("[DownloadFilters] failed to create filters: " + err.Error())
			return err
		}
		defer out.Close()

		if _, err := io.Copy(out, resp.Body); err != nil {
			app.Logger().Error("[DownloadFilters] failed to save filters: " + err.Error())
			return err
		}

		app.Logger().Info("[DownloadFilters] downloaded " + dest)
	}

	return nil
}

func LoadFiltersList(app *pocketbase.PocketBase, path string) error {
	file, err := os.ReadFile(path)
	if err != nil {
		app.Logger().Error("[LoadFiltersList] failed to read path", "error", err)
		return err
	}
	lines := strings.Split(string(file), "\n")
	d := 0
	s := 0
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" || strings.HasPrefix(line, ";") || strings.HasPrefix(line, "!") || strings.HasPrefix(line, "||") {
			continue
		}
		parts := strings.SplitN(line, "##", 2)
		if len(parts) == 2 {
			domain := strings.TrimSpace(parts[0])
			rules := strings.TrimSpace(parts[1])
			if domain != "" && rules != "" {
				if len(AllFilters[domain]) == 0 {
					d = d + 1
				}
				for _, selector := range splitAndTrim(rules, ",") {
					AllFilters[domain] = append(AllFilters[domain], selector)
					s = s + 1
				}
			}
		}
	}

	app.Logger().Info("[LoadFiltersList] loaded " + strconv.Itoa(s) + " filters on " + strconv.Itoa(d) + " domains from " + path)

	return nil
}

// splitAndTrim splits a string by sep and trims whitespace from each element.
func splitAndTrim(s, sep string) []string {
	parts := []string{}
	for _, part := range strings.Split(s, sep) {
		trimmed := strings.TrimSpace(part)
		if trimmed != "" {
			parts = append(parts, trimmed)
		}
	}
	return parts
}
