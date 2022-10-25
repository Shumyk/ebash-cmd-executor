package execute

import (
	"ebash/cmd-executor/util"
	"fmt"
	"os"

	"golang.org/x/crypto/ssh"
)

func CreateSSHClient(identityFilePath, user, hostname string, port int) *ssh.Client {
	defer util.Timer("Create SSH Client")()

	address := BuildAddress(hostname, port)
	clientConf := BuildClientConfig(identityFilePath, user)
	return util.Cautiosly(ssh.Dial("tcp", address, clientConf))("ssh dial vagrant")
}

func BuildClientConfig(identityFilePath, user string) *ssh.ClientConfig {
	privateKey := util.Cautiosly(os.ReadFile(identityFilePath))("read ssh indentity file")
	signer := util.Cautiosly(ssh.ParsePrivateKey(privateKey))("parse ssh private key")
	return &ssh.ClientConfig{
		User:            user,
		Auth:            []ssh.AuthMethod{ssh.PublicKeys(signer)},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
}

func BuildAddress(hostname string, port int) string {
	return fmt.Sprintf("%v:%v", hostname, port)
}
