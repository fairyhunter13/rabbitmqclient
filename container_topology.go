package rabbitmqclient

// SetTopology sets the current topology of this container.
func (c *Container) SetTopology(topo *Topology) *Container {
	if topo != nil {
		c.mutex.Lock()
		c.Topology = topo
		c.mutex.Unlock()
	}
	return c
}

// GetTopology gets the topology of this container.
func (c *Container) GetTopology() *Topology {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	return c.Topology
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
