package task

// Task represents a single executable unit in TaskForge.
type Task struct {
	Name         string   // Name is the unique identifier of the task
	Command      string   // Command is the shell command to execute
	Dependencies []string // Dependencies is a list of task names that must run before this task
}
