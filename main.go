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

type ExecuteRequest struct {
	Command string `json:"command" binding:"required"`
}

const (
	FAILED  = "failed"
	SUCCESS = "success"
)

type ExecuteResponse struct {
	Status string `json:"status"`
	Stdout string `json:"stdout"`
	Stderr string `json:"stderr"`
	Error  string `json:"error"`
}

func SuccessExecuteResponse(stdout, stderr string) *ExecuteResponse {
	return &ExecuteResponse{SUCCESS, stdout, stderr, ""}
}
func FailedExecuteBadRequest(context *gin.Context, err string) {
	context.AbortWithStatusJSON(
		http.StatusBadRequest,
		&ExecuteResponse{Status: FAILED, Error: err},
	)
}
func FailedExecuteInternalError(context *gin.Context, stdout, stderr, err string) {
	context.AbortWithStatusJSON(
		http.StatusInternalServerError,
		&ExecuteResponse{FAILED, stdout, stderr, err},
	)
}

func executePOST(context *gin.Context) {
	request := ExecuteRequest{}
	if err := context.BindJSON(&request); err != nil {
		FailedExecuteBadRequest(context, fmt.Sprintf("Can't parse a body: %v", err))
		return
	}

	go PersistCommand(request.Command)

	stdout, stderr, err := execute.ExecuteCommand(request.Command)
	if err != nil {
		FailedExecuteInternalError(context, stdout, stderr, err.Error())
		return
	}
	context.JSON(http.StatusOK, SuccessExecuteResponse(stdout, stderr))
}

func PersistCommand(command string) {
	// TODO: add real persisting
	// TODO: move to appropriate package
	log.Printf("Received command to execute: [%v]\n", command)
}

func main() {
	router := setupRouter()
	router.Run(":8080")
}
