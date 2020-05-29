package rabbitmqclient

// SetQueueName sets the queue name of the consumer.
func (c *Consumer) SetQueueName(withPrefix bool, name string) *Consumer {
	var newQueueName string
	if name != "" {
		newQueueName = generateQueueName(withPrefix, name)
	} else {
		c.mutex.RLock()
		newQueueName = generateQueueName(withPrefix, c.bind.Key)
		c.mutex.RUnlock()
	}
	c.mutex.Lock()
	c.declare.Name = newQueueName
	c.bind.Name = newQueueName
	c.mutex.Unlock()
	return c
}

// SetExchangeName sets the name of exchange to be binded to the queue of the consumer.
func (c *Consumer) SetExchangeName(name string) *Consumer {
	if name != "" {
		c.mutex.Lock()
		c.bind.Exchange = name
		c.mutex.Unlock()
	}
	return c
}

// SetTopic sets the topic of the queue to the exchange.
func (c *Consumer) SetTopic(topic string) *Consumer {
	if topic != "" {
		c.mutex.Lock()
		c.bind.Key = topic
		c.mutex.Unlock()
		c.SetQueueName(true, "")
	}
	return c
}

// SetQueueDeclare sets the queue declaration topology for this Consumer.
func (c *Consumer) SetQueueDeclare(declare *QueueDeclare) *Consumer {
	if declare != nil {
		c.mutex.Lock()
		c.declare = declare
		c.mutex.Unlock()
	}
	return c
}

// SetQueueBind sets the queue bind topology for this Consumer.
func (c *Consumer) SetQueueBind(bind *QueueBind) *Consumer {
	if bind != nil {
		c.mutex.Lock()
		c.bind = bind
		c.mutex.Unlock()
	}
	return c
}
