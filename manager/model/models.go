package model

type Flags struct {
	Command  string
	Service  string
	Port     string
	Artifact string
}

const PrivateNetworkName = "mossapi-nw-private"
const PublicNetworkName = "mossapi-nw-public"