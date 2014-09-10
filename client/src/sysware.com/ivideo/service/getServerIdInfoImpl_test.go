package service

import (
	"testing"
)

func Test_GetServerIdInfo(t *testing.T) {
	getInfo := &getServerIdInfoImpl{}
	if _, err := getInfo.getInfo("报警服务器1"); nil != err {
		t.Error(err)
	}
}
