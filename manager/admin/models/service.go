package models

import (
	"manager/admin/types"
	"manager/container"
	"manager/libs"
	"os"
)

func GetAllServices() []container.Container {
	return container.FetchAllServices()
}

func IsActiveGateway() bool {
	return container.IsActiveGateway()
}

func CreateService(body types.CreateServiceBody) {
	content := container.GenerateContent(body)
	container.GenerateDockerfile(body.Service, content)
	libs.CopyFileTree(body.Artifact, container.GetServiceDir(body.Service))
	container.BuildAndCreate(body.Service)
}

func RunService(containerID string) {
	container.Run(containerID)
}

func StopService(containerID string) {
	container.StopContainer(containerID)
}

func RemoveService(service string) {
	container.RemoveContainerAndImage(service)
	os.RemoveAll(container.GetServiceDir(service))
}
