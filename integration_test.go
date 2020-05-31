package rabbitmqclient

import (
	"testing"
	"time"

	"github.com/streadway/amqp"
	"github.com/stretchr/testify/assert"
)

const (
	ErrorFormat = "Error: %+v"
)

func TestPublishAndSubscribe(t *testing.T) {
	container, err := NewContainer(testSetup.GetConnection())
	assert.Nil(t, err)

	container.SetExchange(new(ExchangeDeclare).Default()).SetExchangeName("integration-test")
	err = container.Publish("", "integration-test", *new(OtherPublish).SetPersistent().SetBody([]byte("test payload")))
	assert.Nil(t, err)

	var result string
	testHandler := func(ch *amqp.Channel, msg amqp.Delivery) {
		result = string(msg.Body)
	}
	err = container.Consumer().SetTopic("integration-test").Consume(1, testHandler)
	assert.Nil(t, err)

	time.Sleep(1 * time.Second)
	assert.Equal(t, "test payload", result)
}
