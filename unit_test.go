package rabbitmqclient

import (
	"testing"

	"github.com/streadway/amqp"
	"github.com/stretchr/testify/assert"
)

func TestBuildArgs(t *testing.T) {
	t.Run("BuildArgs: Exchange", func(t *testing.T) {
		exchangeDeclare := new(ExchangeDeclare).
			SetName("test").
			SetKind(TypeFanout).
			SetDurable(true).
			SetAutoDelete(true).
			SetInternal(true).
			SetNoWait(true).
			SetArgs(amqp.Table{})
		assert.EqualValues(t, &ExchangeDeclare{
			Name:       "test",
			Kind:       TypeFanout,
			Durable:    true,
			AutoDelete: true,
			Internal:   true,
			NoWait:     true,
			Args:       amqp.Table{},
		}, exchangeDeclare)

		exchangeBind := new(ExchangeBind).
			SetDestination("test").
			SetKey("test").
			SetSource("test").
			SetNoWait(true).
			SetArgs(amqp.Table{})
		assert.EqualValues(t, &ExchangeBind{
			Destination: "test",
			Key:         "test",
			Source:      "test",
			NoWait:      true,
			Args:        amqp.Table{},
		}, exchangeBind)

		exchangeDelete := new(ExchangeDelete).
			SetName("test").
			SetIfUnused(true).
			SetNoWait(true)
		assert.EqualValues(t, &ExchangeDelete{
			Name:     "test",
			IfUnused: true,
			NoWait:   true,
		}, exchangeDelete)
	})

	t.Run("BuildArgs: Queue", func(t *testing.T) {
		queueDeclare := new(QueueDeclare).
			SetArgs(amqp.Table{}).
			SetAutoDelete(true).
			SetDurable(true).
			SetExclusive(true).
			SetName("test").
			SetNoWait(true)
		assert.EqualValues(t, &QueueDeclare{
			Name:       "test",
			Args:       amqp.Table{},
			AutoDelete: true,
			Durable:    true,
			Exclusive:  true,
			NoWait:     true,
		}, queueDeclare)

		queueBind := new(QueueBind).
			SetArgs(amqp.Table{}).
			SetExchange("test").
			SetKey("test").
			SetName("test").
			SetNoWait(true)
		assert.EqualValues(t, &QueueBind{
			Exchange: "test",
			Key:      "test",
			Name:     "test",
			Args:     amqp.Table{},
			NoWait:   true,
		}, queueBind)

		queueDelete := new(QueueDelete).
			SetIfEmpty(true).
			SetIfUnused(true).
			SetName("test").
			SetNoWait(true)
		assert.EqualValues(t, &QueueDelete{
			IfEmpty:  true,
			IfUnused: true,
			Name:     "test",
			NoWait:   true,
		}, queueDelete)
	})

	t.Run("BuildArgs: General", func(t *testing.T) {
		publish := new(Publish).
			SetExchange("test").
			SetKey("test").
			SetOtherPublish(
				*new(OtherPublish).
					SetBody([]byte("test")).
					SetContentEncoding("test").
					SetContentType("test").
					SetExpiration("test").
					SetHeaders(amqp.Table{}).
					SetImmediate(true).
					SetMandatory(true).
					SetPersistent().
					SetPriority(1).
					SetReplyTo("test"),
			)
		assert.EqualValues(t, &Publish{
			Exchange: "test",
			Key:      "test",
			OtherPublish: OtherPublish{
				Immediate: true,
				Mandatory: true,
				Msg: amqp.Publishing{
					Headers:         amqp.Table{},
					ContentEncoding: "test",
					ContentType:     "test",
					Body:            []byte("test"),
					Expiration:      "test",
					DeliveryMode:    DeliveryModePersistent,
					Priority:        1,
					ReplyTo:         "test",
				},
			},
		}, publish)

		publish.SetMsg(amqp.Publishing{})
		assert.EqualValues(t, amqp.Publishing{}, publish.Msg)

		consume := new(Consume).
			SetArgs(amqp.Table{}).
			SetAutoAck(true).
			SetConsumer("test").
			SetExclusive(true).
			SetNoLocal(true).
			SetNoWait(true).
			SetQueue("test")
		assert.EqualValues(t, &Consume{
			Args:      amqp.Table{},
			AutoAck:   true,
			Consumer:  "test",
			Exclusive: true,
			NoLocal:   true,
			NoWait:    true,
			Queue:     "test",
		}, consume)
	})
}

func TestConsumer(t *testing.T) {
	t.Run("Consumer: Variable Struct", func(t *testing.T) {
		container, err := testSetup.NewContainer()
		assert.Nil(t, err)

		consumer := newConsumer(container)
		consumer.SetExchangeName("test")
		assert.Equal(t, "test", consumer.GetQueueBind().Exchange)

		consumer.SetQueueBind(&QueueBind{
			Exchange: "test",
		})
		assert.Equal(t, &QueueBind{Exchange: "test"}, consumer.GetQueueBind())

		consumer.SetConsume(&Consume{
			Queue: "test",
		})
		assert.Equal(t, &Consume{Queue: "test"}, consumer.getConsume())
	})

	t.Run("Consumer: Nil Handler", func(t *testing.T) {
		container, err := testSetup.NewContainer()
		assert.Nil(t, err)

		consumer := newConsumer(container)
		err = consumer.Consume(0, nil)
		assert.NotNil(t, err)
	})
}

func TestContainer(t *testing.T) {
	t.Run("Container: Init Topology", func(t *testing.T) {
		container, err := testSetup.NewContainer()
		assert.Nil(t, err)

		err = container.Init()
		assert.Nil(t, err)

		topology := NewTopology().AddQueueDeclare(QueueDeclare{Name: "test", Durable: true}).AddQueueDelete(QueueDelete{Name: "test"})
		container.SetTopology(topology)
		assert.EqualValues(t, topology, container.GetTopology())

		err = container.Init()
		assert.Nil(t, err)

		ch, err := container.GetInitiatorChannel()
		assert.Nil(t, err)
		assert.NotNil(t, ch)
	})

	t.Run("Container: Nil Connection", func(t *testing.T) {
		// Connection can't be nil
		container, err := NewContainer(nil)
		assert.Nil(t, container)
		assert.NotNil(t, err)
	})

	t.Run("Container: QueueBind Error", func(t *testing.T) {
		container, err := testSetup.NewContainer()
		assert.Nil(t, err)

		// Init topology error: queue to bind is not found
		topology := NewTopology().AddQueueBind(QueueBind{Name: "not-found"})
		container.SetTopology(topology)
		err = container.Init()
		assert.NotNil(t, err)
	})
}

func TestGenerator(t *testing.T) {
	TopologyInitializationFn(nil)(new(amqp.Channel))

	ch, err := testSetup.GetConnection().CreateChannel(DefaultTypeProducer)
	assert.Nil(t, err)
	defer ch.Close()
	QosSetterFn(-1)(ch)

	assert.Equal(t, "test", GenerateExchangeName(false, "test"))
	assert.Equal(t, "test", GenerateConsumerChannelKey(false, "test"))
	assert.Equal(t, "test", GenerateQueueName(false, "test"))
}
