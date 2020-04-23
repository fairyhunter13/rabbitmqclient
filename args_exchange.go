package rabbitmqclient

import "github.com/streadway/amqp"

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

// ExchangeDeclarePassive declares an exchange as passive to assumes that the exchange is already exist.
type ExchangeDeclarePassive struct {
	ExchangeDeclareArgs
}

// Default sets the default values of the struct variables.
func (e *ExchangeDeclareArgs) Default() *ExchangeDeclareArgs {
	e.Name = GenerateExchangeName(true, TypeDirect)
	e.Kind = TypeDirect
	e.Durable = true
	e.AutoDelete = true
	e.Internal = false
	e.NoWait = false
	return e
}

// ExchangeBindArgs specifies the argumenst to bind an exchange to other exchange.
type ExchangeBindArgs struct {
	Destination string
	Key         string
	Source      string
	NoWait      bool
	Args        amqp.Table
}

// ExchangeUnbindArgs unbinds the exchange using the same argument as the exchange bind.
type ExchangeUnbindArgs struct {
	ExchangeBindArgs
}

// ExchangeDelete deletes the exchange when it is already declared.
type ExchangeDelete struct {
	Name     string
	IfUnused bool
	NoWait   bool
}
