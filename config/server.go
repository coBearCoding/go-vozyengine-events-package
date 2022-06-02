package config

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go-vozyengine-events-package/handlers"
	"log"
	"os"
	"strconv"
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

	sns := r.Group(baseUrl + "/sns")
	sns.GET("/topics", handlers.Topic)

	if port != "" {
		r.Run(":" + port)
	} else {
		r.Run()
	}
}
