package rabbitmqclient

func (c *Container) setDefaultTopology() {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	if c.Topology == nil {
		c.Topology = NewTopology()
	}
}

// SetTopology sets the current topology of this container.
func (c *Container) SetTopology(topo *Topology) *Container {
	c.setDefaultTopology()
	if topo != nil {
		c.mutex.Lock()
		c.Topology = topo
		c.mutex.Unlock()
	}
	return c
}

// Init initialize the network topology of rabbitmq.
// The initialization required the cosntructed topology that has been set.
func (c *Container) Init() (err error) {
	if !c.Topology.IsUpdated() {
		return
	}
	err = c.initiator.init(c.Topology)
	return
}
