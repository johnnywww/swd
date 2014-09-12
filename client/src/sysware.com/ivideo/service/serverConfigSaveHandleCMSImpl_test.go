package service

import (
	"net/http"
	"sysware.com/ivideo/common"
	"sysware.com/ivideo/model"
	"testing"
)

func Test_ServerConfigSaveHandleCMSImplSave(t *testing.T) {
	saveInfo := &serverConfigSaveHandleCMSImpl{}
	req, err := http.NewRequest("GET", "http://192.168.128.154/server/saveconfig?servername=中心管理服务器1&servertype=0&dbAddress=jdbc:oracle:thin:@192.168.5.99:1521:orcl&dbUserName=ivideo&dbPassword=sysware&dbDefaultSchema=ivideo&activemqAddress=tcp://192.168.5.154:61616", nil)
	if nil != err {
		t.Error(err)
	}
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	err = req.ParseForm()
	if nil != err {
		t.Error(err)
	}
	serverInfo := &model.ServerInfo{ServerName: "中心管理服务器1", Address: "/home/pc01/source/java/apache-tomcat-7.0.53/webapps/cms", Type: common.SERVER_TYPE_CMS}
	err = saveInfo.Save(req, serverInfo)
	if nil != err {
		t.Error(err)
	}
}
