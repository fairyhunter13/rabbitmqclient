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
