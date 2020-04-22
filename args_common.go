package rabbitmqclient

import "github.com/streadway/amqp"

// PublishArgs specifies the arguments to publish function
type PublishArgs struct {
	Exchange string
	Key      string
	OtherPublishArgs
}

// OtherPublishArgs specifies the others arg for the publish function.
type OtherPublishArgs struct {
	Mandatory bool
	Immediate bool
	Msg       amqp.Publishing
}
