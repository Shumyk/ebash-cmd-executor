package main

import (
	"fmt"
	"io/ioutil"
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
	bodyBytes, _ := ioutil.ReadAll(context.Request.Body)
	body := string(bodyBytes)

	fmt.Printf("Received command to execute: [%v]\n", body)
	message := fmt.Sprintf("Processed result of: [%v]", body)

	context.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": message,
	})
}

func main() {
	router := setupRouter()
	router.Run(":8080")
}
