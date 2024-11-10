package command

import (
	"encoding/json"
	"fmt"
	"manager/admin/types"
	"manager/database/mysql"
	"os"

	"github.com/jmoiron/sqlx"
)

type Backend struct {
	Services []types.CreateServiceBody `json:"services"`
	Tables []mysql.Table	`json:"tables"`
}

func ExecuteBuild(path string, db *sqlx.DB) {
	backend := readFile(path)
	for _, s := range backend.Services {
		fmt.Println(s)
	}

	for _, t := range backend.Tables {
		fmt.Println(t)
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