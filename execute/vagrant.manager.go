package execute

import (
	"ebash/cmd-executor/config"
	"ebash/cmd-executor/util"
	"log"
	"sync"

	"github.com/bmatcuk/go-vagrant"
)

// TODO: 3. vagrant add boxes (?)
// TODO: 4. test async vagrant halt with multiple instances

// TODO: 5. probably move vagrant slice to pool object
// TODO: 5. implement VMs pool
// TODO: 7. VMs pool functionality:
//
//	a. creating new VMs
//	b. self-healing
//	c. concurrent access
type VagrantManager struct {
	vagrants []*AliveVagrant
}

func NewVagrantManager() *VagrantManager {
	return new(VagrantManager)
}

func (vm *VagrantManager) BringUpMachines() {
	for _, path := range config.Vagrant().Vagrantfiles {
		go initClient(vm, path)
	}
}

func initClient(vm *VagrantManager, path string) {
	aliveVagrant := &AliveVagrant{VagrantClient: newVagrantClient(path)}
	vm.vagrants = append(vm.vagrants, aliveVagrant)

	aliveVagrant.Up()
	aliveVagrant.initSSHClient(aliveVagrant.SSHConfig())
	// aliveVagrant.Status()
}

func newVagrantClient(path string) *vagrant.VagrantClient {
	return util.Cautiosly(vagrant.NewVagrantClient(path))("vagrant create client")
}

func (vm *VagrantManager) Shutdown(ch chan<- bool) {
	if !config.Vagrant().Halt {
		log.Println("vagrant halt is disabled")
		ch <- false
		return
	}

	wg := new(sync.WaitGroup)
	wg.Add(len(vm.vagrants))

	for _, v := range vm.vagrants {
		go v.DefinitelyHalt(wg)
	}

	wg.Wait()
	ch <- true
	log.Println("successfully halted all vagrants")
}
