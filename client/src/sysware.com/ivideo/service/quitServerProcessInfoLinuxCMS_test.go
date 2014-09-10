package service

import (
	"sysware.com/ivideo/common"
	"sysware.com/ivideo/model"
	"testing"
)

func Test_QuitCMSServerInfo(t *testing.T) {
	getInfo := &quitServerProcessInfoLinuxCMS{}
	serverInfo := &model.ServerInfo{ServerName: "中心管理服务器1", Address: "/home/pc01/source/java/apache-tomcat-7.0.53/webapps/cms", Type: common.SERVER_TYPE_CMS}
	err := getInfo.Quit(serverInfo)
	if nil != err {
		t.Error(err)
	}
}
