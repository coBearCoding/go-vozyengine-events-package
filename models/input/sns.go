package input

type TopicInput struct {
	Topic string `json:"topic"`
}

type PublishMessageInput struct {
	TopicArn string `json:"topic_arn"`
	Message  string `json:"message"`
}

type SubscriptionInput struct {
	TopicArn string `json:"topic_arn"`
	Protocol string `json:"protocol"`
	EndPoint string `json:"endpoint"`
}
