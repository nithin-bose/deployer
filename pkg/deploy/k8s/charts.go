package k8s

func GetServiceChart(directory string, app string) string {
	return directory + "charts/services/" + app
}

func GetSystemChart(directory string, app string) string {
	return directory + "charts/system/" + app
}

func GetInfraChart(directory string, cloudPlatorm string, app string) string {
	if app == "setup" {
		app = app + "-" + cloudPlatorm
	}
	return directory + "charts/infra/" + app
}

func GetAdminPanelChart(directory string, app string) string {
	return directory + "charts/admin-panels/" + app
}

func GetValFilePath(chart string, file string) string {
	return chart + "/" + file
}
