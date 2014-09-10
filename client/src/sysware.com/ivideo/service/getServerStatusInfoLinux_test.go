package service

import (
	"testing"
)

func Test_GetAPSServerStatusInfo(t *testing.T) {
	getInfo := &getServerStatusInfoLinux{}
	if _, err := getInfo.GetInfo("报警服务器1"); nil != err {
		t.Error(err)
	}
}

func Test_GetSIPServerStatusInfo(t *testing.T) {
	getInfo := &getServerStatusInfoLinux{}
	if _, err := getInfo.GetInfo("SIP服务器1"); nil != err {
		t.Error(err)
	}
}
