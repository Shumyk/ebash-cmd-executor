package vagrant

import (
	"bytes"
	"ebash/cmd-executor/execute/common"
)

type Executor struct{ *Manager }

func (e *Executor) Execute(command string) *common.CommandOutput {
	v := e.vagrants[0] // TODO: this should be changed when vm pool

	session := v.Session()
	defer session.Close()

	stdout, stderr := &bytes.Buffer{}, &bytes.Buffer{}
	session.Stdout, session.Stderr = stdout, stderr
	err := session.Run(command)

	return &common.CommandOutput{Command: command, Stdout: stdout.String(), Stderr: stderr.String(), Error: err}
}
