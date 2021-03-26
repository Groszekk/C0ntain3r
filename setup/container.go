package setup

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
	"Container/models"
	"Container/network"
)

func Parent(config models.Config) {
	cmd := exec.Command("/proc/self/exe", append([]string{"child"}, []string{config.StartBinary}...)...)
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS|syscall.CLONE_NEWPID|syscall.CLONE_NEWNS|syscall.CLONE_NEWIPC|syscall.CLONE_NEWUSER,
		// mapping new user namespace to the ID of host user
		UidMappings: []syscall.SysProcIDMap{
			{
				ContainerID: 0,
				HostID:      os.Getuid(),
				Size:        1,
			},
		},
		GidMappings: []syscall.SysProcIDMap{
			{
				ContainerID: 0,
				HostID:      os.Getgid(),
				Size:        1,
			},
		},
	}
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Env = config.Env

	checkErr(cmd.Start())
	checkErr(network.SetNetworkInterfaces(cmd.Process.Pid, config))
	checkErr(cmd.Wait())
}

func Child(config models.Config) {
	cmd := exec.Command(config.StartBinary, nil...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	pivotRoot(config.RootFileSystem)
	checkErr(syscall.Sethostname([]byte(config.HostName)))

	if err := cmd.Run(); err != nil {
		fmt.Println("err", err.Error())
		os.Exit(2)
	}
}