package main

import (
	"ebash/cmd-executor/execute"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/ping", func(contex *gin.Context) {
		contex.String(http.StatusOK, "pong")
	})

	router.POST("/execute", executePOST)

	return router
}

func executePOST(context *gin.Context) {
	request := ExecuteRequest{}
	if err := context.BindJSON(&request); err != nil {
		failedExecuteBadRequest(context, fmt.Sprintf("Can't parse a body: %v", err))
		return
	}

	stdout, stderr, err := execute.ExecuteCommand(request.Command)
	go PersistCommand(request.Command, stdout, stderr, err)
	context.JSON(http.StatusOK, successExecuteResponse(stdout, stderr, errorDefault(err)))
}

func PersistCommand(command, stdout, stderr string, err error) {
	// TODO: add real persisting
	// TODO: move to appropriate package
	log.Printf("Received command to execute: [%v]", command)
	log.Printf("Stdout:\n%v", stdout)
	log.Printf("Stderr:\n%v", stderr)
	log.Printf("Error:\n%v", errorDefault(err))
}

func main() {
	router := setupRouter()
	router.Run(":8080")
}
