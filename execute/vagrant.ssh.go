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

func (v *AliveVagrant) initSSHSessions() {
	defer util.Timer("initializing vagrant ssh sessions")()

	v.Sessions = new(util.Queue[*ssh.Session])
	for v.Sessions.Size() < config.Vms().SessionPoolSize {
		v.appendNewSession()
	}
	log.Printf("initiated %v SSH sessions", v.Sessions.Size())
}

func (v *AliveVagrant) Session() (session *ssh.Session, close func()) {
	go v.appendNewSession()
	session = v.Sessions.Poll()
	close = func() { go session.Close() }
	return
}

func (v *AliveVagrant) appendNewSession() {
	log.Printf("appending new session")
	v.Sessions.Add(v.newSession())
}

func (v *AliveVagrant) newSession() *ssh.Session {
	return util.Cautiosly(v.Client.NewSession())("create new vagrant session")
}

func buildAddr(c *vagrant.SSHConfig) string {
	return fmt.Sprintf("%v:%v", c.HostName, c.Port)
}
