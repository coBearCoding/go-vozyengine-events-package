package controllers

import (
	"github.com/gin-gonic/gin"
	"go-vozyengine-events-package/awsinternal"
	"net/http"
)

func ListTopicsController(ctx *gin.Context) {

	result := awsinternal.ListTopics()

	ctx.JSON(http.StatusOK, result)
}

func SaveTopicController(ctx *gin.Context) {

}
