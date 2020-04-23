package rabbitmqclient

import (
	"sync"

	"github.com/fairyhunter13/amqpwrapper"
)

// Container is the struct to make custom Publisher and Consumer.
type Container struct {
	publisherManager *PublisherManager
	mutex            *sync.RWMutex
	// Mutex protects the following fields
	globalExchange *ExchangeDeclareArgs
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
		mutex:            new(sync.RWMutex),
		Topology:         NewTopology(),
	}
	return
}

// Publish publishes the message to the default exchange with the default topic.
func (c *Container) Publish(exchange, topic string, args OtherPublishArgs) (err error) {
	c.setDefaultExchange()
	if exchange == "" {
		c.mutex.RLock()
		exchange = c.globalExchange.Name
		c.mutex.RUnlock()
	}
	if topic == "" {
		topic = DefaultTopic
	}
	err = c.publisherManager.Publish(
		PublishArgs{
			Exchange:         exchange,
			Key:              topic,
			OtherPublishArgs: args,
		},
	)
	return
}

// SetExchange sets the exchange of this container.
func (c *Container) SetExchange(exc *ExchangeDeclareArgs) *Container {
	c.setDefaultExchange()
	c.mutex.Lock()
	defer c.mutex.Unlock()
	if exc != nil {
		c.globalExchange = exc
	}
	return c
}

// SetExchangeName sets the exchange name of all publishers and consumers
func (c *Container) SetExchangeName(name string) *Container {
	c.setDefaultExchange()
	c.mutex.Lock()
	defer c.mutex.Unlock()
	if name != "" {
		c.globalExchange.Name = name
	}
	c.AddExchangeDeclare(*c.globalExchange)
	return c
}
