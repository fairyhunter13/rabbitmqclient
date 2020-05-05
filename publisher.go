package rabbitmqclient

import (
	"fmt"
	"sync/atomic"

	"github.com/fairyhunter13/amqpwrapper"
	"github.com/fairyhunter13/rabbitmqclient/args"
	"github.com/fairyhunter13/rabbitmqclient/constant"
	"github.com/fairyhunter13/rabbitmqclient/generator"
	"github.com/streadway/amqp"
)

// PublisherManager manages all the publisher for the Container.
type PublisherManager struct {
	conn           amqpwrapper.IConnectionManager
	channelCounter uint64
	idleChannels   chan uint64
}

func newPublisherManager(conn amqpwrapper.IConnectionManager) (res *PublisherManager) {
	res = new(PublisherManager)
	res.idleChannels = make(chan uint64)
	return
}

// Publish publishes the message with current arguments.
func (pm *PublisherManager) Publish(arg args.Publish) (err error) {
	var (
		idChannel uint64
		isNew     bool
	)
	if pm.conn.IsClosed() {
		err = ErrConnectionAlreadyClosed
		return
	}
	select {
	case idChannel = <-pm.idleChannels:
		isNew = false
		defer func(idChannel uint64) {
			pm.idleChannels <- idChannel
		}(idChannel)
	default:
		isNew = true
	}
	err = pm.publish(idChannel, arg, isNew)
	return
}

func (pm *PublisherManager) publish(idChan uint64, arg args.Publish, isNew bool) (err error) {
	var (
		ch *amqp.Channel
	)
	ch, err = pm.getChannel(idChan, isNew)
	if err != nil {
		return
	}

	err = ch.Publish(arg.Exchange, arg.Key, arg.Mandatory, arg.Immediate, arg.Msg)
	return
}

func (pm *PublisherManager) getChannel(idChan uint64, isNew bool) (ch *amqp.Channel, err error) {
	var (
		keyChannel string
	)
	if isNew {
		idChan = atomic.LoadUint64(&pm.channelCounter)
		atomic.AddUint64(&pm.channelCounter, 1)
	}
	keyChannel = fmt.Sprintf(constant.DefaultKeyProducer, idChan)
	if isNew {
		ch, err = pm.conn.InitChannelAndGet(generator.EmptyChannel, amqpwrapper.InitArgs{
			Key:      keyChannel,
			TypeChan: constant.DefaultTypeProducer,
		})
	} else {
		ch, err = pm.conn.GetChannel(keyChannel, constant.DefaultTypeProducer)
	}
	return
}
