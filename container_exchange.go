package rabbitmqclient

// SetExchange sets the exchange of the global exchange.
func (c *Container) SetExchange(exc *ExchangeDeclare) *Container {
	if exc != nil {
		c.mutex.Lock()
		c.exchange = exc
		c.mutex.Unlock()
	}
	return c
}

// SetExchangeName sets the exchange name of the global exchange.
func (c *Container) SetExchangeName(name string) *Container {
	if name != "" {
		c.mutex.Lock()
		c.exchange.SetName(name)
		c.mutex.Unlock()
	}
	return c
}

// GetGlobalExchange get the exchange declare inside the global exchange.
func (c *Container) GetGlobalExchange() *ExchangeDeclare {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	return c.exchange
}
