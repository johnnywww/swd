package service

import (
	"sysware.com/ivideo/common"
	"sysware.com/ivideo/model"
	"testing"
)

func Test_GetAPSServerInfoConfigFileNameImplGet(t *testing.T) {
	getInfo := &getServerInfoConfigFileNameImpl{}
	serverInfo := &model.ServerInfo{ServerName: "报警服务器1", Address: "/mnt/hgfs/Source/c++/iVideo/Source/APS/Bin/aps", Type: common.SERVER_TYPE_APS}
	cfgFileName, err := getInfo.GetInfo(serverInfo)
	if nil != err {
		t.Error(err)
	}
	if "/mnt/hgfs/Source/c++/iVideo/Source/APS/Bin/config.xml" != cfgFileName {
		t.Error(cfgFileName)
	}
}
