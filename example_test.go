package rabbitmqclient

import (
	"log"
	"time"

	"github.com/streadway/amqp"
)

func ExampleContainer() {
	container, err := testSetup.NewContainer()
	if err != nil {
		log.Panicln(err)
	}

	go func() {
		err := container.
			SetExchangeName("integration-test").
			Publish(
				"",
				"test-normal",
				*new(OtherPublish).
					SetPersistent().
					SetBody([]byte("test payload")),
			)
		if err != nil {
			log.Panicln(err)
		}
	}()

	var result string
	testHandler := func(ch *amqp.Channel, msg amqp.Delivery) {
		msg.Ack(false)
		result = string(msg.Body)
	}
	err = container.
		Consumer().
		SetTopic("test-normal").
		Consume(0, testHandler)
	if err != nil {
		log.Panicln(err)
	}

	time.Sleep(normalTimeSleep)
	if result != "test payload" {
		log.Panicf("Expected: %v Actual: %v doesn't match.", "test payload", result)
	}
}
