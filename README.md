بِسْمِ اللّٰهِ الرَّحْمٰنِ الرَّحِيْمِ

<br/>

السَّلاَمُ عَلَيْكُمْ وَرَحْمَةُ اللهِ وَبَرَكَاتُهُ

<br/>

ٱلْحَمْدُ لِلَّهِ رَبِّ ٱلْعَٰلَمِينَ

ٱلْحَمْدُ لِلَّهِ رَبِّ ٱلْعَٰلَمِينَ

ٱلْحَمْدُ لِلَّهِ رَبِّ ٱلْعَٰلَمِينَ

<br/>

اللَّهُمَّ صَلِّ عَلَى مُحَمَّدٍ ، وَعَلَى آلِ مُحَمَّدٍ ، كَمَا صَلَّيْتَ عَلَى إِبْرَاهِيمَ وَعَلَى آلِ إِبْرَاهِيمَ ، إِنَّكَ حَمِيدٌ مَجِيدٌ ، اللَّهُمَّ بَارِكْ عَلَى مُحَمَّدٍ ، وَعَلَى آلِ مُحَمَّدٍ ، كَمَا بَارَكْتَ عَلَى إِبْرَاهِيمَ ، وَعَلَى آلِ إِبْرَاهِيمَ ، إِنَّكَ حَمِيدٌ مَجِيدٌ

# RabbitMQ Client
[![Coverage Status](https://coveralls.io/repos/github/fairyhunter13/rabbitmqclient/badge.svg?branch=master)](https://coveralls.io/github/fairyhunter13/rabbitmqclient?branch=master)
[![CircleCI](https://circleci.com/gh/fairyhunter13/amqpwrapper.svg?style=svg)](https://circleci.com/gh/fairyhunter13/amqpwrapper)
[![Go Report Card](https://goreportcard.com/badge/github.com/fairyhunter13/rabbitmqclient)](https://goreportcard.com/report/github.com/fairyhunter13/rabbitmqclient)
<a title="Doc for rabbitmqclient" target="_blank" href="https://pkg.go.dev/github.com/fairyhunter13/rabbitmqclient?tab=doc"><img src="https://img.shields.io/badge/go.dev-doc-007d9c?style=flat-square&logo=read-the-docs"></a>

The rabbitmq client is built using the [fairyhunter13/amqpwrapper](https://github.com/fairyhunter13/amqpwrapper) package.
This package can manage the topology of queue's network inside the RabbitMQ.
This package simplifies the management and usage of publishing and consuming for RabbitMQ.

# Example

This is an example how to use this package.

## Publish and Subscribe
This is an example of go code to publish and subscribe using this package.

```go
package main

import (
	"fmt"
	"log"
	"time"

	"github.com/fairyhunter13/amqpwrapper"
	"github.com/fairyhunter13/rabbitmqclient"
	"github.com/streadway/amqp"
)

func main() {
	uriHost := fmt.Sprintf("amqp://guest:guest@%s:5672", "localhost")
	conn, err := amqpwrapper.NewManager(uriHost, amqp.Config{})
	if err != nil {
		log.Panicln(err)
	}

	container, err := rabbitmqclient.NewContainer(conn)
	if err != nil {
		log.Panicln(err)
	}

	err = container.
		SetExchangeName("integration-test").
		Publish(
			"",
			"example",
			*new(rabbitmqclient.OtherPublish).
				SetPersistent().
				SetBody([]byte("test payload")),
		)
	if err != nil {
		log.Panicln(err)
	}

	var result string
	testHandler := func(ch *amqp.Channel, msg amqp.Delivery) {
		msg.Ack(false)
		result = string(msg.Body)
	}
	err = container.
		Consumer().
		SetTopic("example").
		Consume(0, testHandler)
	if err != nil {
		log.Panicln(err)
	}

	time.Sleep(2 * time.Second)
	if result != "test payload" {
		log.Panicf("Expected: %v Actual: %v doesn't match.", "test payload", result)
	}
}
```

## Topology
For this example, see this [test code](integration_test.go) inside `TestRabbitMQNetwork` function.

# Author
**fairyhunter13**

# License
The source code inside this package is available under the [MIT License](LICENSE).
