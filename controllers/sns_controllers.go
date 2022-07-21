package controllers

import (
	"fmt"
	"go-vozyengine-events-package/awsinternal"
	"go-vozyengine-events-package/models/input"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListTopicsController(ctx *gin.Context) {

	result := awsinternal.ListTopics()

	ctx.JSON(http.StatusOK, result)
}

func SaveTopicController(ctx *gin.Context) {
	topicInput := input.TopicInput{}

	if err := ctx.ShouldBindJSON(&topicInput); err != nil {
		fmt.Println(err.Error())
		ctx.JSON(http.StatusInternalServerError, "topic input cannot be null")
	}

	fmt.Println(topicInput.Topic)

	result := awsinternal.CreateTopic(topicInput.Topic)

	ctx.JSON(http.StatusOK, result)
}

func ListSubscriptionsController(ctx *gin.Context) {
	result := awsinternal.ListSubscriptions()

	ctx.JSON(http.StatusOK, result)
}

func SaveSubscriptionsController(ctx *gin.Context) {
	subscriptionInput := input.SubscriptionInput{}

	if err := ctx.ShouldBindJSON(&subscriptionInput); err != nil {
		fmt.Println(err.Error())
		ctx.JSON(http.StatusInternalServerError, "params input cannot be null")
	}

	fmt.Println(subscriptionInput.TopicArn)

	result := awsinternal.CreateSubscription(subscriptionInput.TopicArn, subscriptionInput.Protocol, subscriptionInput.EndPoint)

	ctx.JSON(http.StatusOK, result)
}

func SavePublishController(ctx *gin.Context) {
	publishInput := input.PublishMessageInput{}
	if err := ctx.ShouldBindJSON(&publishInput); err != nil {
		fmt.Println(err.Error())
		ctx.JSON(http.StatusInternalServerError, "params input cannot be null")
	}

	fmt.Println(publishInput.TopicArn)

	result := awsinternal.PublishMessage(publishInput.TopicArn, publishInput.Message)

	ctx.JSON(http.StatusOK, result)
}
