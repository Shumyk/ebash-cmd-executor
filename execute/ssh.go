package execute

import (
	"fmt"
	"log"
	"os"

	"github.com/bmatcuk/go-vagrant"
	"golang.org/x/crypto/ssh"
)

func (v *AliveVagrant) InitSSHSession() {
	sshConfig := v.SshConfig()

	privateKey, err := os.ReadFile(sshConfig.IdentityFile)
	logPanically(err, "read vagrant indentity file")
	signer, err := ssh.ParsePrivateKey(privateKey)
	logPanically(err, "parse vagrant private key")

	clientConfig := &ssh.ClientConfig{
		User:            sshConfig.User,
		Auth:            []ssh.AuthMethod{ssh.PublicKeys(signer)},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	client, err := ssh.Dial("tcp", buildAddr(sshConfig), clientConfig)
	logPanically(err, "ssh dial vagrant")
	session, err := client.NewSession()
	logPanically(err, "create new vagrant session")
	// session.Stdout, session.Stderr = &bytes.Buffer{}, &bytes.Buffer{}

	log.Print("initialized vagrant ssh session")
	v.Session = session
}

func buildAddr(c *vagrant.SSHConfig) string {
	return fmt.Sprintf("%v:%v", c.HostName, c.Port)
}
