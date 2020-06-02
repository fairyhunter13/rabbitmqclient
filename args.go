package rabbitmqclient

import (
	"github.com/streadway/amqp"
)

// Publish specifies the arguments to publish function
type Publish struct {
	Exchange string
	Key      string
	OtherPublish
}

// OtherPublish specifies the others arg for the publish function.
type OtherPublish struct {
	Mandatory bool
	Immediate bool
	Msg       amqp.Publishing
}

// SetContentType sets the content type of message.
func (op *OtherPublish) SetContentType(contentType string) *OtherPublish {
	op.Msg.ContentType = contentType
	return op
}

// SetContentEncoding sets the content encoding of the payload
func (op *OtherPublish) SetContentEncoding(contentEncoding string) *OtherPublish {
	op.Msg.ContentEncoding = contentEncoding
	return op
}

// SetPersistent sets the delivery mode to persistent.
func (op *OtherPublish) SetPersistent() *OtherPublish {
	op.Msg.DeliveryMode = DelvieryModePersistent
	return op
}

// SetBody sets the body payload of publish message
func (op *OtherPublish) SetBody(payload []byte) *OtherPublish {
	op.Msg.Body = payload
	return op
}

// SetHeaders sets the headers for the amqp rabbitmq message
func (op *OtherPublish) SetHeaders(header amqp.Table) *OtherPublish {
	op.Msg.Headers = header
	return op
}

// Consume define the consume arguments of amqp.
type Consume struct {
	Queue     string
	Consumer  string
	AutoAck   bool
	Exclusive bool
	NoLocal   bool
	NoWait    bool
	Args      amqp.Table
}

// SetName sets the name of the consumer.
func (c *Consume) SetName(name string) *Consume {
	c.Consumer = name
	return c
}

// SetOnlyOneConsumer sets the exclusive args of consume to true.
// This makes the queue only have one consumer.
func (c *Consume) SetOnlyOneConsumer() *Consume {
	c.Exclusive = true
	return c
}

// SetQueue sets the queue name of the consume arguments.
func (c *Consume) SetQueue(name string) *Consume {
	c.Queue = name
	return c
}

// SetAutoAck sets the auto ack to true for automatic acknowledgement.
func (c *Consume) SetAutoAck(autoAck bool) *Consume {
	c.AutoAck = autoAck
	return c
}

// SetArgs sets the args of the consume function in amqp.
func (c *Consume) SetArgs(args amqp.Table) *Consume {
	c.Args = args
	return c
}
