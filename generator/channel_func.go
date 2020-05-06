package generator

import "github.com/streadway/amqp"

var (
	// EmptyChannel specifies the empty channel for initializing rabbitmq channel.
	EmptyChannel = func(ch *amqp.Channel) (err error) {
		return
	}
	// TopologyInitializationChannel is a channel function initialize all declarations inside the passed topology.
	TopologyInitializationChannel = func(topo *)
)
