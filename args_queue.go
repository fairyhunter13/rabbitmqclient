package rabbitmqclient

import "github.com/streadway/amqp"

// QueueDeclareArgs specifies the arguments to declare queue.
type QueueDeclareArgs struct {
	Name       string
	Durable    bool
	AutoDelete bool
	Exclusive  bool
	NoWait     bool
}

// Default sets the default values of the struct variables.
func (q *QueueDeclareArgs) Default() *QueueDeclareArgs {
	q.Name = DefaultQueue
	q.Durable = true
	q.AutoDelete = false
	q.Exclusive = false
	q.NoWait = false
	return q
}

// QueueBindArgs specifies the arguments to declare binding of queue to exchange.
type QueueBindArgs struct {
	Name     string
	Key      string
	Exchange string
	NoWait   bool
	Args     amqp.Table
}

// Default sets the default values of the struct variables.
func (q *QueueBindArgs) Default() *QueueBindArgs {
	q.Name = DefaultQueue
	q.Key = DefaultTopic
	q.Exchange = GenerateExchangeName(true, TypeDirect)
	q.NoWait = false
	return q
}
