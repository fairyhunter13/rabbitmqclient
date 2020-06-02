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

// SetName is a setter.
func (q *QueueDeclare) SetName(name string) *QueueDeclare {
	q.Name = name
	return q
}

// SetDurable is a setter.
func (q *QueueDeclare) SetDurable(durable bool) *QueueDeclare {
	q.Durable = durable
	return q
}

// SetAutoDelete is a setter.
func (q *QueueDeclare) SetAutoDelete(delete bool) *QueueDeclare {
	q.AutoDelete = delete
	return q
}

// SetExclusive is a setter.
func (q *QueueDeclare) SetExclusive(exclusive bool) *QueueDeclare {
	q.Exclusive = exclusive
	return q
}

// SetNoWait is a setter.
func (q *QueueDeclare) SetNoWait(noWait bool) *QueueDeclare {
	q.NoWait = noWait
	return q
}

// SetArgs is a setter.
func (q *QueueDeclare) SetArgs(args amqp.Table) *QueueDeclare {
	q.Args = args
	return q
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

// SetName is a setter.
func (q *QueueBind) SetName(name string) *QueueBind {
	q.Name = name
	return q
}

// SetKey is a setter.
func (q *QueueBind) SetKey(key string) *QueueBind {
	q.Key = key
	return q
}

// SetExchange is a setter.
func (q *QueueBind) SetExchange(name string) *QueueBind {
	q.Exchange = name
	return q
}

// SetNoWait is a setter.
func (q *QueueBind) SetNoWait(noWait bool) *QueueBind {
	q.NoWait = noWait
	return q
}

// SetArgs is a setter.
func (q *QueueBind) SetArgs(args amqp.Table) *QueueBind {
	q.Args = args
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
	QueueBind
}

// QueueDelete deletes deletes the bindings, purges the queue, and remove it from the server.
type QueueDelete struct {
	Name     string
	IfUnused bool
	IfEmpty  bool
	NoWait   bool
}

// SetName is a setter.
func (q *QueueDelete) SetName(name string) *QueueDelete {
	q.Name = name
	return q
}

// SetIfUnused is a setter.
func (q *QueueDelete) SetIfUnused(ifUnused bool) *QueueDelete {
	q.IfUnused = ifUnused
	return q
}

// SetIfEmpty is a setter.
func (q *QueueDelete) SetIfEmpty(ifEmpty bool) *QueueDelete {
	q.IfEmpty = ifEmpty
	return q
}

// SetNoWait is a setter.
func (q *QueueDelete) SetNoWait(noWait bool) *QueueDelete {
	q.NoWait = noWait
	return q
}
