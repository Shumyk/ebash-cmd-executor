package execute

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
)

const (
	BASH         = "bash"
	COMMAND_FLAG = "-c"
)

// TODO: 1. run this on vagrant VMs
// TODO: 2. figure out a way to keep open ssh connection to VM
// TODO: 3. implement VMs pool
// TODO: 4. VMs pool functionality:
//					a. creating new VMs
//					b. self-healing
//					c. concurrent access

func ExecuteCommand(command string) (string, string, error) {
	// return executeCommandHostMachine(command)
	return executeCommandOnVirtualMachine(command)
}

func executeCommandOnVirtualMachine(command string) (string, string, error) {
	return executeCommandHostMachine(prepareVagrantCommand(command))
}

func prepareVagrantCommand(command string) string {
	return fmt.Sprintf(
		"(cd %v; vagrant ssh -c \"%v\")",
		vagrants[0].VagrantClient.VagrantfileDir,
		command,
	)
}

func executeCommandHostMachine(command string) (string, string, error) {
	cmd, stdout, stderr := prepareCommand(command)

	log.Printf("Running [%v] command in Bash\n", command)
	err := cmd.Run()
	log.Printf("Finnished executing of [%v] command", command)

	// return CommandOutput{stdou.String(), stderr.String(), err}
	return stdout.String(), stderr.String(), err
}

func prepareCommand(command string) (*exec.Cmd, *bytes.Buffer, *bytes.Buffer) {
	stdout, stderr := &bytes.Buffer{}, &bytes.Buffer{}
	cmd := exec.Command(BASH, COMMAND_FLAG, command)
	cmd.Stdout, cmd.Stderr = stdout, stderr
	return cmd, stdout, stderr
}

type CommandOutput struct {
	stdout string
	stderr string
	err    error
}
