package rabbitmqclient

import (
	"testing"
	"time"

	"github.com/streadway/amqp"
	"github.com/stretchr/testify/assert"
)

func TestPublishAndSubscribe(t *testing.T) {
	container, err := NewContainer(testSetup.GetConnection())
	assert.Nil(t, err)

	container.SetExchange(new(ExchangeDeclare).Default()).SetExchangeName("integration-test")
	err = container.Publish("", "integration-test", *new(OtherPublish).SetPersistent().SetBody([]byte("test payload")))
	assert.Nil(t, err)

	var result string
	testHandler := func(ch *amqp.Channel, msg amqp.Delivery) {
		msg.Ack(true)
		result = string(msg.Body)
	}
	err = container.Consumer().SetTopic("integration-test").Consume(1, testHandler)
	assert.Nil(t, err)

	time.Sleep(1 * time.Second)
	assert.Equal(t, "test payload", result)
}

func TestBuildArgs(t *testing.T) {
	publish := new(OtherPublish).SetContentType("application/json").SetContentEncoding("utf8").SetHeaders(amqp.Table{})
	assert.EqualValues(t, &OtherPublish{
		Msg: amqp.Publishing{
			ContentEncoding: "utf8",
			ContentType:     "application/json",
			Headers:         amqp.Table{},
		},
	}, publish)

	consume := new(Consume).SetName("test").SetAutoAck().SetArgs(amqp.Table{})
	assert.EqualValues(t, &Consume{
		Consumer: "test",
		AutoAck:  true,
		Args:     amqp.Table{},
	}, consume)
}
