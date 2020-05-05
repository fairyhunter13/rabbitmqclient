package args

import "github.com/streadway/amqp"

// Publish specifies the arguments to publish function
type Publish struct {
	Exchange string
	Key      string
	OtherPublish
}

// OtherPublish specifies the others arg for the publish function.
type OtherPublish struct {
	Mandatory bool
	Immediate bool
	Msg       amqp.Publishing
}
