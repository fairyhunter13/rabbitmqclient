package rabbitmqclient

// SetQueueName sets the queue name of the consumer.
// Call SetQueueName after you call the SetTopic method
// in order to prevent the override of the queue name.
func (c *Consumer) SetQueueName(withPrefix bool, name string) *Consumer {
	if name == "" {
		name = c.getTopic()
	}
	newQueueName := GenerateQueueName(withPrefix, name)
	c.mutex.Lock()
	c.declare.SetName(newQueueName)
	c.bind.SetName(newQueueName)
	c.mutex.Unlock()
	return c
}

// SetChannelKey sets the channel key to keep track inside the IConnectionManager.
// Call this function after you call the SetTopic method if you want to use
// the default key for the channel key. The default key is the topic name.
func (c *Consumer) SetChannelKey(withPrefix bool, name string) *Consumer {
	if name == "" {
		name = c.getTopic()
	}
	c.mutex.Lock()
	c.channelKey = GenerateConsumerChannelKey(withPrefix, name)
	c.mutex.Unlock()
	return c
}

// SetExchangeName sets the name of exchange to be binded to the queue of the consumer.
func (c *Consumer) SetExchangeName(name string) *Consumer {
	if name != "" {
		c.mutex.Lock()
		c.bind.SetExchange(name)
		c.mutex.Unlock()
	}
	return c
}

// SetTopic sets the topic of the queue to the exchange.
// Call SetTopic before you call the SetQueueName method
// in order not to override the current queue name.
func (c *Consumer) SetTopic(topic string) *Consumer {
	if topic != "" {
		c.mutex.Lock()
		c.bind.SetKey(topic)
		c.mutex.Unlock()
		c.SetQueueName(true, "")
		c.SetChannelKey(true, "")
	}
	return c
}

// SetConsume sets the consume args for consumer.
// The args list can be seen inside the amqp documentation.
func (c *Consumer) SetConsume(consume *Consume) *Consumer {
	if consume != nil {
		c.mutex.Lock()
		c.consume = consume
		c.mutex.Unlock()
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
