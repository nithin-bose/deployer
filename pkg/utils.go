package pkg

import (
	"fmt"
	"os"
	"text/template"
)

func CreateMobFile(fileName string, fileTemplate string, values interface{}) error {
	dirPath := os.Getenv("HOME") + string(os.PathSeparator) + ConfigFolder
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
