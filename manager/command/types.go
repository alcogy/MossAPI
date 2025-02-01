package command

import (
	"manager/admin/types"
	"manager/container"
	"manager/table"
)

type DumpModel struct {
	Services []container.ContainerFull `json:"services"`
	Tables   []table.Table             `json:"tables"`
}

type Backend struct {
	Services []types.CreateServiceBody `json:"services"`
	Tables   []table.Table             `json:"tables"`
}
