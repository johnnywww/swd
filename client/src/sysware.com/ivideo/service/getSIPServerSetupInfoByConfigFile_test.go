package service

import (
	"testing"
)

func Test_GetSIPServerSetupInfoByConfigFile(t *testing.T) {
	getInfo := &getSIPServerSetupInfoByConfigFile{}
	_, err := getInfo.GetInfo("/mnt/hgfs/Source/c++/iVideo/Source/SIPServer/build/config.xml")
	if nil != err {
		t.Error(err)
	}
}
