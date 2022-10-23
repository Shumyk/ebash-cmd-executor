package execute

import (
	"ebash/cmd-executor/config"
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
	aliveVagrant := &AliveVagrant{VagrantClient: newClient(path)}
	vagrants = append(vagrants, aliveVagrant)

	aliveVagrant.Up()
	aliveVagrant.preInitSSHSession(aliveVagrant.SshConfig())
	aliveVagrant.InitSSHSession()
	aliveVagrant.Status()
}

func newClient(path string) *vagrant.VagrantClient {
	client, err := vagrant.NewVagrantClient(path)
	logPanically(err, "vagrant create client")
	return client
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

func logPanically(err error, action string) {
	if err != nil {
		log.Panicf("could not %v ahh!! [%v]", action, err)
	}
}
