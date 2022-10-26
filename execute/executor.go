package execute

import (
	"ebash/cmd-executor/config"
)

type Executer interface {
	Execute(string) *CommandOutput
}

type CommandOutput struct {
	Command string
	Stdout  string
	Stderr  string
	Error   error
}

var singletonExecutor Executer

func ProvideExecutor() Executer {
	if singletonExecutor != nil {
		return singletonExecutor
	}
	switch runOn := config.Vms().RunOn; runOn {
	case "native":
		singletonExecutor = NewNativeExecutor()
	case "vagrant":
		singletonExecutor = NewVagrantExecutor()
	case "vagrant-ssh":
		singletonExecutor = NewVagrantSSHExecutor()
	case "docker":
		// TODO
		panic("NOT IMPLEMENTED: DOCKER EXECUTOR")
	default:
		panic("invalid option of vms.runOn: " + runOn)
	}
	return singletonExecutor
}
