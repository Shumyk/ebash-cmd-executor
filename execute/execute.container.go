package execute

import (
	"ebash/cmd-executor/config"
	"ebash/cmd-executor/execute/abstract"
	"ebash/cmd-executor/execute/common"
	"ebash/cmd-executor/execute/native"
	"ebash/cmd-executor/execute/vagrant"
)

var manager abstract.Manager
var executor abstract.Executor

func ProvideExecutor() abstract.Executor {
	initialize()
	return executor
}

func ProvideManager() abstract.Manager {
	initialize()
	return manager
}

func initialize() {
	if executor != nil {
		return
	}

	switch runOn := config.Vms().RunOn; runOn {
	case "native":
		manager, executor = new(common.VoidManager), new(native.Executor)
	case "vagrant":
		vagrantManager := new(vagrant.Manager)
		manager, executor = vagrantManager, &vagrant.Executor{Manager: vagrantManager}
	case "vagrant-ssh":
		vagrantManager := new(vagrant.Manager)
		manager, executor = vagrantManager, &vagrant.SSHExecutor{Manager: vagrantManager}
	case "docker":
		// TODO
		panic("NOT IMPLEMENTED: DOCKER EXECUTOR")
	default:
		panic("invalid option of vms.runOn: " + runOn)
	}
}
