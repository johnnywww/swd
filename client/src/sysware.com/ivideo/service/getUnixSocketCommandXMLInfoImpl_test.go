package service

import (
	"sysware.com/ivideo/common"
	"testing"
)

func Test_GetAPSQuitUnixSocketCommandXMLInfoImpl(t *testing.T) {
	getInfo := &getUnixSocketCommandXMLInfoImpl{}
	err := getInfo.GetInfo(common.SERVER_COMMAND_QUIT, &respQuitXmlInfo{}, "@64020000002050000001")
	if err != nil {
		t.Error(err)
	}
}
