package execute

import (
	"ebash/cmd-executor/config"
	"log"
	"sync"

	"github.com/bmatcuk/go-vagrant"
	"golang.org/x/crypto/ssh"
)

type AliveVagrant struct {
	*vagrant.VagrantClient
	*ssh.Session
}

func (v *AliveVagrant) Up() {
	up := v.VagrantClient.Up()
	up.Verbose = config.Vagrant().Verbose
	logPanically(up.Run(), "vagrant up")
}

func (v *AliveVagrant) Status() {
	status := v.VagrantClient.Status()
	status.Verbose = config.Vagrant().Verbose
	logPanically(status.Run(), "vagrant status")

	log.Printf("vagrant status: %v", status.StatusResponse.Status["default"])
	if status.StatusResponse.Error != nil {
		log.Println("vagrant status error:")
		log.Println(status.StatusResponse.Error)
	}
}

func (v *AliveVagrant) SshConfig() *vagrant.SSHConfig {
	sshConfig := v.VagrantClient.SSHConfig()
	sshConfig.Verbose = config.Vagrant().Verbose
	logPanically(sshConfig.Run(), "vagrant ssh config")

	configs := sshConfig.SSHConfigResponse.Configs["default"]
	log.Printf("SSH config [%v]", v.VagrantClient.VagrantfileDir)
	log.Printf("Host			: %v", configs.HostName)
	log.Printf("Port			: %v", configs.Port)
	log.Printf("User			: %v", configs.User)
	log.Printf("Identity file		: %v", configs.IdentityFile)
	return &configs
}

func (v *AliveVagrant) DefinitelyHalt(wg *sync.WaitGroup) {
	if err := v.Halt(); err != nil {
		log.Printf("coudn't halt vagrant %v", v.VagrantClient.VagrantfileDir)
		forceErr := v.ForceHalt()
		logPanically(forceErr, "vagrant force halt")
	}
	wg.Done()
}

func (v *AliveVagrant) Halt() error {
	log.Printf("halting vagrant [%v]", v.VagrantClient.VagrantfileDir)
	halt := v.VagrantClient.Halt()
	halt.Verbose = config.Vagrant().Verbose
	return halt.Run()
}

func (v *AliveVagrant) ForceHalt() error {
	log.Printf("force halting vagrant %v", v.VagrantClient.VagrantfileDir)
	forceHalt := v.VagrantClient.Halt()
	forceHalt.Verbose = config.Vagrant().Verbose
	forceHalt.Force = true
	return forceHalt.Run()
}
