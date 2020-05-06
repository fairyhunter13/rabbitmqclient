package rabbitmqclient

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
	err = c.Initiator.init(c.Topology)
	return
}
