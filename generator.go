package rabbitmqclient

import (
	"github.com/fairyhunter13/amqpwrapper"
	"github.com/streadway/amqp"
)

var (
	// EmptyFn specifies the empty function for initializing rabbitmq channel.
	EmptyFn = func(ch *amqp.Channel) (err error) {
		return
	}
	// TopologyInitializationFn is a function to initialize all declarations inside the passed topology.
	TopologyInitializationFn = func(topo *Topology) (result amqpwrapper.InitializeChannel) {
		result = func(ch *amqp.Channel) (err error) {
			if topo == nil {
				return
			}
			err = topo.DeclareAll(ch)
			return
		}
		return
	}
	// QosSetterFn declares the quality of service inside the channel for consumers.
	QosSetterFn = func(workers int) (result amqpwrapper.InitializeChannel) {
		result = func(ch *amqp.Channel) (err error) {
			if workers <= 0 {
				workers = 1
			}
			err = ch.Qos(workers, 0, false)
			return
		}
		return
	}
)

// GenerateExchangeName creates an exchange name with default prefix if isPrefix is true.
func GenerateExchangeName(isPrefix bool, name string) string {
	if isPrefix {
		return DefaultPrefixExchange + name
	}
	return name
}

// GenerateQueueName creates a queue name with default prefix if isPrefix is true.
func GenerateQueueName(isPrefix bool, name string) string {
	if isPrefix {
		return DefaultPrefixQueue + name
	}
	return name
}

// GenerateConsumerChannelKey creates a consumer key for the channel to keeping track inside the IConnectionManager
// with default prefix if isPrefix is true.
func GenerateConsumerChannelKey(isPrefix bool, name string) string {
	if isPrefix {
		return DefaultPrefixConsumer + name
	}
	return name
}
