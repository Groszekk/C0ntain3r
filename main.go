package main

import (
	"os"
	"Container/setup"
	"Container/models"
)

func main() {
	// example
	vethPair := models.VethPair{
		Veth: nil,
		Name: "cveth",
		Peer: nil,
		IP: "10.0.0.2/24",
		PeerIP: "10.0.0.3/24",
		PeerName: "cvethp",
	}
	config :=
	models.Config{Env:[]string{"PATH=/bin:/usr/bin:/sbin:/usr/sbin:/:/usr/local/bin"},
	StartBinary:"/bin/sh",
	HostName:"c0ntain3r",
	RootFileSystem:"rootfs",
	BridgeIP:"10.0.0.1/24",
	HostInterface:"eth0",
	BridgeName:"c0ntain3r_bridge",
	VethPair: vethPair,
	}
	if len(os.Args) > 1 {
		setup.Child(config)
	}else{
		setup.Parent(config)
	}
}