package service

import (
	"testing"
)

func Test_GetLoaderAvgInfo(t *testing.T) {
	getInfo := &GetLoaderAvgInfoLinux{}
	_, err := getInfo.GetInfo()
	if nil != err {
		t.Error(err)
	}
}
