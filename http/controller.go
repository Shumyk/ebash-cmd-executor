package http

import (
	dto "ebash/cmd-executor/communication"
	"ebash/cmd-executor/execute"
	persistant "ebash/cmd-executor/persistance"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

var executor execute.Executer = execute.ProvideExecutor()

func pingGET(context *gin.Context) {
	context.String(http.StatusOK, "pong")
}

func executePOST(context *gin.Context) {
	request := dto.ExecuteRequest{}
	if err := context.BindJSON(&request); err != nil {
		dto.FailedExecuteBadRequest(context, fmt.Sprintf("Can't parse a body: %v", err))
		return
	}

	output := executor.Execute(request.Command)
	go persistant.PersistCommand(output)

	context.JSON(http.StatusOK, dto.SuccessExecuteResponse(output))
}
