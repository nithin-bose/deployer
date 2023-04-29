package pkg

import (
	"bytes"
	"fmt"
	"os"
	"text/template"
)

func GetConfigFolderPath() string {
	return os.Getenv("HOME") + string(os.PathSeparator) + ConfigFolder
}

func FormatString(textTemplate string, values interface{}) (string, error) {
	tmpl, err := template.New("test").Parse(textTemplate)
	if err != nil {
		return "", err
	}

	output := new(bytes.Buffer)
	err = tmpl.Execute(output, values)
	if err != nil {
		return "", err
	}
	return output.String(), nil
}

func CreateConfigFile(fileName string, fileTemplate string, values interface{}) error {
	dirPath := GetConfigFolderPath()
	err := os.MkdirAll(dirPath, os.ModePerm)
	if err != nil {
		return err
	}

	outputFile := dirPath + string(os.PathSeparator) + fileName
	file, err := os.Create(outputFile)
	if err != nil {
		return err
	}

	tmpl, err := template.New("test").Parse(fileTemplate)
	if err != nil {
		return err
	}

	err = tmpl.Execute(file, values)
	if err != nil {
		return err
	}
	fmt.Println("File created at " + outputFile)
	return nil
}
