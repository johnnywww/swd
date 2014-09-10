package service

import (
	"net/http"
	"sysware.com/ivideo/common"
	"sysware.com/ivideo/model"
	"testing"
)

func Test_ServerConfigSaveHandleSIPImplSave(t *testing.T) {
	saveInfo := &serverConfigSaveHandleSIPImpl{}
	req, err := http.NewRequest("GET", "http://192.168.128.154/server/saveconfig?servername=SIP服务器1&servertype=1&serverId=64020000002000000002&password=123456&domain=64020000&defaultHeartInterval=220&address=192.168.128.5&port=2700&cmsServerAddress=http://192.168.5.138:8080/ws&dbServerAddress=192.168.5.139&dbServerPort=2701&dbServerPassword=fsddsfsd", nil)
	if nil != err {
		t.Error(err)
	}
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	err = req.ParseForm()
	if nil != err {
		t.Error(err)
	}
	serverInfo := &model.ServerInfo{ServerName: "SIP服务器1", Address: "/mnt/hgfs/Source/c++/iVideo/Source/SIPServer/build/sipserver2", Type: common.SERVER_TYPE_SIP}
	err = saveInfo.Save(req, serverInfo)
	if nil != err {
		t.Error(err)
	}
}
