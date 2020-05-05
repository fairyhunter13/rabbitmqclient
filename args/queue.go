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
