package task

import (
	"fmt"
	"log"
	"os/exec"
)

// Runner is responsible for executing tasks.
type Runner struct{}

// NewRunner creates a new task runner.
func NewRunner() *Runner {
	return &Runner{}
}

func (r *Runner) Run(t Task) error {
	log.Printf("Running task: %s\n", t.Name)
	cmd := exec.Command("sh", "-c", t.Command)
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("Error while executing command %s. Error: %s\n", t.Command, err)
		return err
	}
	fmt.Printf("%s", output)
	log.Printf("Task complete: %s\n", t.Name)
	return nil
}
