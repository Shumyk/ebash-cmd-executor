package communication

import (
	"ebash/cmd-executor/execute/abstract"
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	FAILED  = "failed"
	SUCCESS = "success"
)

type ExecuteRequest struct {
	Command string `json:"command" binding:"required"`
}

type ExecuteResponse struct {
	Status string `json:"status"`
	Stdout string `json:"stdout"`
	Stderr string `json:"stderr"`
	Error  string `json:"error"`
}

func SuccessExecuteResponse(output *abstract.CommandOutput) *ExecuteResponse {
	return &ExecuteResponse{SUCCESS, output.Stdout, output.Stderr, NillabeError(output.Error)}
}

func FailedExecuteBadRequest(context *gin.Context, err string) {
	context.AbortWithStatusJSON(
		http.StatusBadRequest,
		&ExecuteResponse{Status: FAILED, Error: err},
	)
}

func FailedExecuteInternalError(context *gin.Context, stdout, stderr string, err error) {
	context.AbortWithStatusJSON(
		http.StatusInternalServerError,
		&ExecuteResponse{FAILED, stdout, stderr, NillabeError(err)},
	)
}

func NillabeError(err error) string {
	if err != nil {
		return err.Error()
	}
	return ""
}
