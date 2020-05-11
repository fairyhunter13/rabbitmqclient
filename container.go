package rabbitmqclient

import (
	"sync"

	"github.com/fairyhunter13/amqpwrapper"
)

// Container is the struct to make custom Publisher and Consumer.
type Container struct {
	publisherManager *publisherManager
	initiator        *initiator
	*globalExchange
	*saver
	mutex *sync.RWMutex
	// mutex protects the following fields
	*Topology
}

// NewContainer return the container of the connection manager for amqp.Wrapper
func NewContainer(conn amqpwrapper.IConnectionManager) (res *Container, err error) {
	if conn == nil {
		err = amqpwrapper.ErrNilArg
		return
	}
	res = &Container{
		publisherManager: newPublisherManager(conn),
		initiator:        newInitiator(conn),
		globalExchange:   newGlobalExchange(),
		saver:            newSaver(),
		mutex:            new(sync.RWMutex),
		Topology:         NewTopology(),
	}
	return
}

// Publish publishes the message to the default exchange with the default topic.
func (c *Container) Publish(exchange, topic string, arg OtherPublish) (err error) {
	if !c.isSaved() {
		err = ErrContainerMustBeSavedFirst
		return
	}
	if exchange == "" {
		exchange = c.GetGlobalExchange().Name
	}
	if topic == "" {
		topic = DefaultTopic
	}
	err = c.Init()
	if err != nil {
		return
	}
	err = c.publisherManager.publish(
		Publish{
			Exchange:     exchange,
			Key:          topic,
			OtherPublish: arg,
		},
	)
	return
}

// Save saves the current global exchange of the saver implementator.
func (c *Container) Save() *Container {
	if !c.isSaved() {
		c.AddExchangeDeclare(*c.GetGlobalExchange())
		c.save()
	}
	return c
}
