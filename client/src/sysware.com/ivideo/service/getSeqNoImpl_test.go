package service

import (
	"testing"
)

func Test_GetSeqNo(t *testing.T) {
	getInfo := &getSeqNoImpl{currValue: 0, step: 1}
	if v := getInfo.GetNo(); v < 1 {
		t.Error("错误值")
	}
}
