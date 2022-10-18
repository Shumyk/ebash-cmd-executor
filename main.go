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

func executePOST(context *gin.Context) {
	request := ExecuteRequest{}
	if error := context.BindJSON(&request); error != nil {
		context.AbortWithStatusJSON(
			http.StatusBadRequest,
			&ExecuteResponse{
				"failed",
				fmt.Sprintf("Can't parse a body: %v", error),
			},
		)
		return
	}

	fmt.Printf("Received command to execute: [%v]\n", request.Command)
	responseMessage := fmt.Sprintf("Processed result of: [%v]", request.Command)
	context.JSON(http.StatusOK, &ExecuteResponse{"success", responseMessage})
}

func main() {
	router := setupRouter()
	router.Run(":8080")
}
