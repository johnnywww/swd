package service

import (
	"testing"
)

func Test_GetServerStatusXMLInfoLinuxCommonGetInfo(t *testing.T) {
	getInfo := &getServerStatusXMLInfoLinuxCommon{}
	_, err := getInfo.getInfo("64020000002050000001")
	if nil != err {
		t.Error(err)
	}
}
