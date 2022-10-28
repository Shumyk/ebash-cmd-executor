package execute

import (
	"fmt"
)

type VagrantSSHExecutor struct {
	manager *VagrantManager
}

func (e *VagrantSSHExecutor) Execute(command string) *CommandOutput {
	sshCommand := fmt.Sprintf(
		"(cd %v; vagrant ssh -c \"%v\")",
		e.manager.vagrants[0].VagrantClient.VagrantfileDir,
		command,
	)
	return bash(sshCommand)
}

func NewVagrantSSHExecutor(manager *VagrantManager) *VagrantSSHExecutor {
	return &VagrantSSHExecutor{manager}
}
