package rabbitmqclient

import (
	"sync"

	"github.com/fairyhunter13/amqpwrapper"
)

// Container is the struct to make custom Publisher and Consumer.
type Container struct {
	publisherManager *PublisherManager
	topology         *Topology
	mutex            *sync.RWMutex
	// Mutex protects the following fields
	globalExchange *ExchangeDeclareArgs
}

// NewContainer return the container of the connection manager for amqp.Wrapper
func NewContainer(conn amqpwrapper.IConnectionManager) (res *Container, err error) {
	if conn == nil {
		err = amqpwrapper.ErrNilArg
		return
	}
	res = new(Container)
	res.publisherManager = newPublisherManager(conn)
	res.mutex = new(sync.RWMutex)
	res.topology = NewTopology()
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
		topic = DefaultKey
	}
	pmArgs := PublishArgs{
		Exchange:         exchange,
		Key:              topic,
		OtherPublishArgs: args,
	}
	err = c.publisherManager.Publish(pmArgs)
	return
}

// SetExchangeName sets the exchange name of all publishers and consumers
func (c *Container) SetExchangeName(name string) *Container {
	c.setDefaultExchange()
	c.mutex.Lock()
	defer c.mutex.Unlock()
	if name != "" {
		c.globalExchange.Name = name
	}
	c.topology.AddExchangeDeclare(*c.globalExchange)
	return c
}

func (c *Container) setDefaultExchange() {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	if c.globalExchange == nil {
		c.globalExchange = new(ExchangeDeclareArgs).Default()
	}
}
