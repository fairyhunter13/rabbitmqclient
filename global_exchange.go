package rabbitmqclient

import "sync"

type globalExchange struct {
	mutex    *sync.RWMutex
	exchange *ExchangeDeclare
}

func newGlobalExchange() *globalExchange {
	return &globalExchange{
		mutex:    new(sync.RWMutex),
		exchange: new(ExchangeDeclare).Default(),
	}
}

// SetExchange sets the exchange of the global exchange.
func (ge *globalExchange) SetExchange(exc *ExchangeDeclare) *globalExchange {
	if exc != nil {
		ge.mutex.Lock()
		ge.exchange = exc
		ge.mutex.Unlock()
	}
	return ge
}

// SetExchangeName sets the exchange name of the global exchange.
func (ge *globalExchange) SetExchangeName(name string) *globalExchange {
	if name != "" {
		ge.mutex.Lock()
		ge.exchange.Name = name
		ge.mutex.Unlock()
	}
	return ge
}

// GetGlobalExchange get the exchange declare inside the global exchange.
func (ge *globalExchange) GetGlobalExchange() *ExchangeDeclare {
	ge.mutex.RLock()
	defer ge.mutex.RUnlock()
	return ge.exchange
}
