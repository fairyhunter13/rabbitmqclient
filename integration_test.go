package rabbitmqclient

import (
	"testing"
	"time"

	"github.com/panjf2000/ants/v2"
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

	t.Run("Publish: Heavy Publish Multi Goroutine", func(t *testing.T) {
		container, err := testSetup.NewContainer()
		assert.Nil(t, err)

		for index := 0; index <= 1000; index++ {
			ants.Submit(func() {
				err = container.Publish(
					"",
					"",
					*new(OtherPublish).SetBody([]byte("test payload")),
				)
				assert.Nil(t, err)
			})
		}
	})

	t.Run("Publish: Heavy Publish Single Goroutine", func(t *testing.T) {
		container, err := testSetup.NewContainer()
		assert.Nil(t, err)

		for index := 0; index <= 1000; index++ {
			err = container.Publish(
				"",
				"",
				*new(OtherPublish).SetBody([]byte("test payload")),
			)
			assert.Nil(t, err)
		}
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

func TestRabbitMQNetwork(t *testing.T) {
	t.Run("RabbitMQNetwork: Init all network architecture", func(t *testing.T) {
		container, err := testSetup.NewContainer()
		assert.Nil(t, err)

		exchangeDeclare := new(ExchangeDeclare).Default()
		exchangeDeclareTest := new(ExchangeDeclare).Default().SetName("test")
		exchangeDeclareQueue := new(ExchangeDeclare).Default().SetName("test-queue")
		exchangeBind := new(ExchangeBind).
			SetSource(exchangeDeclare.Name).
			SetKey("test").
			SetDestination(exchangeDeclareTest.Name)
		exchangeDelete := new(ExchangeDelete).
			SetName(exchangeDeclare.Name)
		exchangeDeleteTest := new(ExchangeDelete).
			SetName(exchangeDeclareTest.Name)

		queueDeclare := new(QueueDeclare).Default()
		queueBind := new(QueueBind).Default().SetExchange(exchangeDeclareQueue.Name)
		queueDelete := new(QueueDelete).SetName(queueDeclare.Name)

		topology := NewTopology().
			// Exchange section
			AddExchangeDeclare(*exchangeDeclare).
			AddExchangeDeclare(*exchangeDeclareTest).
			AddExchangeDeclare(*exchangeDeclareQueue).
			AddExchangeDeclarePassive(ExchangeDeclarePassive{*exchangeDeclare}).
			AddExchangeDeclarePassive(ExchangeDeclarePassive{*exchangeDeclareTest}).
			AddExchangeBind(*exchangeBind).
			AddExchangeUnbind(ExchangeUnbind{*exchangeBind}).
			AddExchangeDelete(*exchangeDelete).
			AddExchangeDelete(*exchangeDeleteTest).
			// Queue section
			AddQueueDeclare(*queueDeclare).
			AddQueueDeclarePassive(QueueDeclarePassive{*queueDeclare}).
			AddQueueBind(*queueBind).
			AddQueueUnbind(QueueUnbind{*queueBind}).
			AddQueueDelete(*queueDelete)

		err = container.SetTopology(topology).Init()
		assert.Nil(t, err)
	})
}

func TestConnectionClose(t *testing.T) {
	// TODO: add connection close testing
}
