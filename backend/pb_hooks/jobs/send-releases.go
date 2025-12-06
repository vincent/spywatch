package jobs

import (
	"net/mail"
	"pocketbase/pb_hooks/db"
	"strings"

	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/tools/mailer"
	"github.com/pocketbase/pocketbase/tools/template"
)

func RegisterSendReleasesJob(app *pocketbase.PocketBase, hooksDir string) {
	app.Cron().MustAdd("SendReleasesJob", "0 6 * * *", func() {
		assignments, err := db.FindWorkspacesForReleases(app)
		if err != nil {
			app.Logger().Error("[SendReleasesJob] cannot load workspace assignments", "error", err)
			return
		}
		if len(assignments) == 0 {
			app.Logger().Info("[SendReleasesJob] nothing to do")
			return
		}

		app.Logger().Error("[SendReleasesJob] starting releases on workspaces", "workspaces", len(assignments))

		// group entities by workspace
		workspaceBodies := map[string][]string{}
		for _, a := range assignments {
			if a.WorkspaceID == "" || a.Body == "" {
				continue
			}
			workspaceBodies[a.WorkspaceID] = append(workspaceBodies[a.WorkspaceID], a.Body)
		}

		// iter on workspaces
		for _, w := range assignments {

			// find user we'll send the release
			users, err := db.FindUsersByWorkspace(app, w.WorkspaceID)
			if err != nil {
				app.Logger().Error("[SendReleasesJob] cannot load users", "error", err)
				continue
			}
			if len(users) == 0 {
				app.Logger().Error("[SendReleasesJob] no users following workspace", "workspace", w.WorkspaceID)
				continue
			}

			// build the release
			release, err := db.BuildRelease(app, workspaceBodies[w.WorkspaceID])
			if err != nil {
				app.Logger().Error("[SendReleasesJob] cannot build release of workspace", "workspace", w.WorkspaceID, "error", err)
				continue
			}
			if len(release) == 0 {
				app.Logger().Error("[SendReleasesJob] built release of workspace is empty", "workspace", w.WorkspaceID)
				continue
			}

			for _, user := range users {
				err = SendReleaseMail(app, hooksDir, w.WorkspaceName, release, user)
				if err != nil {
					app.Logger().Error("[SendReleasesJob] cannot send release of workspace to user", "user", user.Id, "workspace", w.WorkspaceName, "error", err)
					continue
				}
				app.Logger().Info("[SendReleasesJob] sent release of workspace to user", "user", user.Id, "workspace", w.WorkspaceName)
			}
		}
	})
}

func BuildReleaseMailBody(release []*db.ReleaseItem) string {
	html := ""
	lastEntity := ""
	for _, s := range release {
		if lastEntity != s.EntityName {
			if lastEntity != "" {
				html = html + "<hr>"
			}
			html = html + "<h2>" + s.EntityName + "</h2>"
		}
		lastEntity = s.EntityName

		html = html + "<i>" + s.Created + "</i><br>"
		html = html + strings.ReplaceAll(s.Narrative, "\n", "<br>")
		html = html + "<br><br>"
	}
	return html
}

func SendReleaseMail(app *pocketbase.PocketBase, hooksDir string, workspace string, release []*db.ReleaseItem, recipient db.User) error {
	body := "<strong>" + recipient.Name + "</strong>, we have news on your workspace &quot;" + workspace + "&quot;<br><br>" + BuildReleaseMailBody(release)

	registry := template.NewRegistry()
	html, err := registry.LoadFiles(
		hooksDir+"/views/emails/layout.html",
		hooksDir+"/views/emails/content.html",
	).Render(map[string]any{
		"body": body,
	})

	if err != nil {
		app.Logger().Error("[SendReleaseMail] cannot build release mail", "error", err)
		return err
	}

	message := &mailer.Message{
		From: mail.Address{
			Address: app.Settings().Meta.SenderAddress,
			Name:    app.Settings().Meta.SenderName,
		},
		To:      []mail.Address{{Address: recipient.Email}},
		Subject: "Updates on your workspace " + workspace,
		HTML:    html,
		// bcc, cc, attachments and custom headers are also supported...
	}

	return app.NewMailClient().Send(message)
}
