package models

type Config struct {
	Env []string
	StartBinary string
	HostName string
	RootFileSystem string
	BridgeIP string
	HostInterface string
	BridgeName string
	VethPair VethPair
}