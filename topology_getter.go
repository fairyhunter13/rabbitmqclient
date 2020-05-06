package rabbitmqclient

// GetExchangeDeclare return the exchange declare args inside the topology.
func (t *Topology) GetExchangeDeclare() []ExchangeDeclare {
	t.mutex.RLock()
	defer t.mutex.RUnlock()
	return append(t.exchangeDeclare[:0:0], t.exchangeDeclare...)
}

// GetExchangeDeclarePassive return the exchange declare args inside the topology.
func (t *Topology) GetExchangeDeclarePassive() []ExchangeDeclarePassive {
	t.mutex.RLock()
	defer t.mutex.RUnlock()
	return append(t.exchangeDeclarePassive[:0:0], t.exchangeDeclarePassive...)
}

// GetExchangeBind return the exchange bind inside the topology.
func (t *Topology) GetExchangeBind() []ExchangeBind {
	t.mutex.RLock()
	defer t.mutex.RUnlock()
	return append(t.exchangeBind[:0:0], t.exchangeBind...)
}

// GetExchangeUnbind return the exchange bind inside the topology.
func (t *Topology) GetExchangeUnbind() []ExchangeUnbind {
	t.mutex.RLock()
	defer t.mutex.RUnlock()
	return append(t.exchangeUnbind[:0:0], t.exchangeUnbind...)
}

// GetExchangeDelete return the exchange delete inside the topology.
func (t *Topology) GetExchangeDelete() []ExchangeDelete {
	t.mutex.RLock()
	defer t.mutex.RUnlock()
	return append(t.exchangeDelete[:0:0], t.exchangeDelete...)
}

// GetQueueDeclare gets the queue declaration inside the topology.
func (t *Topology) GetQueueDeclare() []QueueDeclare {
	t.mutex.RLock()
	defer t.mutex.RUnlock()
	return append(t.queueDeclare[:0:0], t.queueDeclare...)
}

// GetQueueDeclarePassive gets the queue passive declaration inside the topology.
func (t *Topology) GetQueueDeclarePassive() []QueueDeclarePassive {
	t.mutex.RLock()
	defer t.mutex.RUnlock()
	return append(t.queueDeclarePassive[:0:0], t.queueDeclarePassive...)
}

// GetQueueBind gets the queue bind args inside the topology
func (t *Topology) GetQueueBind() []QueueBind {
	t.mutex.RLock()
	defer t.mutex.RUnlock()
	return append(t.queueBind[:0:0], t.queueBind...)
}

// GetQueueUnbind gets the queue unbind args inside the topology
func (t *Topology) GetQueueUnbind() []QueueUnbind {
	t.mutex.RLock()
	defer t.mutex.RUnlock()
	return append(t.queueUnbind[:0:0], t.queueUnbind...)
}

// GetQueueDelete gets the queue delete args inside the topology
func (t *Topology) GetQueueDelete() []QueueDelete {
	t.mutex.RLock()
	defer t.mutex.RUnlock()
	return append(t.queueDelete[:0:0], t.queueDelete...)
}
