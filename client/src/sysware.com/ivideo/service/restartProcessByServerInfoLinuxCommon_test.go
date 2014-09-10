package service

import (
	"sysware.com/ivideo/common"
	"sysware.com/ivideo/model"
	"testing"
)

func Test_RestartAPSServerInfo(t *testing.T) {
	getInfo := &restartProcessByServerInfoLinuxCommon{}
	serverInfo := &model.ServerInfo{ServerName: "报警服务器1", Address: "/mnt/hgfs/Source/c++/iVideo/Source/APS/Bin/aps", Type: common.SERVER_TYPE_APS}
	err := getInfo.Restart(serverInfo)
	if nil != err {
		t.Error(err)
	}
}
