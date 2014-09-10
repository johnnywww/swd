package model

import (
	"sysware.com/ivideo/service/setup"
	"testing"
)

func Test_GetAPSServerSetupInfo(t *testing.T) {
	setupInfo := &APSServerSetupInfo{}
	err := setup.NewSetupInfoService().GetSetupInfo(setupInfo, "/mnt/hgfs/Source/c++/iVideo/Source/APS/Bin/config.xml")
	if nil != err { //try a unit test on function
		t.Error(err) // 如果不是如预期的那么就报错
	}
}
