package main

import (
	"os/exec"
	"os"
	"github.com/kvz/logstreamer"
	"log"
	"sync"
)

var wg sync.WaitGroup

func createCommandWithPrefixedOutputs(prefix string, name string, args ...string) *exec.Cmd {
	proc := exec.Command(name, args...)

	stdOutLogger := log.New(os.Stdout, prefix + " ", log.Ldate|log.Ltime)
	stdErrLogger := log.New(os.Stdout, prefix + " ", log.Ldate|log.Ltime)

	proc.Stdout = logstreamer.NewLogstreamer(stdOutLogger, "stdout", true)
	proc.Stderr = logstreamer.NewLogstreamer(stdErrLogger, "stderr", true)

	return proc
}

func execCommand(cmd *exec.Cmd) {
	wg.Add(1)

	cmd.Start()
	cmd.Wait()

	wg.Done()
}


func main() {


	wg.Add(5)
	serverProc := createCommandWithPrefixedOutputs("[server]", "go", "run", "server/server.go", "--config", "server/config/server.json")
	jsonApiProc := createCommandWithPrefixedOutputs("[json_api]", "go", "run", "clients/json_api/json_api.go", "--config", "clients/config/json_api_client.json")
	loggerProc := createCommandWithPrefixedOutputs("[logger]", "go", "run", "clients/logger/logger.go", "--config", "clients/config/logger.json")
	mongoProc := createCommandWithPrefixedOutputs("[mongo]", "go", "run", "clients/mongo/mongo.go", "--config", "clients/config/mongo.json")
	eventStreamProc := createCommandWithPrefixedOutputs("[event_stream]", "go", "run", "clients/event_stream/event_stream.go", "--config", "clients/config/event_stream.json")

	go execCommand(serverProc)
	go execCommand(jsonApiProc)
	go execCommand(loggerProc)
	go execCommand(mongoProc)
	go execCommand(eventStreamProc)

	wg.Wait()
}
