package execute

import (
	"bytes"
)

type VagrantExecutor struct {
	manager *VagrantManager
}

func (e *VagrantExecutor) Execute(command string) *CommandOutput {
	v := e.manager.vagrants[0] // TODO: this should be changed when vm pool

	session := v.Session()
	defer session.Close()

	stdout, stderr := &bytes.Buffer{}, &bytes.Buffer{}
	session.Stdout, session.Stderr = stdout, stderr
	err := session.Run(command)

	return &CommandOutput{command, stdout.String(), stderr.String(), err}
}

func NewVagrantExecutor(manager *VagrantManager) *VagrantExecutor {
	return &VagrantExecutor{manager}
}
