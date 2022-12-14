package vagrant

import (
	"ebash/cmd-executor/config"
	"ebash/cmd-executor/util"
	"log"
	"sync"

	"github.com/bmatcuk/go-vagrant"
	"golang.org/x/crypto/ssh"
)

type aliveVagrant struct {
	*vagrant.VagrantClient
	*ssh.Client
}

func (v *aliveVagrant) up() {
	up := v.VagrantClient.Up()
	up.Verbose = config.Vagrant().Verbose
	util.Panically(up.Run(), "vagrant up")
}

func (v *aliveVagrant) status() {
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

func (v *aliveVagrant) definitelyHalt(wg *sync.WaitGroup) {
	if err := v.halt(); err != nil {
		log.Printf("coudn't halt vagrant %v", v.VagrantClient.VagrantfileDir)
		util.Panically(v.forceHalt(), "vagrant force halt")
	}
	wg.Done()
}

func (v *aliveVagrant) halt() error {
	log.Printf("halting vagrant [%v]", v.VagrantClient.VagrantfileDir)
	halt := v.VagrantClient.Halt()
	halt.Verbose = config.Vagrant().Verbose
	return halt.Run()
}

// forceHalt TODO: maybe creating vagrant command could be abstracted with closure to simplify verbose set
func (v *aliveVagrant) forceHalt() error {
	log.Printf("force halting vagrant %v", v.VagrantClient.VagrantfileDir)
	forceHalt := v.VagrantClient.Halt()
	forceHalt.Verbose = config.Vagrant().Verbose
	forceHalt.Force = true
	return forceHalt.Run()
}
