package rabbitmqclient

import (
	"fmt"
	"reflect"

	"github.com/fairyhunter13/amqpwrapper"
	"github.com/streadway/amqp"
)

var (
	// EmptyChannel specifies the empty channel for initializing rabbitmq channel.
	EmptyChannel = func(ch *amqp.Channel) (err error) {
		return
	}
	// TopologyInitializationChannel is a channel function initialize all declarations inside the passed topology.
	TopologyInitializationChannel = func(topo *Topology) (result amqpwrapper.InitializeChannel) {
		result = func(ch *amqp.Channel) (err error) {
			err = iterateDeclare(ch, topo.GetExchangeDeclare())
			if err != nil {
				return
			}
			err = iterateDeclare(ch, topo.GetExchangeDeclarePassive())
			if err != nil {
				return
			}
			return
		}
		return
	}
)

func iterateDeclare(ch *amqp.Channel, list interface{}) (err error) {
	val := reflect.ValueOf(list)
	if !val.IsValid() {
		return
	}
	switch val.Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < val.Len(); i++ {
			err = declare(ch, val.Index(i))
			if err != nil {
				return
			}
		}
	}
	return
}

func declare(ch *amqp.Channel, declaration interface{}) (err error) {
	switch elem := declaration.(type) {
	case ExchangeDeclare:
		err = ch.ExchangeDeclare(elem.Name, elem.Kind, elem.Durable, elem.AutoDelete, elem.Internal, elem.NoWait, elem.Args)
	case ExchangeDeclarePassive:
		err = ch.ExchangeDeclarePassive(elem.Name, elem.Kind, elem.Durable, elem.AutoDelete, elem.Internal, elem.NoWait, elem.Args)
	case ExchangeBind:
		err = ch.ExchangeBind(elem.Destination, elem.Key, elem.Source, elem.NoWait, elem.Args)
	case ExchangeUnbind:
		err = ch.ExchangeUnbind(elem.Destination, elem.Key, elem.Source, elem.NoWait, elem.Args)
	case ExchangeDelete:
		err = ch.ExchangeDelete(elem.Name, elem.IfUnused, elem.NoWait)
	}
	return
}

func generateExchangeName(isPrefix bool, name string) string {
	if isPrefix {
		return fmt.Sprintf("%s%s", DefaultPrefixExchange, name)
	}
	return name
}
