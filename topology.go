package rabbitmqclient

import (
	"reflect"
	"sync"
	"time"

	"github.com/streadway/amqp"
)

// Topology contains all declarations needed to define the topology topology in the rabbitmq.
type Topology struct {
	topoVal reflect.Value
	mutex   *sync.RWMutex
	// Mutex protects the following fields
	// exchange topology
	exchangeDeclare        []ExchangeDeclare
	exchangeDeclarePassive []ExchangeDeclarePassive
	exchangeBind           []ExchangeBind
	exchangeUnbind         []ExchangeUnbind
	exchangeDelete         []ExchangeDelete

	// queue topology
	queueDeclare        []QueueDeclare
	queueDeclarePassive []QueueDeclarePassive
	queueBind           []QueueBind
	queueDelete         []QueueDelete
	queueUnbind         []QueueUnbind

	currentTime *time.Time
	lastTime    *time.Time
}

// NewTopology creates a new topology
func NewTopology() *Topology {
	now := time.Now()
	topo := &Topology{
		mutex:       new(sync.RWMutex),
		currentTime: &now,
		lastTime:    &now,
	}
	topo.topoVal = reflect.ValueOf(topo)
	return topo
}

func (t *Topology) update() *Topology {
	t.mutex.Lock()
	defer t.mutex.Unlock()
	*t.currentTime = time.Now()
	return t
}

func (t *Topology) syncTime() {
	t.lastTime = t.currentTime
}

// IsUpdated checks if the topology has been updated or not.
// IsUpdated also automatically sync the time of last time to the current time if it is updated.
func (t *Topology) IsUpdated() (result bool) {
	t.mutex.Lock()
	defer t.mutex.Unlock()
	result = t.currentTime.After(*t.lastTime)
	if result {
		t.syncTime()
	}
	return
}

var topologyDeclarationList = []string{
	TopologyExchangeDeclare,
	TopologyExchangeDeclarePassive,
	TopologyExchangeBind,
	TopologyExchangeUnbind,
	TopologyExchangeDelete,
	TopologyQueueDeclare,
	TopologyQueueDeclarePassive,
	TopologyQueueBind,
	TopologyQueueUnbind,
	TopologyQueueDelete,
}

// DeclareAll declares all topologies available inside the topology.
func (t *Topology) DeclareAll(ch *amqp.Channel) (err error) {
	if ch == nil {
		err = ErrArgumentsMusntBeEmpty
		return
	}
	for _, declaration := range topologyDeclarationList {
		err = t.Declare(ch, declaration)
		if err != nil {
			return
		}
	}
	return
}

// Declare declares the topology declaration based on the input.
func (t *Topology) Declare(ch *amqp.Channel, declaration string) (err error) {
	if ch == nil || declaration == "" {
		err = ErrArgumentsMusntBeEmpty
		return
	}
	method := t.topoVal.MethodByName("Get" + declaration)
	if method.IsZero() {
		err = ErrMethodNotFound
		return
	}
	returnArr := method.Call([]reflect.Value{})
	if len(returnArr) != 1 {
		err = ErrInvalidFunctionCalled
		return
	}
	kind := returnArr[0].Kind()
	if kind != reflect.Array && kind != reflect.Slice {
		err = ErrInvalidReturnValues
		return
	}
	err = t.iterateDeclare(ch, returnArr[0].Interface())
	return
}

func (t *Topology) iterateDeclare(ch *amqp.Channel, list interface{}) (err error) {
	val := reflect.ValueOf(list)
	if !val.IsValid() {
		return
	}
	switch val.Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < val.Len(); i++ {
			err = t.declare(ch, val.Index(i))
			if err != nil {
				return
			}
		}
	}
	return
}

func (t *Topology) declare(ch *amqp.Channel, declaration interface{}) (err error) {
	switch elem := declaration.(type) {
	case ExchangeDeclare:
		err = ch.ExchangeDeclare(elem.Name, elem.Kind, elem.Durable, elem.AutoDelete, elem.Internal, elem.NoWait, elem.Args)
	case ExchangeDeclarePassive:
		err = ch.ExchangeDeclarePassive(elem.Name, elem.Kind, elem.Durable, elem.AutoDelete, elem.Internal, elem.NoWait, elem.Args)
	case ExchangeBind:
		err = ch.ExchangeBind(elem.Destination, elem.Key, elem.Source, elem.NoWait, elem.Args)
	case ExchangeUnbind:
		err = ch.ExchangeUnbind(elem.Destination, elem.Key, elem.Source, elem.NoWait, elem.Args)
	case ExchangeDelete:
		err = ch.ExchangeDelete(elem.Name, elem.IfUnused, elem.NoWait)
	case QueueDeclare:
		_, err = ch.QueueDeclare(elem.Name, elem.Durable, elem.AutoDelete, elem.Exclusive, elem.NoWait, elem.Args)
	case QueueDeclarePassive:
		_, err = ch.QueueDeclarePassive(elem.Name, elem.Durable, elem.AutoDelete, elem.Exclusive, elem.NoWait, elem.Args)
	case QueueBind:
		err = ch.QueueBind(elem.Name, elem.Key, elem.Exchange, elem.NoWait, elem.Args)
	case QueueUnbind:
		err = ch.QueueUnbind(elem.Name, elem.Key, elem.Exchange, elem.Args)
	case QueueDelete:
		_, err = ch.QueueDelete(elem.Name, elem.IfUnused, elem.IfEmpty, elem.NoWait)
	}
	return
}
