package service

import (
	"sysware.com/ivideo/common"
	"sysware.com/ivideo/model"
	"testing"
)

func Test_GetServerDetailCMSSetupInfo(t *testing.T) {
	getInfo := &getServerDetailSetupInfoCMSImpl{}
	serverInfo := &model.ServerInfo{ServerName: "中心管理服务器1", Address: "/home/pc01/source/java/apache-tomcat-7.0.53/webapps/cms", Type: common.SERVER_TYPE_CMS}
	if _, err := getInfo.GetInfo(serverInfo); nil != err {
		t.Error(err)
	}
}
