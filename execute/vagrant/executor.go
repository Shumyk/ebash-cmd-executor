package vagrant

import (
	"bytes"
	"ebash/cmd-executor/execute/common"

	"golang.org/x/crypto/ssh"
)

type Executor struct{ *Manager }

func (e *Executor) Execute(command string) *common.CommandOutput {
	session := e.pool.vagrant().session()
	defer session.Close()

	stdout, stderr, err := runCommand(session, command)
	return &common.CommandOutput{
		Command: command,
		Stdout:  stdout,
		Stderr:  stderr,
		Error:   err,
	}
}

func runCommand(session *ssh.Session, command string) (stdout, stderr string, err error) {
	stdoutBuf, stderrBuf := bindOutputs(session)
	err = session.Run(command)
	return stdoutBuf.String(), stderrBuf.String(), err
}

func bindOutputs(session *ssh.Session) (stdout, stderr *bytes.Buffer) {
	stdout, stderr = &bytes.Buffer{}, &bytes.Buffer{}
	session.Stdout, session.Stderr = stdout, stderr
	return
}
