package native

import (
	"ebash/cmd-executor/execute/abstract"
	"ebash/cmd-executor/execute/common"
)

type Executor struct{}

func (e *Executor) Execute(command string) *abstract.CommandOutput {
	return common.Bash(command)
}
