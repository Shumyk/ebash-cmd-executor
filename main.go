package main

import (
	dto "ebash/cmd-executor/communication"
	exe "ebash/cmd-executor/execute"
	persistant "ebash/cmd-executor/persistance"
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

func executePOST(context *gin.Context) {
	request := dto.ExecuteRequest{}
	if err := context.BindJSON(&request); err != nil {
		dto.FailedExecuteBadRequest(context, fmt.Sprintf("Can't parse a body: %v", err))
		return
	}

	stdout, stderr, err := exe.ExecuteCommand(request.Command)
	go persistant.PersistCommand(request.Command, stdout, stderr, err)

	context.JSON(http.StatusOK, dto.SuccessExecuteResponse(stdout, stderr, dto.ErrorDefault(err)))
}

func main() {
	router := setupRouter()
	router.Run(":8080")
}
