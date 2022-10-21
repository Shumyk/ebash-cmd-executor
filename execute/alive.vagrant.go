package execute

import (
	"ebash/cmd-executor/config"
	"log"
	"sync"

	"github.com/bmatcuk/go-vagrant"
)

type AliveVagrant struct {
	*vagrant.VagrantClient
}

func (v *AliveVagrant) Up() {
	up := v.VagrantClient.Up()
	up.Verbose = config.GetVagrant().Verbose
	logPanically(up.Run(), "up")
}

func (v *AliveVagrant) Status() {
	status := v.VagrantClient.Status()
	status.Verbose = config.GetVagrant().Verbose
	logPanically(status.Run(), "status")

	log.Printf("vagrant status: %v", status.StatusResponse.Status["default"])
	if status.StatusResponse.Error != nil {
		log.Println("vagrant status error:")
		log.Println(status.StatusResponse.Error)
	}
}

func (v *AliveVagrant) DefinitelyHalt(wg *sync.WaitGroup) {
	if err := v.Halt(); err != nil {
		log.Printf("coudn't halt vagrant %v", v.VagrantClient.VagrantfileDir)
		forceErr := v.ForceHalt()
		logPanically(forceErr, "force halt")
	}
	wg.Done()
}

func (v *AliveVagrant) Halt() error {
	log.Printf("halting vagrant [%v]", v.VagrantClient.VagrantfileDir)
	halt := v.VagrantClient.Halt()
	halt.Verbose = config.GetVagrant().Verbose
	return halt.Run()
}

func (v *AliveVagrant) ForceHalt() error {
	log.Printf("force halting vagrant %v", v.VagrantClient.VagrantfileDir)
	forceHalt := v.VagrantClient.Halt()
	forceHalt.Verbose = config.GetVagrant().Verbose
	forceHalt.Force = true
	return forceHalt.Run()
}
