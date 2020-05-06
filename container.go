package rabbitmqclient

import (
	"sync"
	"sync/atomic"

	"github.com/fairyhunter13/amqpwrapper"
	"github.com/fairyhunter13/rabbitmqclient/args"
	"github.com/fairyhunter13/rabbitmqclient/constant"
)

// Container is the struct to make custom Publisher and Consumer.
type Container struct {
	publisherManager *publisherManager
	Initiator        *initiator
	*Topology
	mutex *sync.RWMutex
	// Mutex protects the following fields
	globalExchange *args.ExchangeDeclare

	savedStatus uint64
}

// NewContainer return the container of the connection manager for amqp.Wrapper
func NewContainer(conn amqpwrapper.IConnectionManager) (res *Container, err error) {
	if conn == nil {
		err = amqpwrapper.ErrNilArg
		return
	}
	res = &Container{
		publisherManager: newPublisherManager(conn),
		Initiator:        newInitiator(conn),
		mutex:            new(sync.RWMutex),
		Topology:         NewTopology(),
	}
	return
}

// Publish publishes the message to the default exchange with the default topic.
func (c *Container) Publish(exchange, topic string, arg args.OtherPublish) (err error) {
	if !c.isSaved() {
		err = ErrContainerMustBeSavedFirst
		return
	}
	if exchange == "" {
		c.mutex.RLock()
		exchange = c.globalExchange.Name
		c.mutex.RUnlock()
	}
	if topic == "" {
		topic = constant.DefaultTopic
	}
	err = c.Init()
	if err != nil {
		return
	}
	err = c.publisherManager.publish(
		args.Publish{
			Exchange:     exchange,
			Key:          topic,
			OtherPublish: arg,
		},
	)
	return
}

// SetExchange sets the exchange of this container.
func (c *Container) SetExchange(exc *args.ExchangeDeclare) *Container {
	c.setDefaultExchange()
	if exc != nil {
		c.mutex.Lock()
		c.globalExchange = exc
		c.mutex.Unlock()
	}
	return c
}

// SetExchangeName sets the exchange name of all publishers and consumers
func (c *Container) SetExchangeName(name string) *Container {
	c.setDefaultExchange()
	if name != "" {
		c.mutex.Lock()
		c.globalExchange.Name = name
		c.mutex.Unlock()
	}
	return c
}

// Save saves the current global exchange of the container.
// Save must be called before calling Publish or Consume function.
func (c *Container) Save() *Container {
	if !c.isSaved() {
		c.setDefaultExchange()
		c.mutex.RLock()
		c.AddExchangeDeclare(*c.globalExchange)
		c.mutex.RUnlock()
		c.save()
	}
	return c
}

func (c *Container) save() {
	atomic.StoreUint64(&c.savedStatus, constant.TrueUint)
}

func (c *Container) isSaved() bool {
	return atomic.LoadUint64(&c.savedStatus) == constant.TrueUint
}
