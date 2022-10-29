package common

type CommandOutput struct {
	Command string
	Stdout  string
	Stderr  string
	Error   error
}
