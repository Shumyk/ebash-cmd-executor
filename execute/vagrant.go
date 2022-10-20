package execute

import (
	"log"

	"github.com/bmatcuk/go-vagrant"
)

type AliveVagrants struct {
	*vagrant.VagrantClient
}

var vagrants []AliveVagrants = make([]AliveVagrants, 0)

func VagrantUp() {
	// TODO: path to Vagrant file from properties
	client, err := vagrant.NewVagrantClient("/Users/shumyk/codeself/shell/cmd-exe")
	if err != nil {
		log.Panicf("could not vagrant client ahh!! [%v]", err)
	}
	vagrants = append(vagrants, AliveVagrants{client})

	upcmd := client.Up()
	upcmd.Verbose = true
	if err := upcmd.Run(); err != nil {
		log.Panicf("could not vagrant up ahh!! [%v]", err)
	}

	status := client.Status()
	status.Verbose = true
	if err := status.Run(); err != nil {
		log.Panicf("could not vagrant status ahh!! [%v]", err)
	}

	log.Printf("vagrant status: [%v]", status.MachineName)
	log.Println(status.StatusResponse.Status)
	log.Println("vagrant status error:")
	log.Println(status.StatusResponse.Error)
	log.Println("vagrant status response:")
	log.Println(status.StatusResponse.ErrorResponse)
}

func CleanUp(ch chan bool) {
	for _, wr := range vagrants {
		log.Printf("starting vagrant halt [%v]", wr.VagrantClient.VagrantfileDir)
		halt := wr.VagrantClient.Halt()
		halt.Verbose = true
		if err := halt.Run(); err != nil {
			log.Printf("coudn't halt vagrant [%v]", wr.VagrantClient.VagrantfileDir)
			log.Printf("forcing vagrant halt")

			forceHalt := wr.VagrantClient.Halt()
			forceHalt.Verbose = true
			forceHalt.Force = true
			if err := forceHalt.Run(); err != nil {
				log.Printf("coudn't force halt vagrant [%v]", wr.VagrantClient.VagrantfileDir)
				log.Printf("destroying vagrant")

				destroy := wr.VagrantClient.Destroy()
				destroy.Verbose = true
				destroy.Run()
			}
		}
	}
	log.Println("successfully halt all vagrants")
	ch <- true
}
