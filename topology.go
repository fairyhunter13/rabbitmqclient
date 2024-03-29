package rabbitmqclient

import (
	"reflect"
	"sync"
	"time"
)

// Topology contains all declarations needed to define the topology topology in the rabbitmq.
type Topology struct {
	topoVal reflect.Value
	mutex   sync.RWMutex
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

	currentTime time.Time
	lastTime    time.Time
}

// NewTopology creates a new topology
func NewTopology() *Topology {
	now := time.Now()
	topo := &Topology{
		currentTime: now,
		lastTime:    now,
	}
	topo.topoVal = reflect.ValueOf(topo)
	return topo
}

func (t *Topology) update() *Topology {
	t.mutex.Lock()
	t.currentTime = time.Now().Add(1 * time.Millisecond)
	t.mutex.Unlock()
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
	result = t.currentTime.After(t.lastTime)
	if result {
		t.syncTime()
	}
	return
}
