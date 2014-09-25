package service

import (
	"testing"
)

func Test_ExecuteAPSLinuxProcForkChild(t *testing.T) {
	executeLinuxProcForkChildImpl := &executeLinuxProcForkChildImpl{}
	err := executeLinuxProcForkChildImpl.Exec("/mnt/hgfs/Source/c++/iVideo/Source/APS/Bin/aps", "/mnt/hgfs/Source/c++/iVideo/Source/APS/Bin")
	if nil != err {
		t.Error(err)
	}
}
