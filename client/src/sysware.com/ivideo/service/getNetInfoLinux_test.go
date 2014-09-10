package service

import (
	"testing"
)

func Test_GetNetInfo(t *testing.T) {
	getInfo := &GetNetInfoLinux{}
	_, err := getInfo.GetInfo()
	if nil != err {
		t.Error(err)
	}
}
