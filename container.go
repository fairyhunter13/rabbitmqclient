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

// SetExchangeName sets the exchange name of all publishers and consumers
func (c *Container) SetExchangeName(name string) *Container {
	c.setExchangeName(name)
	return c
}

func (c *Container) setExchangeName(name string) {
	c.mutex.RLock()
	if c.globalExchange == nil {
		c.mutex.RUnlock()
		c.mutex.Lock()
		c.globalExchange = new(ExchangeDeclareArgs).Default()
		c.mutex.Unlock()
		c.mutex.RLock()
		c.topology.AddExchangeDeclare(*c.globalExchange)
		c.mutex.RUnlock()
	} else {
		c.mutex.RUnlock()
	}
}
