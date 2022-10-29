package abstract

type Manager interface {
	BringUpMachines()
	Shutdown(chan<- bool)
}

type Executor interface {
	Execute(string) *CommandOutput
}

type CommandOutput struct {
	Command string
	Stdout  string
	Stderr  string
	Error   error
}
