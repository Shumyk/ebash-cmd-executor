package execute

import "ebash/cmd-executor/execute/common"

type Manager interface {
	BringUpMachines()
	Shutdown(chan<- bool)
}

type Executor interface {
	Execute(string) *common.CommandOutput
}
