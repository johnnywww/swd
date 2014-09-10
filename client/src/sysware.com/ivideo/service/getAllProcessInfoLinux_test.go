package service

import (
	"testing"
)

func Test_GetAllProcessInfo(t *testing.T) {
	getInfo := &GetAllProcessInfoLinux{}
	_, err := getInfo.GetInfo()
	if nil != err {
		t.Error(err)
	}
}
