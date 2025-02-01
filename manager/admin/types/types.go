package types

type Message struct {
	Message string `json:"message"`
}

// CreateServiceBody is build service(container) structure data.
type CreateServiceBody struct {
	Service  string `json:"service"`
	Base     string `json:"base"`
	Artifact string `json:"artifact"`
	Options  string `json:"options"`
	Execute  string `json:"execute"`
}

// InfrastructureInfo is for signal admin top.
type InfrastructureInfo struct {
	Gateway  bool `json:"gateway"`
	Database bool `json:"database"`
}
