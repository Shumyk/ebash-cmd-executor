package execute

type NativeExecutor struct{}

func (e *NativeExecutor) Execute(command string) *CommandOutput {
	return bash(command)
}

func NewNativeExecutor() Executor {
	return new(NativeExecutor)
}
