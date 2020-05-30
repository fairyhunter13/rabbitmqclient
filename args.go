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
	if contentType != "" {
		op.Msg.ContentType = contentType
	}
	return op
}

// SetContentEncoding sets the content encoding of the payload
func (op *OtherPublish) SetContentEncoding(contentEncoding string) *OtherPublish {
	if contentEncoding != "" {
		op.Msg.ContentEncoding = contentEncoding
	}
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
	op.Msg.Headers = amqp.Table{}
	if header != nil {
		op.Msg.Headers = header
	}
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
	queueName := DefaultQueue
	if name != "" {
		queueName = name
	}
	c.Queue = queueName
	return c
}

// SetAutoAck sets the auto ack to true for automatic acknowledgement.
// For manual acknowledgement, don't call this function.
func (c *Consume) SetAutoAck() *Consume {
	c.AutoAck = true
	return c
}
