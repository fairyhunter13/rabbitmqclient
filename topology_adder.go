package rabbitmqclient

import "github.com/fairyhunter13/rabbitmqclient/args"

// AddExchangeDeclare add the exchange declare args to the topology.
func (t *Topology) AddExchangeDeclare(arg args.ExchangeDeclare) *Topology {
	t.mutex.Lock()
	t.exchangeDeclare = append(t.exchangeDeclare, arg)
	t.mutex.Unlock()
	t.update()
	return t
}

// AddExchangeDeclarePassive adds the exchange declare if it is not available.
func (t *Topology) AddExchangeDeclarePassive(arg args.ExchangeDeclarePassive) *Topology {
	t.mutex.Lock()
	t.exchangeDeclarePassive = append(t.exchangeDeclarePassive, arg)
	t.mutex.Unlock()
	t.update()
	return t
}

// AddExchangeBind adds the exchange bind to the topology.
func (t *Topology) AddExchangeBind(arg args.ExchangeBind) *Topology {
	t.mutex.Lock()
	t.exchangeBind = append(t.exchangeBind, arg)
	t.mutex.Unlock()
	t.update()
	return t
}

// AddExchangeUnbind adds the exchange unbind to the topology.
func (t *Topology) AddExchangeUnbind(arg args.ExchangeUnbind) *Topology {
	t.mutex.Lock()
	t.exchangeUnbind = append(t.exchangeUnbind, arg)
	t.mutex.Unlock()
	t.update()
	return t
}

// AddExchangeDelete adds the exchange delete to the topology.
func (t *Topology) AddExchangeDelete(arg args.ExchangeDelete) *Topology {
	t.mutex.Lock()
	t.exchangeDelete = append(t.exchangeDelete, arg)
	t.mutex.Unlock()
	t.update()
	return t
}

// AddQueueDeclare adds the queue declaration into the topology
func (t *Topology) AddQueueDeclare(arg args.QueueDeclare) *Topology {
	t.mutex.Lock()
	t.queueDeclare = append(t.queueDeclare, arg)
	t.mutex.Unlock()
	t.update()
	return t
}

// AddQueueDeclarePassive adds the queue passive declaration into the topology
func (t *Topology) AddQueueDeclarePassive(arg args.QueueDeclarePassive) *Topology {
	t.mutex.Lock()
	t.queueDeclarePassive = append(t.queueDeclarePassive, arg)
	t.mutex.Unlock()
	t.update()
	return t
}

// AddQueueBind adds the queue bind args to the topology
func (t *Topology) AddQueueBind(arg args.QueueBind) *Topology {
	t.mutex.Lock()
	t.queueBind = append(t.queueBind, arg)
	t.mutex.Unlock()
	t.update()
	return t
}

// AddQueueUnbind adds the queue unbind args to the topology
func (t *Topology) AddQueueUnbind(arg args.QueueUnbind) *Topology {
	t.mutex.Lock()
	t.queueUnbind = append(t.queueUnbind, arg)
	t.mutex.Unlock()
	t.update()
	return t
}

// AddQueueDelete adds the queue delete args to the topology
func (t *Topology) AddQueueDelete(arg args.QueueDelete) *Topology {
	t.mutex.Lock()
	t.queueDelete = append(t.queueDelete, arg)
	t.mutex.Unlock()
	t.update()
	return t
}
