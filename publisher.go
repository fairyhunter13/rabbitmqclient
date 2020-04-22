package rabbitmqclient

import (
	"fmt"
	"sync/atomic"

	"github.com/fairyhunter13/amqpwrapper"
	"github.com/streadway/amqp"
)

const (
	// DefaultProducerKey specifies the key for producer.
	DefaultProducerKey = "producer.%d"
	// DefaultProducerType specifies the default type of channel for the amqpwrapper.
	DefaultProducerType = amqpwrapper.Producer
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
func (pm *PublisherManager) Publish(args PublishArgs) (err error) {
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
	err = pm.publish(idChannel, args, isNew)
	return
}

func (pm *PublisherManager) publish(idChan uint64, args PublishArgs, isNew bool) (err error) {
	var (
		ch *amqp.Channel
	)
	ch, err = pm.getChannel(idChan, isNew)
	if err != nil {
		return
	}

	err = ch.Publish(args.Exchange, args.Key, args.Mandatory, args.Immediate, args.Msg)
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
	keyChannel = fmt.Sprintf(DefaultProducerKey, idChan)
	if isNew {
		ch, err = pm.conn.InitChannelAndGet(EmptyChannel, amqpwrapper.InitArgs{
			Key:      keyChannel,
			TypeChan: DefaultProducerType,
		})
	} else {
		ch, err = pm.conn.GetChannel(keyChannel, DefaultProducerType)
	}
	return
}
