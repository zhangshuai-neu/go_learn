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
	attr := syscall.SysProcAttr{}
	attr.Cloneflags |= syscall.CLONE_NEWUTS
	attr.Cloneflags |= syscall.CLONE_NEWIPC
	attr.Cloneflags |= syscall.CLONE_NEWPID
	attr.Cloneflags |= syscall.CLONE_NEWNS
	attr.Cloneflags |= syscall.CLONE_NEWUSER
	attr.Cloneflags |= syscall.CLONE_NEWNET
	cmd.SysProcAttr = &attr

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	os.Exit(0)
}
