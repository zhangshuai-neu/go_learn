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
	cmd.SysProcAttr = &attr

	// 配置用户命名空间
	uidMap := syscall.SysProcIDMap{}
	uidMap.ContainerID = 1994
	uidMap.HostID = 0
	uidMap.Size = 1
	uidMapSlice := []syscall.SysProcIDMap{}
	uidMapSlice = append(uidMapSlice, uidMap)

	gidMap := syscall.SysProcIDMap{}
	gidMap.ContainerID = 1994
	gidMap.HostID = 0
	gidMap.Size = 1
	gidMapSlice := []syscall.SysProcIDMap{}
	gidMapSlice = append(uidMapSlice, gidMap)

	cmd.SysProcAttr.UidMappings = uidMapSlice
	cmd.SysProcAttr.GidMappings = gidMapSlice

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	os.Exit(0)
}
