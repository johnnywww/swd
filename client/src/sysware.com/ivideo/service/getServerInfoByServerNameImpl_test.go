package service

import (
	"testing"
)

func Test_GetServerInfoByServerNameImplGetAPSInfo(t *testing.T) {
	getInfo := &getServerInfoByServerNameImpl{}
	if _, err := getInfo.getInfo("SIP服务器1"); nil != err {
		t.Error(err)
	}
}
