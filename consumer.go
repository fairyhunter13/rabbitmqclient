package rabbitmqclient

import "sync"

// Consumer defines the behavior class for the consumer
type Consumer struct {
	container *Container
	*saver
	mutex *sync.RWMutex
	// mutex protects the following fields
	declare *QueueDeclare
	bind    *QueueBind
}

func newConsumer(container *Container) *Consumer {
	return &Consumer{
		container: container,
		mutex:     new(sync.RWMutex),
		declare:   new(QueueDeclare).Default(),
		bind:      new(QueueBind).Default(),
		saver:     newSaver(),
	}
}

// Save saves the consumer state.
func (c *Consumer) Save() *Consumer {
	if !c.isSaved() {
		c.container.AddQueueDeclare(*c.declare)
		c.container.AddQueueBind(*c.bind)
		c.save()
	}
	return c
}
