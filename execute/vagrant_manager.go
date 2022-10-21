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

// TODO: 2. vagrant up in loop for several machine
// TODO: 3. vagrant add boxes
// TODO: 4. test async vagrant halt with multiple instances

// TODO: 5. probably move vagrant slice to pool object
var vagrants []*AliveVagrant = make([]*AliveVagrant, VAGRANTS_SIZE)

func VagrantsUp() {
	client, err := vagrant.NewVagrantClient(VAGRANT_PATH)
	logPanically(err, "create client")

	aliveVagrant := &AliveVagrant{client}
	vagrants = append(vagrants, aliveVagrant)

	aliveVagrant.Up()
	go aliveVagrant.Status()
}

func HaltVagrants(ch chan<- bool) {
	wg := new(sync.WaitGroup)
	wg.Add(len(vagrants))

	for _, v := range vagrants {
		go v.DefinitelyHalt(wg)
	}

	wg.Wait()
	ch <- true
	log.Println("successfully halt all vagrants")
}

func logPanically(err error, action string) {
	if err != nil {
		log.Panicf("could not vagrant %v ahh!! [%v]", action, err)
	}
}
