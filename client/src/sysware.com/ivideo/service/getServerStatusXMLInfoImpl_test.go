package service

import (
	"testing"
)

func Test_GetServerStatusXMLInfoImplGetCommonInfo(t *testing.T) {
	getInfo := &getServerStatusXMLInfoImpl{}
	_, err := getInfo.getInfo("64020000002050000001")
	if nil != err {
		t.Error(err)
	}
}

func Test_GetServerStatusXMLInfoImplGetCMSInfo(t *testing.T) {
	getInfo := &getServerStatusXMLInfoImpl{}
	_, err := getInfo.getInfo("")
	if nil != err {
		t.Error(err)
	}
}
