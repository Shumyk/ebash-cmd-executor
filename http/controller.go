package http

import (
	dto "ebash/cmd-executor/communication"
	"ebash/cmd-executor/execute"
	"ebash/cmd-executor/execute/abstract"
	persistent "ebash/cmd-executor/persistance"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

var executor abstract.Executor

func setup() {
	executor = execute.ProvideExecutor()
}

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
	go persistent.PersistCommand(output)

	context.JSON(http.StatusOK, dto.SuccessExecuteResponse(output))
}
