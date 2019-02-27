package pkg

const (
	AppName      = "deployer"
	Version      = "3.1.0"
	ConfigFolder = ".deployer"

	//Charts
	// TraefikChart = "traefik-server"

	//Helm user service account
	HelmServiceUser = "helm"

	// Gitlab
	GitlabServer = "https://gitlab.com/api/v4"
	MaxChanges   = 50
	BodyTemplate = `
{{if .ReleaserName }}
Released by: **{{ .ReleaserName }}**
{{end}}

The changes are as follows:

{{range .Changes}}
* {{ . }}
{{end}}

---
Automated with https://gitlab.com/nithinbose/deployer
`
)
