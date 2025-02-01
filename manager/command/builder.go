package command

import (
	"encoding/json"
	"fmt"
	"manager/admin/models"
	"manager/admin/types"
	"manager/container"
	"manager/libs"
	"manager/table"
	"os"

	"regexp"

	"github.com/jmoiron/sqlx"
	"gopkg.in/yaml.v3"
)

func ExecuteBuild(path string, db *sqlx.DB) {
	backend := readFile(path)
	for _, s := range backend.Services {
		buildService(s)
	}

	for _, t := range backend.Tables {
		buildTable(t, db)
	}
}

func readFile(path string) Backend {
	fileInfo, err := os.Stat(path)
	if err != nil {
		panic(err)
	}

	if fileInfo == nil || fileInfo.IsDir() {
		panic("File path is not correct.")
	}

	file, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	var backend Backend
	if isYaml(path) {
		err = yaml.Unmarshal(file, &backend)
	} else {
		err = json.Unmarshal(file, &backend)
	}

	if err != nil {
		panic(err)
	}

	return backend
}

func buildService(body types.CreateServiceBody) {
	content := container.GenerateContent(body)
	container.GenerateDockerfile(body.Service, content)
	libs.CopyFileTree(body.Artifact, container.GetServiceDir(body.Service))
	container.BuildAndCreate(body.Service)
}

func buildTable(table table.Table, db *sqlx.DB) {
	if len(table.Columns) == 0 {
		return
	}
	models.CreateTable(db, table)
}

func isYaml(path string) bool {
	pattern := `.yml$`
	re, err := regexp.Compile(pattern)
	if err != nil {
		fmt.Println("Error compiling regex:", err)
		return false
	}
	matches := re.FindAllString(path, -1)

	return len(matches) > 0
}
