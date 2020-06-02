package rabbitmqclient

import (
	"testing"
	"time"

	"github.com/streadway/amqp"
	"github.com/stretchr/testify/assert"
)

func TestPublish(t *testing.T) {
	t.Run("Publish: Empty Topic", func(t *testing.T) {
		container, err := testSetup.NewContainer()
		assert.Nil(t, err)

		err = container.Publish(
			"",
			"",
			*new(OtherPublish).SetBody([]byte("test payload")),
		)
		assert.Nil(t, err)
	})

	t.Run("Publish: Error Topology Init", func(t *testing.T) {
		container, err := testSetup.NewContainer()
		assert.Nil(t, err)

		container.SetTopology(
			container.GetTopology().
				AddQueueUnbind(QueueUnbind{QueueBind: QueueBind{Name: "publish-no-unbind"}}),
		)

		err = container.Publish(
			"",
			"",
			*new(OtherPublish).SetBody([]byte("test payload")),
		)
		assert.NotNil(t, err)
	})
}

func TestPublishSubscribe(t *testing.T) {
	t.Run("PublishSubscribe: Normal Case", func(t *testing.T) {
		container, err := testSetup.NewContainer()
		assert.Nil(t, err)

		container.SetExchange(new(ExchangeDeclare).Default()).SetExchangeName("integration-test")
		err = container.Publish(
			"",
			"test-normal",
			*new(OtherPublish).SetPersistent().SetBody([]byte("test payload")),
		)
		assert.Nil(t, err)

		var result string
		testHandler := func(ch *amqp.Channel, msg amqp.Delivery) {
			msg.Ack(false)
			result = string(msg.Body)
		}
		consumer := container.Consumer()
		err = consumer.
			SetTopic("test-normal").
			SetQueueDeclare(consumer.GetQueueDeclare()).
			Consume(0, testHandler)
		assert.Nil(t, err)

		time.Sleep(200 * time.Millisecond)
		assert.Equal(t, "test payload", result)
	})
}
