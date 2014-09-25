package service

import (
	"testing"
)

func Test_ExecuteAPSLinuxProcStartProcess(t *testing.T) {
	executeLinuxProcStartProcessImpl := &executeLinuxProcStartProcessImpl{}
	err := executeLinuxProcStartProcessImpl.Exec("/mnt/hgfs/Source/c++/iVideo/Source/APS/Bin/aps", "/mnt/hgfs/Source/c++/iVideo/Source/APS/Bin")
	if nil != err {
		t.Error(err)
	}
}
