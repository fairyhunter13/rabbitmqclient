package rabbitmqclient

import "sync"

type queueBinder struct {
	mutex   *sync.RWMutex
	declare *QueueDeclare
	bind    *QueueBind
}

func newQueueBinder() *queueBinder {
	return &queueBinder{
		mutex:   new(sync.RWMutex),
		declare: new(QueueDeclare).Default(),
		bind:    new(QueueBind).Default(),
	}
}

func (q *queueBinder) SetQueueDeclare(declare *QueueDeclare) *queueBinder {
	if declare != nil {
		q.mutex.Lock()
		q.declare = declare
		q.mutex.Unlock()
	}
	return q
}

func (q *queueBinder) SetQueueBind(bind *QueueBind) *queueBinder {
	if bind != nil {
		q.mutex.Lock()
		q.bind = bind
		q.mutex.Unlock()
	}
	return q
}
