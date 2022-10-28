package execute

import (
	"ebash/cmd-executor/config"
	"log"
)

type ExecutorManager interface {
	BringUpMachines()
	Shutdown(chan<- bool)
}

var executorManager ExecutorManager

func InitializeExecutors() ExecutorManager {
	if executorManager != nil {
		log.Println("yeah singleton is working")
		return executorManager
	}

	switch runOn := config.Vms().RunOn; runOn {
	case "native":
		singletonExecutor = NewNativeExecutor()
	case "vagrant":
		vagrantManager := NewVagrantManager()
		executorManager, singletonExecutor = vagrantManager, NewVagrantExecutor(vagrantManager)
	case "vagrant-ssh":
		vagrantManager := NewVagrantManager()
		executorManager, singletonExecutor = vagrantManager, NewVagrantSSHExecutor(vagrantManager)
	case "docker":
		// TODO
		panic("NOT IMPLEMENTED: DOCKER EXECUTOR")
	default:
		panic("invalid option of vms.runOn: " + runOn)
	}
	return executorManager
}
