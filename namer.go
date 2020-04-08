package rabbitmqclient

import "fmt"

// List of default constants for this library.
const (
	DefaultPrefixExchange = "amqp."
	DefaultQueue          = "default"
	DefaultKey            = "default"
)

// List of all exchange type in rabbitmq
const (
	TypeDirect  = `direct`
	TypeFanout  = `fanout`
	TypeTopic   = `topic`
	TypeHeaders = `headers`
)

// GenerateExchangeName generates the default exchange name of this library.
func GenerateExchangeName(isPrefix bool, name string) string {
	if isPrefix {
		return fmt.Sprintf("%s%s", DefaultPrefixExchange, name)
	}
	return fmt.Sprintf("%s", name)
}
