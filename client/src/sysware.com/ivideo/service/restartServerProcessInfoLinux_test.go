package service

import (
	"testing"
)

func Test_RestartAPSByServerName(t *testing.T) {
	getInfo := &restartServerProcessInfoLinux{}
	err := getInfo.Restart("报警服务器1", "64020000002050000002")
	if nil != err {
		t.Error(err)
	}
}

func Test_RestartSIPByServerName(t *testing.T) {
	getInfo := &restartServerProcessInfoLinux{}
	err := getInfo.Restart("SIP服务器1", "64020000002000000001")
	if nil != err {
		t.Error(err)
	}
}

func Test_RestartCMSByServerName(t *testing.T) {
	getInfo := &restartServerProcessInfoLinux{}
	err := getInfo.Restart("中心管理服务器1", "")
	if nil != err {
		t.Error(err)
	}
}
