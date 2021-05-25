package k8s

import (
	"fmt"
	"os"
)

func GetServiceChart(directory string, app string) string {
	return directory + "charts/services/" + app
}

func GetSystemChart(directory string, app string) string {
	return directory + "charts/system/" + app
}

func GetCommonChart(directory string, app string) string {
	return directory + "charts/common/" + app
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

func GetValFilePath(chart string, environment string) (string, error) {
	valFile := fmt.Sprintf("%s-values.yaml", environment)
	filePath := chart + "/" + valFile
	_, err := os.Stat(filePath)
	if err != nil {
		return "", err
	}
	return filePath, nil
}

func GetDefaultValFilePath(chart string) (string, error) {
	valFile := "values.yaml"
	filePath := chart + "/" + valFile
	_, err := os.Stat(filePath)
	if err != nil {
		return "", err
	}
	return filePath, nil
}
