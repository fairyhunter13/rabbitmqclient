package rabbitmqclient

import (
	"github.com/fairyhunter13/amqpwrapper"
	"github.com/fairyhunter13/rabbitmqclient/constant"
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
	if topo == nil {
		err = ErrTopologyMustNotBeNil
		return
	}
	args := amqpwrapper.InitArgs{
		Key:      constant.DefaultKeyInitiator,
		TypeChan: constant.DefaultTypeProducer,
	}
	i.conn.InitChannelAndGet()
	return
}
