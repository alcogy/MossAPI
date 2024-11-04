package models

import (
	"manager/container"
	"manager/database/redis"
	"manager/libs"
)

func GetAllServices() []container.Container {
	return container.AllContainers()
}

func CreateService(service string, port string, artifactPath string) {
	content := container.GenerateContent(service, nil)
	container.GenerateDockerfile(service, content)
	libs.CopyFileTree(artifactPath, container.GetServiceDir(service))
	container.BuildAndCreate(service, port)
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
