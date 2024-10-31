package container

type Container struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Port   string `json:"port"`
	Status string `json:"status"`
}

func GetAllContainers() []Container {
	return []Container{
		{
			ID:     "123456A",
			Name:   "customer",
			Port:   "12010",
			Status: "Running",
		}, {
			ID:     "4GREW56",
			Name:   "product",
			Port:   "12011",
			Status: "Running",
		},
		{
			ID:     "TT069F3G",
			Name:   "project",
			Port:   "12012",
			Status: "Running",
		},
	}
}