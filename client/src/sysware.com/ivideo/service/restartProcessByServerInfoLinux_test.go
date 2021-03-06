package service

import (
	"sysware.com/ivideo/common"
	"sysware.com/ivideo/model"
	"testing"
)

func Test_RestartAPSServerInfo1(t *testing.T) {
	getInfo := &restartProcessByServerInfoLinux{}
	serverInfo := &model.ServerInfo{ServerName: "报警服务器1", Address: "/mnt/hgfs/Source/c++/iVideo/Source/APS/Bin/aps", Type: common.SERVER_TYPE_APS}
	err := getInfo.Restart(serverInfo)
	if nil != err {
		t.Error(err)
	}
}

func Test_RestartMTSServerInfo1(t *testing.T) {
	getInfo := &restartProcessByServerInfoLinux{}
	serverInfo := &model.ServerInfo{ServerName: "转发服务器1", Address: "/mnt/hgfs/Source/c++/iVideo/Source/mts/build/mts2", Type: common.SERVER_TYPE_MTS}
	err := getInfo.Restart(serverInfo)
	if nil != err {
		t.Error(err)
	}
}
