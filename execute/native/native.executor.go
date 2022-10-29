package native

import (
	"ebash/cmd-executor/execute/common"
)

type Executor struct{}

func (e *Executor) Execute(command string) *common.CommandOutput {
	return common.Bash(command)
}
