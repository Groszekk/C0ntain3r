package setup

import (
	"os"
	"path/filepath"
	"syscall"
)

func pivotRoot(newRoot string) {
	old := filepath.Join(newRoot, "/.p_root")
	checkErr(syscall.Mount(newRoot, newRoot, "", syscall.MS_BIND|syscall.MS_REC, ""))
	checkErr(os.MkdirAll(old, 0700))
	checkErr(syscall.PivotRoot(newRoot, old))
	checkErr(os.Chdir("/"))
	checkErr(syscall.Mount("proc", "/proc", "proc", 0, ""))
	checkErr(syscall.Unmount("/.p_root", syscall.MNT_DETACH))
	checkErr(os.RemoveAll("/.p_root"))
}