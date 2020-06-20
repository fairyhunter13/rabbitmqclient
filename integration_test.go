package rabbitmqclient

import (
	"sync"
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

	t.Run("Publish: Heavy Publish Multi Goroutine", func(t *testing.T) {
		container, err := testSetup.NewContainer()
		assert.Nil(t, err)

		for index := 0; index <= 1000; index++ {
			go func() {
				err := container.Publish(
					"",
					"",
					*new(OtherPublish).SetBody([]byte("test payload")),
				)
				assert.Nil(t, err)
			}()
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

type TestChecker struct {
	mutex  sync.RWMutex
	Result string
}

func (t *TestChecker) SetResult(result string) *TestChecker {
	t.mutex.Lock()
	t.Result = result
	t.mutex.Unlock()
	return t
}

func (t *TestChecker) GetResult() string {
	t.mutex.RLock()
	defer t.mutex.RUnlock()
	return t.Result
}

func TestPublishSubscribe(t *testing.T) {
	t.Run("PublishSubscribe: Normal Case", func(t *testing.T) {
		container, err := testSetup.NewContainer()
		assert.Nil(t, err)

		container.SetExchange(new(ExchangeDeclare).Default()).SetExchangeName("integration-test")

		go func() {
			err := container.Publish(
				"",
				"test-normal",
				*new(OtherPublish).SetPersistent().SetBody([]byte("test normal payload")),
			)
			assert.Nil(t, err)
		}()

		testCheck := new(TestChecker)
		testHandler := func(ch *amqp.Channel, msg amqp.Delivery) {
			msg.Ack(true)
			testCheck.SetResult(string(msg.Body))
		}

		consumer := container.Consumer()
		err = consumer.
			SetTopic("test-normal").
			SetQueueDeclare(consumer.GetQueueDeclare()).
			Consume(0, testHandler)
		assert.Nil(t, err)

		time.Sleep(normalTimeSleep)
		assert.Equal(t, "test normal payload", testCheck.GetResult())
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
	t.Run("ConnectionClosed: Publish Error", func(t *testing.T) {
		container, err := testSetup.NewContainerAndConnection()
		assert.Nil(t, err)

		err = container.Close()
		assert.Nil(t, err)

		err = container.Publish("", "", *new(OtherPublish).SetBody([]byte("test publih")))
		assert.Equal(t, ErrConnectionAlreadyClosed, err)
	})

	t.Run("ConnectionClosed: Consume Error", func(t *testing.T) {
		container, err := testSetup.NewContainerAndConnection()
		assert.Nil(t, err)

		err = container.Close()
		assert.Nil(t, err)

		testHandler := func(ch *amqp.Channel, msg amqp.Delivery) {
			return
		}

		err = container.Consumer().Consume(0, testHandler)
		assert.Equal(t, ErrConnectionAlreadyClosed, err)
	})
}
