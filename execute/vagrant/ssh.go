package vagrant

import (
	"ebash/cmd-executor/config"
	"ebash/cmd-executor/execute/common"
	"ebash/cmd-executor/util"
	"log"

	"github.com/bmatcuk/go-vagrant"
	"golang.org/x/crypto/ssh"
)

func (v *AliveVagrant) SSHConfig() *vagrant.SSHConfig {
	defer util.Timer("gathering vagrant ssh config")()

	sshConfig := v.VagrantClient.SSHConfig()
	sshConfig.Verbose = config.Vagrant().Verbose
	util.Panically(sshConfig.Run(), "vagrant ssh config")

	sshConfigs := sshConfig.SSHConfigResponse.Configs["default"]
	if config.Vagrant().Verbose {
		log.Printf("SSH config [%v]", v.VagrantClient.VagrantfileDir)
		log.Printf("Host			: %v", sshConfigs.HostName)
		log.Printf("Port			: %v", sshConfigs.Port)
		log.Printf("User			: %v", sshConfigs.User)
		log.Printf("Identity file		: %v", sshConfigs.IdentityFile)
	}
	return &sshConfigs
}

func (v *AliveVagrant) initSSHClient(c *vagrant.SSHConfig) {
	v.Client = common.CreateSSHClient(c.IdentityFile, c.User, c.HostName, c.Port)
}

func (v *AliveVagrant) Session() *ssh.Session {
	return util.Cautiosly(v.Client.NewSession())("create new vagrant session")
}
