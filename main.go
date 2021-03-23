package main

import (
	"os"
	"Container/setup"
	"Container/models"
)

func main() {
	config :=
	models.Config{Env:[]string{"PATH=/bin:/usr/bin:/sbin:/usr/sbin:/:/usr/local/bin"},
	StartBinary:"/bin/sh",
	HostName:"c0ntain3r",
	RootFileSystem:"rootfs",
	BridgeIP:"10.0.0.1/24",
	HostInterface:"eth0",
	BridgeName:"c0ntain3r_bridge",
	}
	if len(os.Args) > 1 {
		setup.Child(config)
	}else{
		setup.Parent(config)
	}
}