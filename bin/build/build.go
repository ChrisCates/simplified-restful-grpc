package main

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
)

type configStruct struct {
	displayCmds  bool
	displayArgs  bool
	extraLogging bool
}

var config = configStruct{true, true, true}
var count = 1

func main() {
	grpcRoot, e1 := os.LookupEnv("GRPC_ROOT")
	gopath, e2 := os.LookupEnv("GOPATH")

	if e1 == true && e2 == true {
		//Build handler for base.proto
		runCmd("protoc", []string{"-I", "/usr/local/include", "-I", grpcRoot + "/src", "-I", gopath + "/src", "-I", gopath + "/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis", "--go_out=plugins=grpc:./src", grpcRoot + "/src/grpc/api.proto"})
		//Build gateway for base.proto
		runCmd("protoc", []string{"-I", "/usr/local/include", "-I", grpcRoot + "/src", "-I", gopath + "/src", "-I", gopath + "/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis", "--grpc-gateway_out=logtostderr=true:./src", grpcRoot + "/src/grpc/api.proto"})
	} else {
		fmt.Println("Error getting path variables...")
	}
}

func runCmd(cmd string, args []string) {

	message := genMsgOk(cmd, args)
	out, err := exec.Command(cmd, args...).Output()

	if err != nil {
		message = genMsgBad(cmd, args)
	}

	fmt.Println(message)

	if config.extraLogging == true {
		fmt.Println("Here's the output:")
		fmt.Println(string(out))
		fmt.Println("")
		if err != nil {
			fmt.Println("Here's the errors:")
			fmt.Println(err)
			fmt.Println("")
			fmt.Println("*If you need to debug further, you should debug in a terminal session...")
		}
	}

	count = count + 1
}

func genMsgOk(cmd string, args []string) string {
	message := "✅ Ran command #" + strconv.Itoa(count) + " successfully\n"

	if config.displayCmds == true {
		message = formatCmd(message, cmd, args)
	}

	if config.displayArgs == true {
		message = formatArgs(message, args)
	}

	return message
}

func genMsgBad(cmd string, args []string) string {
	message := "🛑 Ran command #" + strconv.Itoa(count) + " with errors 💩\n"
	message = formatCmd(message, cmd, args)
	message = formatArgs(message, args)

	return message
}

func formatCmd(message string, cmd string, args []string) string {
	message = message + "\nRaw command:\n\n\t" + cmd
	for _, arg := range args {
		message = message + " " + arg
	}
	message = message + "\n"
	return message
}

func formatArgs(message string, args []string) string {
	message = message + "\nArguments used:\n\n"
	for i, arg := range args {
		message = message + "\targ" + strconv.Itoa(i) + ": " + arg + "\n"
	}
	message = message + "\n"
	return message
}
