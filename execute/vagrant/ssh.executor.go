package vagrant

import (
	"ebash/cmd-executor/execute/common"
	"fmt"
)

type SSHExecutor struct{ *Manager }

func (e *SSHExecutor) Execute(command string) *common.CommandOutput {
	sshCommand := fmt.Sprintf(
		"(cd %v; vagrant ssh -c \"%v\")",
		e.pool.vagrant().VagrantClient.VagrantfileDir,
		command,
	)
	return common.Bash(sshCommand)
}
