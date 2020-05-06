package rabbitmqclient

import (
	"github.com/fairyhunter13/amqpwrapper"
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
		Key:      DefaultKeyInitiator,
		TypeChan: DefaultTypeProducer,
	}
	_, err = i.conn.InitChannelAndGet(TopologyInitializationChannel(topo), args)
	return
}
