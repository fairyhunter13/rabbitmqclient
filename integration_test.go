package rabbitmqclient

import (
	"testing"
	"time"

	"github.com/streadway/amqp"
	"github.com/stretchr/testify/assert"
)

func TestPublishAndSubscribe(t *testing.T) {
	container, err := testSetup.GetContainer()
	assert.Nil(t, err)

	container.SetExchange(new(ExchangeDeclare).Default()).SetExchangeName("integration-test")
	err = container.Publish(
		"",
		"test-publish-subscribe",
		*new(OtherPublish).SetPersistent().SetBody([]byte("test payload")),
	)
	assert.Nil(t, err)

	var result string
	testHandler := func(ch *amqp.Channel, msg amqp.Delivery) {
		msg.Ack(true)
		result = string(msg.Body)
	}
	consumer := container.Consumer()
	err = consumer.
		SetTopic("test-publish-subscribe").
		SetQueueDeclare(consumer.GetQueueDeclare().SetAutoDelete(true)).
		Consume(1, testHandler)
	assert.Nil(t, err)

	time.Sleep(1 * time.Second)
	assert.Equal(t, "test payload", result)
}
