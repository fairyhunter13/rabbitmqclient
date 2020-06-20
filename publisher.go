package rabbitmqclient

import (
	"fmt"
	"sync/atomic"

	"github.com/fairyhunter13/amqpwrapper"
	"github.com/streadway/amqp"
)

type publisherManager struct {
	conn           amqpwrapper.IConnectionManager
	channelCounter uint64
	idleChannels   chan uint64
}

func newPublisherManager(conn amqpwrapper.IConnectionManager) (res *publisherManager) {
	res = new(publisherManager)
	res.idleChannels = make(chan uint64)
	res.conn = conn
	return
}

func (pm *publisherManager) publish(arg Publish) (err error) {
	if pm.conn.IsClosed() {
		err = ErrConnectionAlreadyClosed
		return
	}

	var (
		idChannel uint64
		isNew     bool
		more      bool
	)
	select {
	case idChannel, more = <-pm.idleChannels:
		if !more {
			return
		}
		isNew = false
	default:
		isNew = true
	}
	defer pm.releaseChannel(&idChannel)
	err = pm.publishWithChannel(&idChannel, arg, isNew)
	return
}

func (pm *publisherManager) publishWithChannel(idChan *uint64, arg Publish, isNew bool) (err error) {
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

func (pm *publisherManager) getChannel(idChan *uint64, isNew bool) (ch *amqp.Channel, err error) {
	var (
		keyChannel string
	)
	if isNew {
		*idChan = atomic.LoadUint64(&pm.channelCounter)
		atomic.AddUint64(&pm.channelCounter, 1)
	}
	keyChannel = fmt.Sprintf(DefaultKeyProducer, *idChan)
	if isNew {
		ch, err = pm.conn.InitChannelAndGet(EmptyFn, amqpwrapper.InitArgs{
			Key:      keyChannel,
			TypeChan: DefaultTypeProducer,
		})
	} else {
		ch, err = pm.conn.GetChannel(keyChannel, DefaultTypeProducer)
	}
	return
}

func (pm *publisherManager) releaseChannel(idChannel *uint64) {
	idLocal := *idChannel
	go func() {
		pm.idleChannels <- idLocal
	}()
}

func (pm *publisherManager) close() {
	close(pm.idleChannels)
}
