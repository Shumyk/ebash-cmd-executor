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
var vagrants []*AliveVagrant

func VagrantsUp() {
	for _, path := range config.Vagrant().Vagrantfiles {
		go initClient(path)
	}
}

func initClient(path string) {
	aliveVagrant := &AliveVagrant{VagrantClient: newVagrantClient(path)}
	vagrants = append(vagrants, aliveVagrant)

	aliveVagrant.Up()
	aliveVagrant.initSSHClient(aliveVagrant.SSHConfig())
	aliveVagrant.initSSHSessions()
	// aliveVagrant.Status()
}

func newVagrantClient(path string) *vagrant.VagrantClient {
	return util.Cautiosly(vagrant.NewVagrantClient(path))("vagrant create client")
}

func HaltVagrants(ch chan<- bool) {
	if !config.Vagrant().Halt {
		log.Println("vagrant halt is disabled")
		ch <- false
		return
	}

	wg := new(sync.WaitGroup)
	wg.Add(len(vagrants))

	for _, v := range vagrants {
		go v.DefinitelyHalt(wg)
	}

	wg.Wait()
	ch <- true
	log.Println("successfully halted all vagrants")
}
