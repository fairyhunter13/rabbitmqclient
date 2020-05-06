package rabbitmqclient

import (
	"sync"
	"time"

	"github.com/fairyhunter13/rabbitmqclient/args"
)

// Topology contains all declarations needed to define the topology in the rabbitmq.
type Topology struct {
	mutex *sync.RWMutex
	// Mutex protects the following fields
	// exchange topology
	exchangeDeclare        []args.ExchangeDeclare
	exchangeDeclarePassive []args.ExchangeDeclarePassive
	exchangeBind           []args.ExchangeBind
	exchangeUnbind         []args.ExchangeUnbind
	exchangeDelete         []args.ExchangeDelete

	// queue topology
	queueDeclare        []args.QueueDeclare
	queueDeclarePassive []args.QueueDeclarePassive
	queueBind           []args.QueueBind
	queueDelete         []args.QueueDelete
	queueUnbind         []args.QueueUnbind

	currentTime *time.Time
	lastTime    *time.Time
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
