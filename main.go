package main

import (
	"context"
	dto "ebash/cmd-executor/communication"
	exe "ebash/cmd-executor/execute"
	persistant "ebash/cmd-executor/persistance"
	"fmt"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	go exe.VagrantUp()

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	router := setupRouter()
	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	go listenAndServe(server)
	<-ctx.Done()

	vagrantChannel := make(chan bool)
	go exe.CleanUp(vagrantChannel)
	stop()
	log.Println("shutting down gracefully, press ctrl+c again to kill me ;(")

	<-vagrantChannel
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalln("server forced to shutdown: ", err)
	}

	log.Println("adios!")
}

func listenAndServe(server *http.Server) {
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("failed to serve: [%v]", err)
	}
}

func setupRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/ping", func(contex *gin.Context) {
		contex.String(http.StatusOK, "pong")
	})

	router.POST("/execute", executePOST)

	return router
}

func executePOST(context *gin.Context) {
	request := dto.ExecuteRequest{}
	if err := context.BindJSON(&request); err != nil {
		dto.FailedExecuteBadRequest(context, fmt.Sprintf("Can't parse a body: %v", err))
		return
	}

	stdout, stderr, err := exe.ExecuteCommand(request.Command)
	go persistant.PersistCommand(request.Command, stdout, stderr, err)

	context.JSON(http.StatusOK, dto.SuccessExecuteResponse(stdout, stderr, err))
}
