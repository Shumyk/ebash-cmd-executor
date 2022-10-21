package execute

import (
	"log"
	"strings"
	"sync"

	"github.com/bmatcuk/go-vagrant"
)

const (
	// TODO: path to Vagrant file from properties
	VAGRANT_PATHS  = "/Users/shumyk/codeself/shell/cmd-exe"
	PATH_SEPARATOR = ","
)

// TODO: 3. vagrant add boxes (?)
// TODO: 4. test async vagrant halt with multiple instances

// TODO: 5. probably move vagrant slice to pool object
var vagrants []*AliveVagrant

func VagrantsUp() {
	paths := strings.Split(VAGRANT_PATHS, PATH_SEPARATOR)
	for _, path := range paths {
		go initClient(path)
	}
}

func initClient(path string) {
	aliveVagrant := &AliveVagrant{newClient(path)}
	vagrants = append(vagrants, aliveVagrant)

	aliveVagrant.Up()
	aliveVagrant.Status()
}

func newClient(path string) *vagrant.VagrantClient {
	client, err := vagrant.NewVagrantClient(path)
	logPanically(err, "create client")
	return client
}

func HaltVagrants(ch chan<- bool) {
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
		log.Panicf("could not vagrant %v ahh!! [%v]", action, err)
	}
}
