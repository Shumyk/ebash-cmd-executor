package execute

import (
	"ebash/cmd-executor/config"
	"ebash/cmd-executor/util"
	"fmt"
	"log"
	"os"

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

func (v *AliveVagrant) initSSHClient(sshConfig *vagrant.SSHConfig) {
	defer util.Timer("preInitSSHSession")()

	privateKey := util.Cautiosly(os.ReadFile(sshConfig.IdentityFile))("read vagrant indentity file")
	signer := util.Cautiosly(ssh.ParsePrivateKey(privateKey))("parse vagrant private key")

	clientConf := &ssh.ClientConfig{
		User:            sshConfig.User,
		Auth:            []ssh.AuthMethod{ssh.PublicKeys(signer)},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	v.Client = util.Cautiosly(ssh.Dial("tcp", buildAddr(sshConfig), clientConf))("ssh dial vagrant")
}

func (v *AliveVagrant) initSSHSession() {
	defer util.Timer("initializing vagrant ssh session")()
	v.Session = util.Cautiosly(v.Client.NewSession())("create new vagrant session")
}

func (v *AliveVagrant) reinitSSHSession() {
	defer util.Timer("ReinitSSHSession")()
	v.Session.Close()
	v.initSSHSession()
}

func buildAddr(c *vagrant.SSHConfig) string {
	return fmt.Sprintf("%v:%v", c.HostName, c.Port)
}
