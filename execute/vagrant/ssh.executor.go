package vagrant

import (
	"ebash/cmd-executor/execute/abstract"
	"ebash/cmd-executor/execute/common"
	"fmt"
)

type SSHExecutor struct{ *Manager }

func (e *SSHExecutor) Execute(command string) *abstract.CommandOutput {
	sshCommand := fmt.Sprintf(
		"(cd %v; vagrant ssh -c \"%v\")",
		e.vagrants[0].VagrantClient.VagrantfileDir,
		command,
	)
	return common.Bash(sshCommand)
}
