package execute

import (
	"log"

	"github.com/bmatcuk/go-vagrant"
)

func VagrantUp() {
	// TODO: path to Vagrant file from properties
	client, err := vagrant.NewVagrantClient("/Users/shumyk/codeself/shell/cmd-exe")
	if err != nil {
		log.Panicf("could not vagrant client ahh!! [%v]", err)
	}

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
