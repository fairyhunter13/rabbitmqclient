package args

import (
	"github.com/fairyhunter13/rabbitmqclient/constant"
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
	op.Msg.DeliveryMode = constant.DelvieryModePersistent
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
