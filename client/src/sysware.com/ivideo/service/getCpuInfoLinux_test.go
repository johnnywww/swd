package service

import (
	"testing"
)

func Test_GetCpuInfo(t *testing.T) {
	getInfo := &GetCpuInfoLinux{}
	_, err := getInfo.GetInfo()
	if nil != err {
		t.Error(err)
	}
}
