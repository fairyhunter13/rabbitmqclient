package rabbitmqclient

import "sync/atomic"

type saver struct {
	savedStatus uint64
}

func newSaver() *saver {
	return &saver{}
}

func (s *saver) save() {
	atomic.StoreUint64(&s.savedStatus, TrueUint)
}

func (s *saver) isSaved() bool {
	return atomic.LoadUint64(&s.savedStatus) == TrueUint
}
