package execute

import (
	"bytes"
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
	cmd, stdout, stderr := prepareCommand(command)
	log.Printf("Running [%v] command in Bash\n", command)
	err := cmd.Run()
	log.Printf("Finnished executing of [%v] command", command)
	return stdout.String(), stderr.String(), err
}

func prepareCommand(command string) (*exec.Cmd, *bytes.Buffer, *bytes.Buffer) {
	stdout := &bytes.Buffer{}
	stderr := &bytes.Buffer{}
	cmd := exec.Command(BASH, COMMAND_FLAG, command)
	cmd.Stdout = stdout
	cmd.Stderr = stderr
	return cmd, stdout, stderr
}
