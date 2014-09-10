package service

import (
	"testing"
)

func Test_GetListenPortInfo(t *testing.T) {
	getInfo := &GetListenPortInfoLinux{}
	_, err := getInfo.GetInfo()
	if nil != err {
		t.Error(err)
	}
}
