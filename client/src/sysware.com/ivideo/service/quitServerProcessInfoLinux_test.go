package service

import (
	"sysware.com/ivideo/common"
	"sysware.com/ivideo/model"
	"testing"
)

func Test_QuitAPSServerInfo1(t *testing.T) {
	getInfo := &quitServerProcessInfoLinux{}
	serverInfo := &model.ServerInfo{ServerName: "报警服务器1", Address: "/mnt/hgfs/Source/c++/iVideo/Source/APS/Bin/aps", Type: common.SERVER_TYPE_APS}
	err := getInfo.Quit(serverInfo)
	if nil != err {
		t.Error(err)
	}
}
