package args

import (
	"github.com/fairyhunter13/rabbitmqclient/constant"
	"github.com/fairyhunter13/rabbitmqclient/generator"
	"github.com/streadway/amqp"
)

// QueueDeclare specifies the arguments to declare queue.
type QueueDeclare struct {
	Name       string
	Durable    bool
	AutoDelete bool
	Exclusive  bool
	NoWait     bool
}

// Default sets the default values of the struct variables.
func (q *QueueDeclare) Default() *QueueDeclare {
	q.Name = constant.DefaultQueue
	q.Durable = true
	q.AutoDelete = false
	q.Exclusive = false
	q.NoWait = false
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

// Default sets the default values of the struct variables.
func (q *QueueBind) Default() *QueueBind {
	q.Name = constant.DefaultQueue
	q.Key = constant.DefaultTopic
	q.Exchange = generator.GenerateExchangeName(true, constant.TypeDirect)
	q.NoWait = false
	return q
}

// QueueDelete deletes deletes the bindings, purges the queue, and remove it from the server.
type QueueDelete struct {
	Name     string
	IfUnused bool
	IfEmpty  bool
	NoWait   bool
}

// QueueUnbind removes the bindings of the queue to an exchange.
type QueueUnbind struct {
	Name     string
	Key      string
	Exchange string
	args     amqp.Table
}
