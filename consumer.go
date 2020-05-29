package rabbitmqclient

import (
	"sync"

	"github.com/streadway/amqp"
)

// Consumer defines the behavior class for the consumer
type Consumer struct {
	container *Container
	*saver
	mutex *sync.RWMutex
	// mutex protects the following fields
	declare *QueueDeclare
	bind    *QueueBind
}

// Handler defines the handler type for this rabbitmqclient
type Handler func(ch *amqp.Channel, msg amqp.Delivery)

func newConsumer(container *Container) *Consumer {
	return &Consumer{
		container: container,
		mutex:     new(sync.RWMutex),
		declare:   new(QueueDeclare).Default(),
		bind:      new(QueueBind).Default(),
		saver:     newSaver(),
	}
}

// Consume consumes the message using the number of workers and handler passed.
func (c *Consumer) Consume(workers uint64, handler Handler) (err error) {
	return
}

// Save saves the consumer state.
func (c *Consumer) Save() *Consumer {
	if !c.isSaved() {
		c.container.AddQueueDeclare(*c.GetQueueDeclare())
		c.container.AddQueueBind(*c.GetQueueBind())
		c.save()
	}
	return c
}
