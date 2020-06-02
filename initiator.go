package rabbitmqclient

import (
	"github.com/fairyhunter13/amqpwrapper"
	"github.com/streadway/amqp"
)

type initiator struct {
	conn amqpwrapper.IConnectionManager
}

func newInitiator(conn amqpwrapper.IConnectionManager) *initiator {
	return &initiator{
		conn: conn,
	}
}

func (i *initiator) init(topo *Topology) (err error) {
	args := amqpwrapper.InitArgs{
		Key:      DefaultKeyInitiator,
		TypeChan: DefaultTypeProducer,
	}
	_, err = i.conn.InitChannelAndGet(TopologyInitializationFn(topo), args)
	return
}

func (i *initiator) getChannel() (*amqp.Channel, error) {
	return i.conn.GetChannel(DefaultKeyInitiator, DefaultTypeProducer)
}
