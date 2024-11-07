package types

type Message struct {
	Message string `json:"message"`
}

type CreateServiceBody struct {
	Service  string `json:"service"`
	Artifact string `json:"artifact"`
	Options  string `json:"options"`
	Execute  string `json:"execute"`
}
