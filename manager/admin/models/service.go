package models

import (
	"manager/admin/types"
	"manager/container"
	"manager/database/redis"
	"manager/libs"
)

func GetAllServices() []container.Container {
	return container.AllContainers()
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
	redis.DeleteService(service)
}
