package service

import (
	"testing"
)

func Test_ExecuteAPSLinuxProc(t *testing.T) {
	executeLinuxProc := &executeLinuxProcImpl{}
	err := executeLinuxProc.Exec("/mnt/hgfs/Source/c++/iVideo/Source/APS/Bin/aps", "/mnt/hgfs/Source/c++/iVideo/Source/APS/Bin")
	if nil != err {
		t.Error(err)
	}
}
