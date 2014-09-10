package service

import (
	"testing"
)

func Test_GetSystemTimeInfo(t *testing.T) {
	getInfo := &GetSystemTimeInfoLinux{}
	_, err := getInfo.GetInfo()
	if nil != err {
		t.Error(err)
	}
}
