package rabbitmqclient

import (
	"sync"
	"time"
)

// Topology contains all declarations needed to define the topology in the rabbitmq.
type Topology struct {
	mutex *sync.RWMutex
	// Mutex protects the following fields
	exchangeDeclareArgs []ExchangeDeclareArgs
	queueDeclareArgs    []QueueDeclareArgs
	queueBindArgs       []QueueBindArgs
	currentTime         *time.Time
	lastTime            *time.Time
}

// NewTopology creates a new topology
func NewTopology() *Topology {
	now := time.Now()
	return &Topology{
		mutex:       new(sync.RWMutex),
		currentTime: &now,
		lastTime:    &now,
	}
}

// Update updates the last time of the topology.
func (t *Topology) Update() *Topology {
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

// AddExchangeDeclare add the exchange declare args to the topology.
func (t *Topology) AddExchangeDeclare(args ExchangeDeclareArgs) *Topology {
	t.mutex.Lock()
	defer t.mutex.Unlock()
	t.exchangeDeclareArgs = append(t.exchangeDeclareArgs, args)
	return t
}

// GetExchangeDeclare return the exchange declare args inside the topology.
func (t *Topology) GetExchangeDeclare() []ExchangeDeclareArgs {
	t.mutex.RLock()
	defer t.mutex.RUnlock()
	return append(t.exchangeDeclareArgs[:0:0], t.exchangeDeclareArgs...)
}

// AddQueueDeclare adds the queue declaration into the topology
func (t *Topology) AddQueueDeclare(args QueueDeclareArgs) *Topology {
	t.mutex.Lock()
	defer t.mutex.Unlock()
	t.queueDeclareArgs = append(t.queueDeclareArgs, args)
	return t
}

// GetQueueDeclare gets the queue declaration inside the topology.
func (t *Topology) GetQueueDeclare() []QueueDeclareArgs {
	t.mutex.RLock()
	defer t.mutex.RUnlock()
	return append(t.queueDeclareArgs[:0:0], t.queueDeclareArgs...)
}

// AddQueueBind adds the queue bind args to the topology
func (t *Topology) AddQueueBind(args QueueBindArgs) *Topology {
	t.mutex.Lock()
	defer t.mutex.Unlock()
	t.queueBindArgs = append(t.queueBindArgs, args)
	return t
}

// GetQueueBind gets the queue bind args inside the topology
func (t *Topology) GetQueueBind() []QueueBindArgs {
	t.mutex.RLock()
	defer t.mutex.RUnlock()
	return append(t.queueBindArgs[:0:0], t.queueBindArgs...)
}
