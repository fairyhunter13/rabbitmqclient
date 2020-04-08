package rabbitmqclient

import "github.com/fairyhunter13/amqpwrapper"

// Container is the struct to make custom Publisher and Consumer.
type Container struct {
	globalExchange   *ExchangeDeclareArgs
	publisherManager *PublisherManager
}

// NewContainer return the container of the connection manager for amqp.Wrapper
func NewContainer(conn amqpwrapper.IConnectionManager) (res *Container, err error) {
	if conn == nil {
		err = amqpwrapper.ErrNilArg
		return
	}
	res = new(Container)
	res.publisherManager = newPublisherManager(conn)
	return
}
