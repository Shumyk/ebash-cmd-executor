package vagrant

import (
	"ebash/cmd-executor/config"
	"ebash/cmd-executor/util"
	"log"
	"sync"

	"github.com/bmatcuk/go-vagrant"
)

// TODO: 3. vagrant add boxes (?)
// TODO: 4. test async vagrant halt with multiple instances

// Manager TODO: 7. VMs pool functionality:
//
//	a. creating new VMs
//	b. self-healing
//	c. concurrent access
type Manager struct{ pool *pool }

func NewManager() *Manager {
	return &Manager{new(pool)}
}

func (vm *Manager) BringUpMachines() {
	for _, path := range config.Vagrant().Vagrantfiles {
		go initClient(vm, path)
	}
}

func initClient(vm *Manager, path string) {
	aliveVagrant := &aliveVagrant{VagrantClient: newVagrantClient(path)}
	vm.pool.add(aliveVagrant)

	aliveVagrant.up()
	aliveVagrant.initSSHClient(aliveVagrant.sshConfig())
	// aliveVagrant.status()
}

func newVagrantClient(path string) *vagrant.VagrantClient {
	return util.Cautiosly(vagrant.NewVagrantClient(path))("vagrant create client")
}

func (vm *Manager) Shutdown(ch chan<- bool) {
	if !config.Vagrant().Halt {
		log.Println("vagrant halt is disabled")
		ch <- false
		return
	}

	wg := new(sync.WaitGroup)
	wg.Add(vm.pool.size())

	for _, v := range vm.pool.elements() {
		go v.definitelyHalt(wg)
	}

	wg.Wait()
	ch <- true
	log.Println("successfully halted all vagrants")
}
