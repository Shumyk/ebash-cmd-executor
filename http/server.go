package http

import (
	"context"
	"ebash/cmd-executor/config"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func Server() *http.Server {
	server := &http.Server{
		Addr:    ":" + config.App().Port,
		Handler: setupRouter(),
	}
	go ListenAndServe(server)
	return server
}

func setupRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/ping", pingGET)
	router.POST("/execute", executePOST)
	println()

	return router
}

func ListenAndServe(server *http.Server) {
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Failed to serve: [%v]", err)
	}
}

func ShutdownServer(server *http.Server) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalln("Server forced to shutdown: ", err)
	}
}
