package container

type Container struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	State  string `json:"state"`
	Status string `json:"status"`
}

type ContainerFull struct {
	ID      string   `json:"id"`
	Names   []string `json:"names"`
	ImageID string   `json:"imageId"`
	Command string   `json:"command"`
	Created int64    `json:"created"`
	State   string   `json:"state"`
	Status  string   `json:"status"`
}