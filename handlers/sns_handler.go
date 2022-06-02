package handlers

import (
	"github.com/gin-gonic/gin"
	"go-vozyengine-events-package/controllers"
)

func Topic(ctx *gin.Context) {
	if ctx.Request.Method == "GET" {
		controllers.ListTopicsController(ctx)
	} else {
		controllers.SaveTopicController(ctx)
	}
}
