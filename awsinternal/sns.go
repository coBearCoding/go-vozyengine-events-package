package awsinternal

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/sns"
	"log"
)

func ListTopics() *sns.ListTopicsOutput {
	cfg := LoadConfiguration()

	client := sns.NewFromConfig(cfg)

	input := &sns.ListTopicsInput{}

	result, err := client.ListTopics(context.TODO(), input)
	if err != nil {
		log.Printf("SNS Listing Error: %v \n", err)
		return nil
	}

	return result

}
