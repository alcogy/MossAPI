package model

type Flags struct {
	Command  string
	Service  string
	Artifact string
	Execute  string
}

const PrivateNetworkName = "mossapi-nw-private"
const PublicNetworkName = "mossapi-nw-public"