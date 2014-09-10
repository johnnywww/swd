package service

import (
	"sync"
)

var seqMutex sync.Mutex

func init() {
	seqMutex = sync.Mutex{}
}

type getSeqNoImpl struct {
	step      int
	currValue int
}

func (getSeqNoImpl *getSeqNoImpl) GetNo() int {
	seqMutex.Lock()
	defer seqMutex.Unlock()
	getSeqNoImpl.currValue = getSeqNoImpl.currValue + getSeqNoImpl.step
	return getSeqNoImpl.currValue
}
