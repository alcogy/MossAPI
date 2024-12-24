package command

import (
	"encoding/json"
	"manager/admin/models"
	"manager/admin/types"
	"manager/container"
	"manager/libs"
	"manager/table"
	"os"

	"github.com/jmoiron/sqlx"
)

type Backend struct {
	Services []types.CreateServiceBody `json:"services"`
	Tables   []table.Table             `json:"tables"`
}

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
	err = json.Unmarshal(file, &backend)
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