package models

import (
	"manager/container"
)

func GetAllContainers() []container.Container {

	containers := container.AllContainers()
	return containers
	// return []container.Container{
	// 	{
	// 		ID:     "123456A",
	// 		Name:   "customer",
	// 		Port:   "12010",
	// 		Status: "Running",
	// 	}, {
	// 		ID:     "4GREW56",
	// 		Name:   "product",
	// 		Port:   "12011",
	// 		Status: "Running",
	// 	},
	// 	{
	// 		ID:     "TT069F3G",
	// 		Name:   "project",
	// 		Port:   "12012",
	// 		Status: "Running",
	// 	},
	// }
}