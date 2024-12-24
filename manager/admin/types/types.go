package types

type Message struct {
	Message string `json:"message"`
}

type CreateServiceBody struct {
	Service  string `json:"service"`
	Base     string `json:"base"`
	Artifact string `json:"artifact"`
	Options  string `json:"options"`
	Execute  string `json:"execute"`
}

type InfrastructureInfo struct {
	Gateway  bool `json:"gateway"`
	Database bool `json:"database"`
}