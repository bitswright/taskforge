package main

import (
	"github.com/bitswright/taskforge/internals/task"
)

func main() {
	helloTask := task.Task{
		Name:    "hello",
		Command: "echo 'Hello World!'",
	}
	runner := task.NewRunner()
	runner.Run(helloTask)
}
