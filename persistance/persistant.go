package persistant

import (
	dto "ebash/cmd-executor/communication"
	"log"
)

func PersistCommand(command, stdout, stderr string, err error) {
	// TODO: add real persisting
	log.Printf("Received command to execute: [%v]", command)
	log.Printf("Stdout:\n%v", stdout)
	log.Printf("Stderr:\n%v", stderr)
	log.Printf("Error:\n%v", dto.ErrorDefault(err))
}
