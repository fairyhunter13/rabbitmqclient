package rabbitmqclient

import (
	"github.com/fairyhunter13/amqpwrapper"
	"github.com/streadway/amqp"
)

var (
	// EmptyChannel specifies the empty channel for initializing rabbitmq channel.
	EmptyChannel = func(ch *amqp.Channel) (err error) {
		return
	}
	// TopologyInitializationChannel is a channel function initialize all declarations inside the passed topology.
	TopologyInitializationChannel = func(topo *Topology) (result amqpwrapper.InitializeChannel) {
		result = func(ch *amqp.Channel) (err error) {
			if topo == nil {
				return
			}
			err = topo.DeclareAll(ch)
			return
		}
		return
	}
)

func generateExchangeName(isPrefix bool, name string) string {
	if isPrefix {
		return DefaultPrefixExchange + name
	}
	return name
}

func generateQueueName(isPrefix bool, name string) string {
	if isPrefix {
		return DefaultPrefixQueue + name
	}
	return name
}
