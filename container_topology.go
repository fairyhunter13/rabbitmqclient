package rabbitmqclient

// SetTopology sets the current topology of this container.
func (c *Container) SetTopology(topo *Topology) *Container {
	c.setDefaultTopology()
	c.mutex.Lock()
	defer c.mutex.Unlock()
	if topo != nil {
		c.topology = topo
	}
	return c
}

// Init initialize the network topology of rabbitmq.
// The initialization required the cosntructed topology that has been set.
func (c *Container) Init() (err error) {
	return
}
