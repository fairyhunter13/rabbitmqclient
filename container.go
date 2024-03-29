package rabbitmqclient

import (
	"sync"

	"github.com/fairyhunter13/amqpwrapper"
	"github.com/streadway/amqp"
)

// Container is the struct to make custom Publisher and Consumer.
type Container struct {
	publisherManager *publisherManager
	initiator        *initiator
	conn             amqpwrapper.IConnectionManager
	*saver

	mutex sync.RWMutex
	// mutex protects the following fields
	exchange *ExchangeDeclare
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
		conn:             conn,
		saver:            newSaver(),
		exchange:         new(ExchangeDeclare).Default(),
		Topology:         NewTopology(),
	}
	return
}

// GetConnection return the underlying connection manager.
func (c *Container) GetConnection() amqpwrapper.IConnectionManager {
	return c.conn
}

// GetInitiatorChannel gets the initiator channel from the connection manager.
func (c *Container) GetInitiatorChannel() (*amqp.Channel, error) {
	return c.initiator.getChannel()
}

// Publish publishes the message to the default exchange with the default topic.
func (c *Container) Publish(exchange, topic string, arg OtherPublish) (err error) {
	c.Save()
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

// Consumer creates a new consumer.
func (c *Container) Consumer() *Consumer {
	c.Save()
	return newConsumer(c)
}

// Save saves the current global exchange of the saver implementator.
func (c *Container) Save() *Container {
	if !c.isSaved() {
		c.AddExchangeDeclare(*c.GetGlobalExchange())
		c.save()
	}
	return c
}

// Close closes all the resources inside the container.
func (c *Container) Close() error {
	err := c.conn.Close()
	c.publisherManager.close()
	return err
}
