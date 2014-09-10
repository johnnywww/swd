package service

import (
	"testing"
)

func Test_GetAPSServerSetupInfo(t *testing.T) {
	getInfo := &getServerSetupInfoImpl{}
	if _, _, err := getInfo.GetInfo("报警服务器1"); nil != err {
		t.Error(err)
	}
}

func Test_GetSIPServerSetupInfo(t *testing.T) {
	getInfo := &getServerSetupInfoImpl{}
	if _, _, err := getInfo.GetInfo("SIP服务器1"); nil != err {
		t.Error(err)
	}
}
