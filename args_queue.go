package rabbitmqclient

import (
	"github.com/streadway/amqp"
)

// QueueDeclare specifies the arguments to declare queue.
type QueueDeclare struct {
	Name       string
	Durable    bool
	AutoDelete bool
	Exclusive  bool
	NoWait     bool
	Args       amqp.Table
}

// Default sets the default values of the struct variables.
func (q *QueueDeclare) Default() *QueueDeclare {
	q.Name = DefaultQueue
	q.Durable = true
	q.AutoDelete = false
	q.Exclusive = false
	q.NoWait = false
	q.Args = amqp.Table{}
	return q
}

// QueueDeclarePassive declares the queue if it isn't already declared
type QueueDeclarePassive struct {
	QueueDeclare
}

// QueueBind specifies the arguments to declare binding of queue to exchange.
type QueueBind struct {
	Name     string
	Key      string
	Exchange string
	NoWait   bool
	Args     amqp.Table
}

// SetExchange sets the exchange of the queue binding.
func (q *QueueBind) SetExchange(name string) *QueueBind {
	if name != "" {
		q.Exchange = name
	}
	return q
}

// Default sets the default values of the struct variables.
func (q *QueueBind) Default() *QueueBind {
	q.Name = DefaultQueue
	q.Key = DefaultTopic
	q.Exchange = DefaultExchange
	q.NoWait = false
	return q
}

// QueueUnbind removes the bindings of the queue to an exchange.
type QueueUnbind struct {
	Name     string
	Key      string
	Exchange string
	Args     amqp.Table
}

// QueueDelete deletes deletes the bindings, purges the queue, and remove it from the server.
type QueueDelete struct {
	Name     string
	IfUnused bool
	IfEmpty  bool
	NoWait   bool
}
