package config

import (
	"go-vozyengine-events-package/handlers"
	"log"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func InitializeServer() {
	_ = godotenv.Load(".env")
	port := os.Getenv("APP_PORT")
	production, _ := strconv.ParseBool(os.Getenv("PRODUCTION"))

	if production {
		createRoutes("")
	} else {
		createRoutes(port)
	}
}

func createRoutes(port string) {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	log.Print("Gin Cold Start")

	baseUrl := "/api/events"

	//SNS - TOPICS
	sns := r.Group(baseUrl + "/sns")
	sns.GET("/topics", handlers.Topic)
	sns.POST("/topics", handlers.Topic)

	//SNS - SUBSCRIPTIONS
	sns.GET("/subscriptions", handlers.Subscription)
	sns.POST("/subscriptions", handlers.Subscription)

	//SNS - PUBLISH MESSAGE
	sns.POST("/publish-message", handlers.Publish)

	if port != "" {
		r.Run(":" + port)
	} else {
		r.Run()
	}
}
