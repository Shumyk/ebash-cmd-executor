package execute

import (
	"log"
	"sync"

	"github.com/bmatcuk/go-vagrant"
)

const (
	// TODO: path to Vagrant file from properties
	VAGRANT_PATH  = "/Users/shumyk/codeself/shell/cmd-exe"
	VAGRANTS_SIZE = 0
)

// TODO: 1. refactor struct from slice wrapper
// TODO: 2. vagrant up in loop for several machine
// TODO: 3. vagrant add boxes
// TODO: 4. test async vagrant halt with multiple instances
type AliveVagrant struct {
	*vagrant.VagrantClient
}

var vagrants []*AliveVagrant = make([]*AliveVagrant, VAGRANTS_SIZE)

func VagrantUp() {
	client, err := vagrant.NewVagrantClient(VAGRANT_PATH)
	logPanically(err, "create client")
	vagrants = append(vagrants, &AliveVagrant{client})

	upcmd := client.Up()
	upcmd.Verbose = true
	logPanically(upcmd.Run(), "up")

	go VagrantStatus(client.Status())
}

func VagrantStatus(status *vagrant.StatusCommand) {
	logPanically(status.Run(), "status")

	log.Printf("vagrant status: %v", status.StatusResponse.Status["default"])
	if status.StatusResponse.Error != nil {
		log.Println("vagrant status error:")
		log.Println(status.StatusResponse.Error)
	}
}

func HaltVagrants(ch chan<- bool) {
	wg := new(sync.WaitGroup)
	wg.Add(len(vagrants))

	for _, v := range vagrants {
		go v.DefinitelyHaltVagrant(wg)
	}

	wg.Wait()
	ch <- true
	log.Println("successfully halt all vagrants")
}

func (v *AliveVagrant) DefinitelyHaltVagrant(wg *sync.WaitGroup) {
	log.Printf("starting vagrant halt [%v]", v.VagrantClient.VagrantfileDir)
	halt := v.VagrantClient.Halt()
	halt.Verbose = true

	if err := halt.Run(); err != nil {
		log.Printf("coudn't halt vagrant %v", v.VagrantClient.VagrantfileDir)
		log.Printf("forcing vagrant halt %v", v.VagrantClient.VagrantfileDir)
		v.forceHaltVagrant()
	}

	wg.Done()
}

func (v *AliveVagrant) forceHaltVagrant() {
	forceHalt := v.VagrantClient.Halt()
	forceHalt.Verbose = true
	forceHalt.Force = true
	logPanically(forceHalt.Run(), "force halt")
}

func logPanically(err error, action string) {
	if err != nil {
		log.Panicf("could not vagrant %v ahh!! [%v]", action, err)
	}
}
