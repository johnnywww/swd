package service

import (
	"sysware.com/ivideo/common"
	"sysware.com/ivideo/model"
	"testing"
)

func Test_GetAPSProcessInfoByServerInfo(t *testing.T) {
	getInfo := &getProcessInfoByServerInfoLinux{}
	serverInfo := &model.ServerInfo{ServerName: "报警服务器1", Address: "/mnt/hgfs/Source/c++/iVideo/Source/APS/Bin/aps", Type: common.SERVER_TYPE_APS}
	_, err := getInfo.GetInfo(serverInfo)
	if nil != err {
		t.Error(err)
	}
}

func Test_GetCMSProcessInfoByServerInfo(t *testing.T) {
	getInfo := &getProcessInfoByServerInfoLinux{}
	serverInfo := &model.ServerInfo{ServerName: "中心管理服务器1", Address: "/home/pc01/source/java/apache-tomcat-7.0.53/webapps/cms", Type: common.SERVER_TYPE_CMS}
	_, err := getInfo.GetInfo(serverInfo)
	if nil == err {
		t.Error("程序居然运行了")
	}
}
