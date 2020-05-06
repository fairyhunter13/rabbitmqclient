package rabbitmqclient

func (c *Container) setDefaultExchange() {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	if c.globalExchange == nil {
		c.globalExchange = new(ExchangeDeclare).Default()
	}
}

func (c *Container) setDefaultTopology() {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	if c.Topology == nil {
		c.Topology = NewTopology()
	}
}
