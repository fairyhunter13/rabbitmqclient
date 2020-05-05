package rabbitmqclient

import "github.com/fairyhunter13/rabbitmqclient/args"

func (c *Container) setDefaultExchange() {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	if c.globalExchange == nil {
		c.globalExchange = new(args.ExchangeDeclare).Default()
	}
}

func (c *Container) setDefaultTopology() {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	if c.Topology == nil {
		c.Topology = NewTopology()
	}
}
