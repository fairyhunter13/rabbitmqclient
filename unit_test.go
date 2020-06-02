package rabbitmqclient

import (
	"testing"

	"github.com/streadway/amqp"
	"github.com/stretchr/testify/assert"
)

func TestBuildArgs(t *testing.T) {
	publish := new(OtherPublish).SetContentType("application/json").SetContentEncoding("utf8").SetHeaders(amqp.Table{})
	assert.EqualValues(t, &OtherPublish{
		Msg: amqp.Publishing{
			ContentEncoding: "utf8",
			ContentType:     "application/json",
			Headers:         amqp.Table{},
		},
	}, publish)

	consume := new(Consume).SetName("test").SetAutoAck(true).SetArgs(amqp.Table{})
	assert.EqualValues(t, &Consume{
		Consumer: "test",
		AutoAck:  true,
		Args:     amqp.Table{},
	}, consume)
}
