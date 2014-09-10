package service

import (
	"testing"
)

func Test_RestartAPS(t *testing.T) {
	getInfo := &restartServerProcessInfoLinux{}
	err := getInfo.Restart("报警服务器1")
	if nil != err {
		t.Error(err)
	}
}

func Test_RestartSIP(t *testing.T) {
	getInfo := &restartServerProcessInfoLinux{}
	err := getInfo.Restart("SIP服务器1")
	if nil != err {
		t.Error(err)
	}
}

func Test_RestartCMS(t *testing.T) {
	getInfo := &restartServerProcessInfoLinux{}
	err := getInfo.Restart("中心管理服务器1")
	if nil != err {
		t.Error(err)
	}
}
