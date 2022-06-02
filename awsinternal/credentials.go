package awsinternal

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"log"
	"os"
)

var (
	awsAccessKey string
	awsSecretKey string
	awsRegion    string
)

func initCredentials() {
	awsAccessKey = os.Getenv("AWS_ACCESS_KEY_ID")
	awsSecretKey = os.Getenv("AWS_SECRET_ACCESS_KEY")
	awsRegion = os.Getenv("AWS_REGION")
}

func LoadConfiguration() aws.Config {
	initCredentials()

	credentialsProvider := credentials.NewStaticCredentialsProvider(awsAccessKey, awsSecretKey, "")

	_, err := credentialsProvider.Retrieve(context.TODO())
	if err != nil {
		log.Println(err)
		return aws.Config{}
	}

	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithCredentialsProvider(credentialsProvider),
		config.WithRegion(awsRegion))

	if err != nil {
		log.Println(err)
		return aws.Config{}
	}

	return cfg

}
