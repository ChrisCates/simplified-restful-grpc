package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"sync"
)

func main() {
	grpcRoot, e1 := os.LookupEnv("GRPC_ROOT")

	if e1 == true {
		pid1 := func() { runCmd("go", []string{"run", grpcRoot + "/src/grpc_server/grpc.server.go"}) }
		pid2 := func() { runCmd("go", []string{"run", grpcRoot + "/src/rest_server/rest.server.go"}) }
		fmt.Println("Running both servers parallel... Wait for messages to appear that they've started up.")
		parallelize(pid1, pid2)
	} else {
		fmt.Println("Error getting path variables...")
	}
}

func parallelize(functions ...func()) {
	var waitGroup sync.WaitGroup
	waitGroup.Add(len(functions))

	defer waitGroup.Wait()

	for _, function := range functions {
		go func(copy func()) {
			defer waitGroup.Done()
			copy()
		}(function)
	}
}

func runCmd(cmd string, args []string) {

	command := exec.Command(cmd, args...)

	stdout, _ := command.StdoutPipe()

	command.Start()

	for {
		r := bufio.NewReader(stdout)
		line, _, _ := r.ReadLine()
		fmt.Println(string(line))
	}

}
