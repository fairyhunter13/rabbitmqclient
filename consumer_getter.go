package rabbitmqclient

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

func (c *Consumer) getContainer() *Container {
	return c.container
}

func (c *Consumer) getQueue() string {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	return c.declare.Name
}

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

func (c *Consumer) getConsume() *Consume {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	return c.consume
}
