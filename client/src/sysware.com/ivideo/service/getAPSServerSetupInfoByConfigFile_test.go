package service

import (
	"testing"
)

func Test_GetAPSServerSetupInfoByConfigFile(t *testing.T) {
	getInfo := &getAPSServerSetupInfoByConfigFile{}
	if _, err := getInfo.GetInfo("/mnt/hgfs/Source/c++/iVideo/Source/APS/Bin/config.xml"); nil != err {
		t.Error(err)
	}
}
