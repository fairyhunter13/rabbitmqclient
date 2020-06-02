package rabbitmqclient

import "github.com/fairyhunter13/amqpwrapper"

// List of all uint constants for boolean
const (
	FalseUint uint64 = iota
	TrueUint
)

// List of default config used for this library.
const (
	// prefixes
	DefaultPrefixExchange = "amqp."
	DefaultPrefixQueue    = "queue."
	DefaultPrefixConsumer = "consumer."

	DefaultQueue      = "default"
	DefaultTopic      = "default"
	DefaultChannelKey = "default"
)

// List of default variable configuration for this library
var (
	DefaultExchange = GenerateExchangeName(true, TypeDirect)
)

const (
	// DefaultKeyInitiator sets the key channel for the initiator of topology in the producer connection.
	DefaultKeyInitiator = "initiator"
	// DefaultKeyProducer specifies the key channel for producer.
	DefaultKeyProducer = "producer.%d"
)

// List of channel types for amqpwrapper
const (
	// DefaultTypeProducer specifies the default type of channel for the amqpwrapper.
	DefaultTypeProducer = amqpwrapper.Producer
	// DefaultTypeConsumer specifies the default type of channel for the amqpwrapper.
	DefaultTypeConsumer = amqpwrapper.Consumer
)

// List of all exchange type in rabbitmq
const (
	TypeDirect  = `direct`
	TypeFanout  = `fanout`
	TypeTopic   = `topic`
	TypeHeaders = `headers`
)

// List of all delivery modes for amqp rabbitmqclient
const (
	DeliveryModeTransient  = 1
	DeliveryModePersistent = 2
)

// List of all topology constants
const (
	TopologyExchangeDeclare        = `ExchangeDeclare`
	TopologyExchangeDeclarePassive = `ExchangeDeclarePassive`
	TopologyExchangeBind           = `ExchangeBind`
	TopologyExchangeUnbind         = `ExchangeUnbind`
	TopologyExchangeDelete         = `ExchangeDelete`
	TopologyQueueDeclare           = `QueueDeclare`
	TopologyQueueDeclarePassive    = `QueueDeclarePassive`
	TopologyQueueBind              = `QueueBind`
	TopologyQueueUnbind            = `QueueUnbind`
	TopologyQueueDelete            = `QueueDelete`
)
