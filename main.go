package main

import (
	"context"
	exe "ebash/cmd-executor/execute"
	"ebash/cmd-executor/http"
	"log"
	"os/signal"
	"syscall"
)

func main() {
	go exe.VagrantsUp()

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	server := http.Server()
	<-ctx.Done()

	// halt vagrants
	haltVagrants := make(chan bool)
	go exe.HaltVagrants(haltVagrants)

	// stop context
	stop()
	log.Println("shutting down gracefully, press ctrl+c again to kill me ;(")
	http.ShutdownServer(server)
	<-haltVagrants
	log.Println("adios!")
}
