package models

import (
	"manager/container"
	"manager/libs"
)

func GetAllServices() []container.Container {

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

func CreateService(service string, port string, artifactPath string) {
	content := container.GenerateContent(service, nil)
	container.GenerateDockerfile(service, content)
	libs.CopyFileTree(artifactPath, container.GetServiceDir(service))
	container.BuildAndCreate(service, port)
}


func Run(service string) {
	// TODO Run
}

func Stop(service string) {
	// TODO Stop
}

func Remove(service string) {
	// TODO Remove
}