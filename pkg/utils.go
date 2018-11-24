package pkg

import (
	"fmt"
	"log"
	"os"
	"text/template"
)

func FatalF(format string, a ...interface{}) {
	log.Panicf(format, a...)
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
