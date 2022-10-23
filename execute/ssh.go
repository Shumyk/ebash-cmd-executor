package execute

import (
	"ebash/cmd-executor/config"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/bmatcuk/go-vagrant"
	"golang.org/x/crypto/ssh"
)

func (v *AliveVagrant) SshConfig() *vagrant.SSHConfig {
	defer timer("gathering vagrant ssh config")()

	sshConfig := v.VagrantClient.SSHConfig()
	sshConfig.Verbose = config.Vagrant().Verbose
	logPanically(sshConfig.Run(), "vagrant ssh config")

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

func (v *AliveVagrant) preInitSSHSession(sshConfig *vagrant.SSHConfig) {
	defer timer("preInitSSHSession")()

	privateKey, err := os.ReadFile(sshConfig.IdentityFile)
	logPanically(err, "read vagrant indentity file")
	signer, err := ssh.ParsePrivateKey(privateKey)
	logPanically(err, "parse vagrant private key")

	clientConf := &ssh.ClientConfig{
		User:            sshConfig.User,
		Auth:            []ssh.AuthMethod{ssh.PublicKeys(signer)},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	v.Client = cautiosly(ssh.Dial("tcp", buildAddr(sshConfig), clientConf))("ssh dial vagrant")
}

func (v *AliveVagrant) InitSSHSession() {
	defer timer("initializing vagrant ssh session")()
	v.Session = cautiosly(v.Client.NewSession())("create new vagrant session")
}

func (v *AliveVagrant) ReinitSSHSession() {
	defer timer("ReinitSSHSession")()
	v.Session.Close()
	v.InitSSHSession()
}

// TODO: functions below should be moved to util package
func timer(name string) func() {
	start := time.Now()
	return func() {
		log.Printf("%v took %v", name, time.Since(start))
	}
}

func buildAddr(c *vagrant.SSHConfig) string {
	return fmt.Sprintf("%v:%v", c.HostName, c.Port)
}

// TODO: reuse it everywhere instead of logPanically
func cautiosly[T any](result T, err error) func(string) T {
	return func(action string) T {
		logPanically(err, action)
		return result
	}
}
