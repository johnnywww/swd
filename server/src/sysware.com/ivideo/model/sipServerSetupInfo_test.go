package model

import (
	"sysware.com/ivideo/service/setup"
	"sysware.com/ivideo/utils"
	"testing"
)

func Test_GetSIPServerSetupInfo(t *testing.T) {
	setupInfo := &SIPServerSetupInfo{}
	err := setup.NewSetupInfoService().GetSetupInfo(setupInfo, "/mnt/hgfs/Source/c++/iVideo/Source/SIPServer/config.xml")
	if nil != err { //try a unit test on function
		t.Error(err) // 如果不是如预期的那么就报错
	}
	if utils.IsEmptyStr(setupInfo.ActiveMQInfo.AddressInfo.IP) {
		t.Error("没有ActiveMQ地址")
	}
	if utils.IsEmptyStr(setupInfo.ActiveMQInfo.Topic) {
		t.Error("没有ActiveMQ主题")
	}
}
