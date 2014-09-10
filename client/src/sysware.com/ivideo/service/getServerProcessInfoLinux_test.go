package service

import (
	"testing"
)

func Test_GetAPSServerProcessInfo(t *testing.T) {
	getInfo := &getServerProcessInfoLinux{}
	_, err := getInfo.GetInfo("报警服务器1")
	if nil != err {
		t.Error(err)
	}
}

func Test_GetCMSServerProcessInfo(t *testing.T) {
	getInfo := &getServerProcessInfoLinux{}
	_, err := getInfo.GetInfo("中心管理服务器1")
	if nil != err {
		t.Error(err)
	}
}
