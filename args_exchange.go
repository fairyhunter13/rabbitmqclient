package rabbitmqclient

import (
	"github.com/streadway/amqp"
)

// ExchangeDeclare specifies the arguments to used in declaring exchange.
type ExchangeDeclare struct {
	Name       string
	Kind       string
	Durable    bool
	AutoDelete bool
	Internal   bool
	NoWait     bool
	Args       amqp.Table
}

// SetName is a setter.
func (e *ExchangeDeclare) SetName(name string) *ExchangeDeclare {
	e.Name = name
	return e
}

// SetKind is a setter.
func (e *ExchangeDeclare) SetKind(kind string) *ExchangeDeclare {
	e.Kind = kind
	return e
}

// SetDurable is a setter.
func (e *ExchangeDeclare) SetDurable(durable bool) *ExchangeDeclare {
	e.Durable = durable
	return e
}

// SetAutoDelete is a setter.
func (e *ExchangeDeclare) SetAutoDelete(autoDelete bool) *ExchangeDeclare {
	e.AutoDelete = autoDelete
	return e
}

// SetInternal is a setter.
func (e *ExchangeDeclare) SetInternal(internal bool) *ExchangeDeclare {
	e.Internal = internal
	return e
}

// SetNoWait is a setter.
func (e *ExchangeDeclare) SetNoWait(noWait bool) *ExchangeDeclare {
	e.NoWait = noWait
	return e
}

// SetArgs is a setter.
func (e *ExchangeDeclare) SetArgs(args amqp.Table) *ExchangeDeclare {
	e.Args = args
	return e
}

// Default sets the default values of the struct variables.
func (e *ExchangeDeclare) Default() *ExchangeDeclare {
	e.Name = DefaultExchange
	e.Kind = TypeDirect
	e.Durable = true
	e.AutoDelete = true
	e.Internal = false
	e.NoWait = false
	return e
}

// ExchangeDeclarePassive declares an exchange as passive to assumes that the exchange is already exist.
type ExchangeDeclarePassive struct {
	ExchangeDeclare
}

// ExchangeBind specifies the argumenst to bind an exchange to other exchange.
type ExchangeBind struct {
	Destination string
	Key         string
	Source      string
	NoWait      bool
	Args        amqp.Table
}

// SetDestination is a setter.
func (e *ExchangeBind) SetDestination(destination string) *ExchangeBind {
	e.Destination = destination
	return e
}

// SetKey is a setter.
func (e *ExchangeBind) SetKey(key string) *ExchangeBind {
	e.Key = key
	return e
}

// SetSource is a setter.
func (e *ExchangeBind) SetSource(source string) *ExchangeBind {
	e.Source = source
	return e
}

// SetNoWait is a setter.
func (e *ExchangeBind) SetNoWait(noWait bool) *ExchangeBind {
	e.NoWait = noWait
	return e
}

// SetArgs is a setter.
func (e *ExchangeBind) SetArgs(args amqp.Table) *ExchangeBind {
	e.Args = args
	return e
}

// ExchangeUnbind unbinds the exchange using the same argument as the exchange bind.
type ExchangeUnbind struct {
	ExchangeBind
}

// ExchangeDelete deletes the exchange when it is already declared.
type ExchangeDelete struct {
	Name     string
	IfUnused bool
	NoWait   bool
}

// SetName is a setter.
func (e *ExchangeDelete) SetName(name string) *ExchangeDelete {
	e.Name = name
	return e
}

// SetIfUnused is a setter.
func (e *ExchangeDelete) SetIfUnused(ifUnused bool) *ExchangeDelete {
	e.IfUnused = ifUnused
	return e
}

// SetNoWait is a setter.
func (e *ExchangeDelete) SetNoWait(noWait bool) *ExchangeDelete {
	e.NoWait = noWait
	return e
}
