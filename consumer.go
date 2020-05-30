package rabbitmqclient

import (
	"sync"

	"github.com/fairyhunter13/amqpwrapper"
	"github.com/streadway/amqp"
)

// Consumer defines the behavior class for the consumer
type Consumer struct {
	container *Container
	*saver
	mutex *sync.RWMutex
	// mutex protects the following fields
	declare    *QueueDeclare
	bind       *QueueBind
	channelKey string
}

// Handler defines the handler type for this rabbitmqclient
type Handler func(ch *amqp.Channel, msg amqp.Delivery)

func newConsumer(container *Container) *Consumer {
	return &Consumer{
		container:  container,
		mutex:      new(sync.RWMutex),
		declare:    new(QueueDeclare).Default(),
		bind:       new(QueueBind).Default(),
		saver:      newSaver(),
		channelKey: DefaultChannelKey,
	}
}

func (c *Consumer) getContainer() *Container {
	return c.container
}

// Consume consumes the message using the number of workers and handler passed.
func (c *Consumer) Consume(workers int, handler Handler) (err error) {
	if handler == nil {
		err = ErrArgumentsMusntBeEmpty
		return
	}
	if workers <= 0 {
		workers = 1
	}
	c.Save()
	err = c.getContainer().Init()
	if err != nil {
		return
	}
	var ch *amqp.Channel
	ch, err = c.initChannel(workers)
	if err != nil {
		return
	}

	return
}

func (c *Consumer) initChannel(workers int) (ch *amqp.Channel, err error) {
	args := amqpwrapper.InitArgs{
		Key:      c.getChannelKey(),
		TypeChan: DefaultTypeConsumer,
	}
	ch, err = c.getContainer().GetConnection().InitChannelAndGet(QosSetterFn(workers), args)
	return
}

// Save saves the consumer state.
func (c *Consumer) Save() *Consumer {
	if !c.isSaved() {
		c.getContainer().AddQueueDeclare(*c.GetQueueDeclare())
		c.getContainer().AddQueueBind(*c.GetQueueBind())
		c.save()
	}
	return c
}
