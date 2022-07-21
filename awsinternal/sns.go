package awsinternal

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/service/sns"
)

func ListTopics() *sns.ListTopicsOutput {
	cfg := LoadConfiguration()

	client := sns.NewFromConfig(cfg)

	input := &sns.ListTopicsInput{}

	result, err := client.ListTopics(context.TODO(), input)
	if err != nil {
		log.Printf("SNS Listing Topic Error: %v \n", err)
		return nil
	}

	return result

}

func CreateTopic(topic string) *sns.CreateTopicOutput {
	cfg := LoadConfiguration()
	client := sns.NewFromConfig(cfg)

	input := &sns.CreateTopicInput{
		Name: &topic,
	}

	result, err := client.CreateTopic(context.TODO(), input)

	if err != nil {
		log.Printf("SNS Creating Topic Error: %v \n", err)
		return nil
	}

	return result
}

func CreateSubscription(topicArn string, protocol string, endpoint string) *sns.SubscribeOutput {
	cfg := LoadConfiguration()
	client := sns.NewFromConfig(cfg)

	/*Examples for subscription Endpoint can be found here:
	https://aws.github.io/aws-sdk-go-v2/docs/code-examples/sns/subscribe/
	https://docs.aws.amazon.com/sns/latest/dg/example_sns_Subscribe_Lambda_section.html
	*/

	input := &sns.SubscribeInput{
		TopicArn:              &topicArn,
		Protocol:              &protocol,
		Endpoint:              &endpoint,
		ReturnSubscriptionArn: true,
	}

	result, err := client.Subscribe(context.TODO(), input)

	if err != nil {
		log.Printf("SNS Creating Topic Error: %v \n", err)
		return nil
	}

	return result
}

func ListSubscriptions() *sns.ListSubscriptionsOutput {
	cfg := LoadConfiguration()

	client := sns.NewFromConfig(cfg)

	input := &sns.ListSubscriptionsInput{}

	result, err := client.ListSubscriptions(context.TODO(), input)

	if err != nil {
		log.Printf("SNS List Subscriptions Error: %v \n", err)
		return nil
	}

	return result
}

func ListSubscriptionsByTopic(topic string) *sns.ListSubscriptionsByTopicOutput {
	cfg := LoadConfiguration()

	client := sns.NewFromConfig(cfg)

	input := &sns.ListSubscriptionsByTopicInput{
		TopicArn: &topic,
	}

	result, err := client.ListSubscriptionsByTopic(context.TODO(), input)

	if err != nil {
		log.Printf("SNS List Subscription By Topic Error: %v \n", err)
		return nil
	}

	return result
}

func PublishMessage(topic string, msg string) *sns.PublishOutput {
	cfg := LoadConfiguration()

	client := sns.NewFromConfig(cfg)

	input := &sns.PublishInput{
		Message:  &msg,
		TopicArn: &topic,
	}

	result, err := client.Publish(context.TODO(), input)
	if err != nil {
		log.Printf("SNS List Subscription By Topic Error: %v \n", err)
		return nil
	}

	return result
}
