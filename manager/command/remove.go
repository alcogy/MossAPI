package command

import (
	"manager/container"
	"os"

	"github.com/jmoiron/sqlx"
)

func RemoveService(service string, db *sqlx.DB) {
	container.RemoveContainerAndImage(service)
	os.RemoveAll(container.GetServiceDir(service))
}