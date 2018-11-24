package pkg

import (
	"fmt"
	"os"
	"text/template"
)

func FatalF(format string, a ...interface{}) {
	fmt.Fprintf(os.Stderr, format, a...)
	os.Exit(2)
}

func CreateMobFile(fileName string, fileTemplate string, values interface{}) {
	dirPath := os.Getenv("HOME") + string(os.PathSeparator) + ConfigFolder
	err := os.MkdirAll(dirPath, os.ModePerm)
	if err != nil {
		FatalF("An error occurred:\n %s \n", err.Error())
	}

	outputFile := dirPath + string(os.PathSeparator) + fileName
	file, err := os.Create(outputFile)
	if err != nil {
		FatalF("An error occurred:\n %s \n", err.Error())
	}

	tmpl, err := template.New("test").Parse(fileTemplate)
	if err != nil {
		FatalF("An error occurred:\n %s \n", err.Error())
	}

	err = tmpl.Execute(file, values)
	if err != nil {
		FatalF("An error occurred:\n %s \n", err.Error())
	}
	fmt.Println("File created at " + outputFile)
}
