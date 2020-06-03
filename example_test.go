package rabbitmqclient

import (
	"log"
	"time"

	"github.com/panjf2000/ants/v2"
	"github.com/streadway/amqp"
)

func ExampleContainer() {
	container, err := testSetup.NewContainer()
	if err != nil {
		log.Panicln(err)
	}

	ants.Submit(func() {
		err = container.
			SetExchange(new(ExchangeDeclare).Default()).
			SetExchangeName("integration-test").
			Publish(
				"",
				"test-normal",
				*new(OtherPublish).SetPersistent().SetBody([]byte("test payload")),
			)
		if err != nil {
			log.Panicln(err)
		}
	})

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
	if err != nil {
		log.Panicln(err)
	}

	time.Sleep(200 * time.Millisecond)
	if result != "test payload" {
		log.Panicf("Expected: %v Actual: %v doesn't match.", "test payload", result)
	}
}
