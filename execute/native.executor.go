package execute

type NativeExecutor struct{}

func (e *NativeExecutor) Execute(command string) *CommandOutput {
	return bash(command)
}

func NewNativeExecutor() Executer {
	return new(NativeExecutor)
}
