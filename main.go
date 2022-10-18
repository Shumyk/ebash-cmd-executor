package main

import (
	"fmt"
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

type ExecuteResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func SuccessExecuteResponse(message string) *ExecuteResponse {
	return &ExecuteResponse{"success", message}
}
func FailedExecuteResponse(message string) *ExecuteResponse {
	return &ExecuteResponse{"failed", message}
}

func executePOST(context *gin.Context) {
	request := ExecuteRequest{}
	if error := context.BindJSON(&request); error != nil {
		context.AbortWithStatusJSON(
			http.StatusBadRequest,
			FailedExecuteResponse(fmt.Sprintf("Can't parse a body: %v", error)),
		)
		return
	}

	go PersistCommand(request.Command)
	commandOutput := ExecuteCommand(request.Command)
	context.JSON(http.StatusOK, SuccessExecuteResponse(commandOutput))
}

func PersistCommand(command string) {
	// TODO: add real persisting
	// TODO: move to appropriate package
	fmt.Printf("Received command to execute: [%v]\n", command)
}

func ExecuteCommand(command string) string {
	// TODO: implement executing
	// TODO: move to appropriate package
	return fmt.Sprintf("Processed result of: [%v]", command)
}

func main() {
	router := setupRouter()
	router.Run(":8080")
}
