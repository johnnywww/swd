package service

import (
	"testing"
)

func Test_GetMemInfo(t *testing.T) {
	getInfo := &GetMemInfoLinux{}
	_, err := getInfo.GetInfo()
	if nil != err {
		t.Error(err)
	}
}
