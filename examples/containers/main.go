package main

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

//how to run a container in docker: docker run <image> <command>
//how we will run a container: go run main.go run <command>
func main() {
	switch os.Args[1] {
	case "run":
		run()
	case "child":
		child()
	default:
		panic("idk what you want homie")
	}
}

func run() {
	//args[2] is the main command, args[3:] are the arguments
	cmd := exec.Command("/proc/self/exe", append([]string{"child"}, os.Args[2:]...)...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	//namespaces
	cmd.SysProcAttr = &syscall.SysProcAttr{
		//clone_newuts is used to clone the unix time
		Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID | syscall.CLONE_NEWNS,
	}

	checkErr(cmd.Run())
}

func child() {
	fmt.Printf("running %v as PID %d\n", os.Args[2:], os.Getpid())

	//args[2] is the main command, args[3:] are the arguments
	cmd := exec.Command(os.Args[2], os.Args[3:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	checkErr(cmd.Run())
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
