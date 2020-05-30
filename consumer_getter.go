package rabbitmqclient

func (c *Consumer) getTopic() string {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	return c.bind.Key
}

func (c *Consumer) getChannelKey() string {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	return c.channelKey
}

// GetQueueDeclare get queue declaration for this consumer.
func (c *Consumer) GetQueueDeclare() *QueueDeclare {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	return c.declare
}

// GetQueueBind gets the queue binder for this consumer.
func (c *Consumer) GetQueueBind() *QueueBind {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	return c.bind
}
