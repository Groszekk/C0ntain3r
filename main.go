package main

import (
	"os"
	"Container/setup"
	"Container/models"
)

func main() {
	config :=
	models.Config{Env:[]string{"PATH=/bin:/usr/bin:/sbin:/usr/sbin:/:/usr/local/bin"}, StartBinary:"/bin/sh", HostName:"c0ntain3r", RootFileSystem:"rootfs",}
	switch os.Args[1] { // little hack for re-execution
	case "run":
		setup.Parent(config)
	case "child":
		setup.Child(config)
	}
}