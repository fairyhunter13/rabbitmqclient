package rabbitmqclient

import "github.com/fairyhunter13/amqpwrapper"

// Maker is the struct to make custom Publisher and Consumer.
type Maker struct {
	conn amqpwrapper.IConnectionManager
}
