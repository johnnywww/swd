package service

import (
	"testing"
)

func Test_GetServerStatusXMLInfoLinuxCMSGetInfo(t *testing.T) {
	getInfo := &getServerStatusXMLInfoLinuxCMS{}
	_, err := getInfo.getInfo("")
	if nil != err {
		t.Error(err)
	}
}
