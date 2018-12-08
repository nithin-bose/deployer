package pkg

const (
	AppName          = "deployer"
	Version          = "2.0.2"
	ConfigFolder     = ".deployer"
	ConfigFolderPath = "~/.deployer"

	//Charts
	// TraefikChart = "traefik-server"

	//Helm user service account
	HelmServiceUser = "helm"

	// Helm deploy constants
	StagingValsFile    = "values.yaml"
	ProductionValsFile = "prod-values.yaml"

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
Automated with github.com/nithin-bose/deployer
`
)
