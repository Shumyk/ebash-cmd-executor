package main

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

func successExecuteResponse(stdout, stderr, exitCode string) *ExecuteResponse {
	return &ExecuteResponse{SUCCESS, stdout, stderr, exitCode}
}

func failedExecuteBadRequest(context *gin.Context, err string) {
	context.AbortWithStatusJSON(
		http.StatusBadRequest,
		&ExecuteResponse{Status: FAILED, Error: err},
	)
}

func failedExecuteInternalError(context *gin.Context, stdout, stderr, err string) {
	context.AbortWithStatusJSON(
		http.StatusInternalServerError,
		&ExecuteResponse{FAILED, stdout, stderr, err},
	)
}

func errorDefault(err error) string {
	if err != nil {
		return err.Error()
	}
	return ""
}
