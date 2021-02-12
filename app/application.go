package app

import (
	"github.com/gin-gonic/gin"
	"github.com/sial-soft/users-api/logger"
	"log"
)

var (
	router = gin.Default()
)

func StartApplication() {
	mapUrls()
	logger.Info("server is starting")
	log.Fatal(router.Run(":8080"))
}
