package persistant

import (
	dto "ebash/cmd-executor/communication"
	exe "ebash/cmd-executor/execute"
	"log"
)

func PersistCommand(output *exe.CommandOutput) {
	// TODO: add real persisting
	log.Printf("Received command to execute: [%v]", output.Command)
	log.Printf("Stdout:\n%v", output.Stdout)
	log.Printf("Stderr:\n%v", output.Stderr)
	log.Printf("Error:\n%v", dto.NillabeError(output.Error))
}
