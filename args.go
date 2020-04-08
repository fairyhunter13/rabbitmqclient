package rabbitmqclient

import "github.com/streadway/amqp"

// Topology contains all declarations needed to define the topology in the rabbitmq.
type Topology struct {
	ExchangeDeclareArgs []ExchangeDeclareArgs
	QueueDeclareArgs    []QueueDeclareArgs
	QueueBindArgs       []QueueBindArgs
}

// ExchangeDeclareArgs specifies the arguments to used in declaring exchange.
type ExchangeDeclareArgs struct {
	Name       string
	Kind       string
	Durable    bool
	AutoDelete bool
	Internal   bool
	NoWait     bool
	Args       amqp.Table
}

// Default sets the default values of the struct variables.
func (e *ExchangeDeclareArgs) Default() *ExchangeDeclareArgs {
	e.Name = GenerateExchangeName(true, TypeDirect)
	e.Kind = TypeDirect
	e.Durable = true
	e.AutoDelete = true
	e.Internal = false
	e.NoWait = false
	e.Args = amqp.Table{}
	return e
}

// QueueDeclareArgs specifies the arguments to declare queue.
type QueueDeclareArgs struct {
	Name       string
	Durable    bool
	AutoDelete bool
	Exclusive  bool
	NoWait     bool
	Args       amqp.Table
}

// Default sets the default values of the struct variables.
func (q *QueueDeclareArgs) Default() *QueueDeclareArgs {
	q.Name = DefaultQueue
	q.Durable = true
	q.AutoDelete = false
	q.Exclusive = false
	q.NoWait = false
	q.Args = amqp.Table{}
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
	q.Key = DefaultKey
	q.Exchange = GenerateExchangeName(true, TypeDirect)
	q.NoWait = false
	q.Args = amqp.Table{}
	return q
}
