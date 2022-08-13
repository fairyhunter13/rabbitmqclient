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

	mutex sync.RWMutex
	// mutex protects the following fields
	declare    *QueueDeclare
	bind       *QueueBind
	consume    *Consume
	channelKey string
}

// Handler defines the handler type for this rabbitmqclient
type Handler func(ch *amqp.Channel, msg amqp.Delivery)

func newConsumer(container *Container) *Consumer {
	consumer := &Consumer{
		container:  container,
		saver:      newSaver(),
		declare:    new(QueueDeclare).Default(),
		bind:       new(QueueBind).Default().SetExchange(container.GetGlobalExchange().Name),
		consume:    new(Consume),
		channelKey: DefaultChannelKey,
	}
	return consumer
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
	err = c.initChannelManager(workers)
	if err != nil {
		return
	}
	c.runConsumers(workers, *c.getConsume(), handler)
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

func (c *Consumer) initChannelManager(workers int) (err error) {
	_, err = c.getContainer().GetConnection().InitChannelAndGet(
		QosSetterFn(workers),
		amqpwrapper.InitArgs{
			Key:      c.getChannelKey(),
			TypeChan: DefaultTypeConsumer,
		},
	)
	return
}

func (c *Consumer) runConsumers(workers int, consumeArgs Consume, handler Handler) {
	consumeArgs.SetQueue(c.getQueue())
	if workers == 1 {
		consumeArgs.SetExclusive(true)
	}
	for numWorker := 1; numWorker <= workers; numWorker++ {
		go func() {
			c.loopMessage(consumeArgs, handler)
		}()
	}
}

func (c *Consumer) loopMessage(args Consume, handler Handler) {
	for {
		connection := c.getContainer().GetConnection()
		if connection.IsClosed() {
			return
		}
		ch, err := connection.GetChannel(c.getChannelKey(), DefaultTypeConsumer)
		if err != nil {
			return
		}
		deliveryCh, err := ch.Consume(
			args.Queue,
			args.Consumer,
			args.AutoAck,
			args.Exclusive,
			args.NoLocal,
			args.NoWait,
			args.Args,
		)
		if err != nil {
			continue
		}
		for msg := range deliveryCh {
			handler(ch, msg)
		}
	}
}
