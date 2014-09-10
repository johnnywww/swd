package service

import (
	"sysware.com/ivideo/common"
	"sysware.com/ivideo/utils"
	"testing"
)

func Test_GetXMLInfo(t *testing.T) {
	getInfo := &getServerRequestXMLInfoImpl{}
	if utils.IsEmptyStr(getInfo.GetXML(common.SERVER_COMMAND_QUERYSTATUS)) {
		t.Error("错误的xml请求信息")
	}
}
