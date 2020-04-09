package rabbitmqclient

import "github.com/streadway/amqp"

// PublishArgs specifies the arguments to publish function
type PublishArgs struct {
	Exchange  string
	Key       string
	Mandatory bool
	Immediate bool
	Msg       amqp.Publishing
}
