package handlers

import (
	"go-vozyengine-events-package/controllers"

	"github.com/gin-gonic/gin"
)

func Topic(ctx *gin.Context) {
	if ctx.Request.Method == "GET" {
		controllers.ListTopicsController(ctx)
	} else {
		controllers.SaveTopicController(ctx)
	}
}

func Subscription(ctx *gin.Context) {
	if ctx.Request.Method == "GET" {
		controllers.ListSubscriptionsController(ctx)
	} else {
		controllers.SaveSubscriptionsController(ctx)
	}
}

func Publish(ctx *gin.Context) {
	if ctx.Request.Method == "POST" {
		controllers.SavePublishController(ctx)
	}
}
