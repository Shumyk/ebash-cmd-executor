package communication

import (
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

func SuccessExecuteResponse(stdout, stderr, exitCode string) *ExecuteResponse {
	return &ExecuteResponse{SUCCESS, stdout, stderr, exitCode}
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

func ErrorDefault(err error) string {
	if err != nil {
		return err.Error()
	}
	return ""
}
