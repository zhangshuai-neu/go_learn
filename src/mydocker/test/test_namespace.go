package main

import (
	"log"
	"os"
	"os/exec"
	"syscall"
)

func main() {
	cmd := exec.Command("sh")
	// 配置namespace
	cmd.SysProcAttr.Cloneflags |= syscall.CLONE_NEWUTS
	cmd.SysProcAttr.Cloneflags |= syscall.CLONE_NEWIPC
	cmd.SysProcAttr.Cloneflags |= syscall.CLONE_NEWPID
	cmd.SysProcAttr.Cloneflags |= syscall.CLONE_NEWNS

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}
