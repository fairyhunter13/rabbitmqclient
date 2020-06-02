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

// SetExchange is a setter.
func (p *Publish) SetExchange(exchange string) *Publish {
	p.Exchange = exchange
	return p
}

// SetKey is a setter.
func (p *Publish) SetKey(key string) *Publish {
	p.Key = key
	return p
}

// SetOtherPublish is a setter.
func (p *Publish) SetOtherPublish(publish OtherPublish) *Publish {
	p.OtherPublish = publish
	return p
}

// OtherPublish specifies the others arg for the publish function.
type OtherPublish struct {
	Mandatory bool
	Immediate bool
	Msg       amqp.Publishing
}

// SetMandatory is a setter.
func (op *OtherPublish) SetMandatory(mandatory bool) *OtherPublish {
	op.Mandatory = mandatory
	return op
}

// SetImmediate is a setter.
func (op *OtherPublish) SetImmediate(immediate bool) *OtherPublish {
	op.Immediate = immediate
	return op
}

// SetMsg is a setter.
func (op *OtherPublish) SetMsg(msg amqp.Publishing) *OtherPublish {
	op.Msg = msg
	return op
}

// Method list for publishing body

// SetContentType sets the content type of message.
func (op *OtherPublish) SetContentType(contentType string) *OtherPublish {
	op.Msg.ContentType = contentType
	return op
}

// SetContentEncoding sets the content encoding of the payload.
func (op *OtherPublish) SetContentEncoding(contentEncoding string) *OtherPublish {
	op.Msg.ContentEncoding = contentEncoding
	return op
}

// SetPersistent sets the delivery mode to persistent.
func (op *OtherPublish) SetPersistent() *OtherPublish {
	op.Msg.DeliveryMode = DeliveryModePersistent
	return op
}

// SetBody sets the body payload of publish message.
func (op *OtherPublish) SetBody(payload []byte) *OtherPublish {
	op.Msg.Body = payload
	return op
}

// SetHeaders sets the headers for the amqp rabbitmq message.
func (op *OtherPublish) SetHeaders(header amqp.Table) *OtherPublish {
	op.Msg.Headers = header
	return op
}

// SetPriority sets the priority of the message.
func (op *OtherPublish) SetPriority(priority uint8) *OtherPublish {
	op.Msg.Priority = priority
	return op
}

// SetReplyTo sets the reply address for the RPC.
func (op *OtherPublish) SetReplyTo(replyTo string) *OtherPublish {
	op.Msg.ReplyTo = replyTo
	return op
}

// SetExpiration sets the message expiration specification.
func (op *OtherPublish) SetExpiration(expiration string) *OtherPublish {
	op.Msg.Expiration = expiration
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

// SetQueue sets the queue name of the consume arguments.
func (c *Consume) SetQueue(name string) *Consume {
	c.Queue = name
	return c
}

// SetConsumer sets the name of the consumer.
func (c *Consume) SetConsumer(name string) *Consume {
	c.Consumer = name
	return c
}

// SetAutoAck sets the auto ack to true for automatic acknowledgement.
func (c *Consume) SetAutoAck(autoAck bool) *Consume {
	c.AutoAck = autoAck
	return c
}

// SetExclusive sets the exclusive attribute to the consumer.
func (c *Consume) SetExclusive(exclusive bool) *Consume {
	c.Exclusive = exclusive
	return c
}

// SetNoLocal sets the no local attribute of consumer.
func (c *Consume) SetNoLocal(noLocal bool) *Consume {
	c.NoLocal = noLocal
	return c
}

// SetNoWait sets the no wait attribute when consuming.
func (c *Consume) SetNoWait(noWait bool) *Consume {
	c.NoWait = noWait
	return c
}

// SetArgs sets the args of the consume function in amqp.
func (c *Consume) SetArgs(args amqp.Table) *Consume {
	c.Args = args
	return c
}
