package service

import ()

type GetSeqNo interface {
	GetNo() int
}

func NewGetSeqNo() GetSeqNo {
	return &getSeqNoImpl{currValue: 0, step: 1}
}
