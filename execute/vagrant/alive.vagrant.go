package vagrant

import (
	"ebash/cmd-executor/config"
	"ebash/cmd-executor/util"
	"log"
	"sync"

	"github.com/bmatcuk/go-vagrant"
	"golang.org/x/crypto/ssh"
)

type AliveVagrant struct {
	*vagrant.VagrantClient
	*ssh.Client
}

func (v *AliveVagrant) Up() {
	up := v.VagrantClient.Up()
	up.Verbose = config.Vagrant().Verbose
	util.Panically(up.Run(), "vagrant up")
}

func (v *AliveVagrant) Status() {
	defer util.Timer("vagrant status")()

	status := v.VagrantClient.Status()
	status.Verbose = config.Vagrant().Verbose
	util.Panically(status.Run(), "vagrant status")

	log.Printf("vagrant status: %v", status.StatusResponse.Status["default"])
	if status.StatusResponse.Error != nil {
		log.Println("vagrant status error:")
		log.Println(status.StatusResponse.Error)
	}
}

func (v *AliveVagrant) DefinitelyHalt(wg *sync.WaitGroup) {
	if err := v.Halt(); err != nil {
		log.Printf("coudn't halt vagrant %v", v.VagrantClient.VagrantfileDir)
		util.Panically(v.ForceHalt(), "vagrant force halt")
	}
	wg.Done()
}

func (v *AliveVagrant) Halt() error {
	log.Printf("halting vagrant [%v]", v.VagrantClient.VagrantfileDir)
	halt := v.VagrantClient.Halt()
	halt.Verbose = config.Vagrant().Verbose
	return halt.Run()
}

// ForceHalt TODO: maybe creating vagrant command could be abstracted with closure to simplify verbose set
func (v *AliveVagrant) ForceHalt() error {
	log.Printf("force halting vagrant %v", v.VagrantClient.VagrantfileDir)
	forceHalt := v.VagrantClient.Halt()
	forceHalt.Verbose = config.Vagrant().Verbose
	forceHalt.Force = true
	return forceHalt.Run()
}
