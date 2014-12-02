package service

import (
	"testing"
)

func Test_GetMTSServerSetupInfoByConfigFile(t *testing.T) {
	getInfo := &getMTSServerSetupInfoByConfigFile{}
	if _, err := getInfo.GetInfo("/mnt/hgfs/Source/c++/iVideo/Source/mts/build/config.xml"); nil != err {
		t.Error(err)
	}
}
