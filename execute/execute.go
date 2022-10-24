package execute

import (
	"bytes"
	"ebash/cmd-executor/config"
	"errors"
	"fmt"
	"log"
	"os/exec"
)

const (
	BASH         = "bash"
	COMMAND_FLAG = "-c"
)

// TODO: 3. implement VMs pool
// TODO: 4. VMs pool functionality:
//					a. creating new VMs
//					b. self-healing
//					c. concurrent access
// TODO: 5. ssh pool
//			i. fix limit of 10 open ssh session in vagrant
//			ii. added check to create session when no available in ssh pool

func ExecuteCommand(command string) *CommandOutput {
	switch runOn := config.Vms().RunOn; runOn {
	case "native":
		return executeCommandHostMachine(command)
	case "vagrant":
		return executeCommandOnVirtualMachine(command)
	default:
		errMsg := fmt.Sprintf("unknown target machine in configuration: %v", runOn)
		return &CommandOutput{Error: errors.New(errMsg)}
	}
}

func executeCommandOnVirtualMachine(command string) *CommandOutput {
	v := vagrants[0] // TODO: this should be changed when vm/ssh pool

	session := v.Session()
	defer func() { go session.Close() }()

	stdout, stderr := &bytes.Buffer{}, &bytes.Buffer{}
	session.Stdout, session.Stderr = stdout, stderr
	err := session.Run(command)

	return &CommandOutput{command, stdout.String(), stderr.String(), err}
}

func prepareVagrantSshCommand(command string) string {
	return fmt.Sprintf(
		"(cd %v; vagrant ssh -c \"%v\")",
		vagrants[0].VagrantClient.VagrantfileDir,
		command,
	)
}

func executeCommandHostMachine(command string) *CommandOutput {
	cmd, stdout, stderr := prepareCommand(command)

	log.Printf("Running [%v] command in Bash\n", command)
	err := cmd.Run()
	log.Printf("Finnished executing of [%v] command", command)

	return &CommandOutput{command, stdout.String(), stderr.String(), err}
}

func prepareCommand(command string) (*exec.Cmd, *bytes.Buffer, *bytes.Buffer) {
	stdout, stderr := &bytes.Buffer{}, &bytes.Buffer{}
	cmd := exec.Command(BASH, COMMAND_FLAG, command)
	cmd.Stdout, cmd.Stderr = stdout, stderr
	return cmd, stdout, stderr
}

type CommandOutput struct {
	Command string
	Stdout  string
	Stderr  string
	Error   error
}
