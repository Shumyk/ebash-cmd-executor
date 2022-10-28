package execute

type Executor interface {
	Execute(string) *CommandOutput
}

type CommandOutput struct {
	Command string
	Stdout  string
	Stderr  string
	Error   error
}

var singletonExecutor Executor

func ProvideExecutor() Executor {
	if singletonExecutor == nil {
		InitializeExecutors()
	}
	return singletonExecutor
}
