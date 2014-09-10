package service

import (
	"testing"
)

func Test_GetProcessInfo(t *testing.T) {
	getInfo := &GetProcessInfoLinux{}
	_, err := getInfo.GetInfo("/mnt/hgfs/Source/c++/iVideo/Source/APS/Bin/aps")
	if nil != err {
		t.Error(err)
	}
}

func Test_GetNotExistProcessInfo(t *testing.T) {
	getInfo := &GetProcessInfoLinux{}
	_, err := getInfo.GetInfo("/mnt/hgfs/Source/c++/iVideo/Source/APS/Bin/aps1")
	if nil == err {
		t.Error("程序居然运行了")
	}
}
