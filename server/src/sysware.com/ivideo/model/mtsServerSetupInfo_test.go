package model

import (
	"sysware.com/ivideo/service/setup"
	"sysware.com/ivideo/utils"
	"testing"
)

func Test_GetMTSServerSetupInfo(t *testing.T) {
	setupInfo := &MTSServerSetupInfo{}
	err := setup.NewSetupInfoService().GetSetupInfo(setupInfo, "/mnt/hgfs/Source/c++/iVideo/Source/mts/config.xml")
	if nil != err { //try a unit test on function
		t.Error(err) // 如果不是如预期的那么就报错
	}
	if utils.IsEmptyStr(setupInfo.MTSInfo.AddressInfo.IP) {
		t.Error("没有服务器地址")
	}
	if utils.IsEmptyStr(setupInfo.SipServerInfo.AddressInfo.IP) {
		t.Error("没有SIP服务器地址")
	}
	if utils.IsEmptyStr(setupInfo.CMSServerInfo.Address) {
		t.Error("没有CMS地址")
	}
}
