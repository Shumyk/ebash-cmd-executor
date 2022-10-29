package persistant

import (
	dto "ebash/cmd-executor/communication"
	"ebash/cmd-executor/config"
	exe "ebash/cmd-executor/execute/common"
	"log"
)

func PersistCommand(output *exe.CommandOutput) {
	// TODO: add real persisting
	if !config.Persistence().Enabled {
		return
	}
	log.Printf("Received command to persist: [%v]", output.Command)
	log.Printf("Stdout:\n%v", output.Stdout)
	log.Printf("Stderr:\n%v", output.Stderr)
	log.Printf("Error:\n%v", dto.NillabeError(output.Error))
}
