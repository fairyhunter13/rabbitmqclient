package rabbitmqclient

import "github.com/fairyhunter13/amqpwrapper"

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

func Publish()
