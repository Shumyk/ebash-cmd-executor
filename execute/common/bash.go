package common

import (
	"bytes"
	"ebash/cmd-executor/execute/abstract"
	"log"
	"os/exec"
)

const (
	BASH        = "bash"
	CommandFlag = "-c"
)

func Bash(command string) *abstract.CommandOutput {
	cmd, stdout, stderr := prepareCommand(command)

	log.Printf("Running [%v] command on local machine\n", command)
	err := cmd.Run()
	log.Printf("Finnished executing of [%v] command", command)

	return &abstract.CommandOutput{command, stdout.String(), stderr.String(), err}
}

func prepareCommand(command string) (*exec.Cmd, *bytes.Buffer, *bytes.Buffer) {
	stdout, stderr := &bytes.Buffer{}, &bytes.Buffer{}
	cmd := exec.Command(BASH, CommandFlag, command)
	cmd.Stdout, cmd.Stderr = stdout, stderr
	return cmd, stdout, stderr
}
