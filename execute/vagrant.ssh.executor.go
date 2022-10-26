package execute

import (
	"fmt"
)

type VagrantSSHExecutor struct{}

func (e *VagrantSSHExecutor) Execute(command string) *CommandOutput {
	sshCommand := fmt.Sprintf(
		"(cd %v; vagrant ssh -c \"%v\")",
		vagrants[0].VagrantClient.VagrantfileDir,
		command,
	)
	return bash(sshCommand)
}

func NewVagrantSSHExecutor() Executer {
	return new(VagrantSSHExecutor)
}
