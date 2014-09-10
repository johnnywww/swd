package service

import (
	"sysware.com/ivideo/common"
	"sysware.com/ivideo/model"
	"testing"
)

func Test_GetServerDetailSetupInfo(t *testing.T) {
	getInfo := &getServerDetailSetupInfoImpl{}
	serverInfo := &model.ServerInfo{ServerName: "报警服务器1", Address: "/mnt/hgfs/Source/c++/iVideo/Source/APS/Bin/aps", Type: common.SERVER_TYPE_APS}
	if _, err := getInfo.GetInfo(serverInfo); nil != err {
		t.Error(err)
	}
}
