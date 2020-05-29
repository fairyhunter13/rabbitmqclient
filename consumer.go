package rabbitmqclient

import "sync"

// Consumer defines the behavior class for the consumer
type Consumer struct {
	mutex *sync.RWMutex
	// mutex protects the following fields
	declare *QueueDeclare
	bind    *QueueBind
}

func newConsumer() *Consumer {
	return &Consumer{
		mutex:   new(sync.RWMutex),
		declare: new(QueueDeclare).Default(),
		bind:    new(QueueBind).Default(),
	}
}
