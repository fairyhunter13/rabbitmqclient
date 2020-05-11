package rabbitmqclient

import "sync/atomic"

type saver struct {
	savedStatus    uint64
	topo           *Topology
	globalExchange *globalExchange
}

func newSaver(topo *Topology, globalExchange *globalExchange) *saver {
	return &saver{
		topo:           topo,
		globalExchange: globalExchange,
	}
}

// Save saves the current global exchange of the saver implementator.
// Save must be called before calling Publish or Consume function.
func (s *saver) Save() {
	if !s.isSaved() {
		s.topo.AddExchangeDeclare(*s.globalExchange.GetGlobalExchange())
		s.save()
	}
	return
}

func (s *saver) save() {
	atomic.StoreUint64(&s.savedStatus, TrueUint)
}

func (s *saver) isSaved() bool {
	return atomic.LoadUint64(&s.savedStatus) == TrueUint
}
