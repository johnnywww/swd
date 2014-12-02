package service

import (
	"net/http"
	"sysware.com/ivideo/common"
	"sysware.com/ivideo/model"
	"testing"
)

func Test_ServerConfigSaveHandleMTSImplSave(t *testing.T) {
	saveInfo := &serverConfigSaveHandleMTSImpl{}
	req, err := http.NewRequest("GET", "http://192.168.128.154/server/saveconfig?servername=转发服务器1&servertype=2&serverId=65020000002020000001&password=1234567&registerInterval=4500&defaultHeartInterval=220&address=192.168.5.154&port=5063&sipServerId=65020000002000000002&sipServerDomain=65020000&sipServerAddress=192.168.5.138&sipServerPort=5060&cmsServerAddress=http://192.168.5.138:8080/ws", nil)
	if nil != err {
		t.Error(err)
	}
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	err = req.ParseForm()
	if nil != err {
		t.Error(err)
	}
	serverInfo := &model.ServerInfo{ServerName: "转发服务器1", Address: "/mnt/hgfs/Source/c++/iVideo/Source/mts/build/mts2", Type: common.SERVER_TYPE_MTS}
	err = saveInfo.Save(req, serverInfo)
	if nil != err {
		t.Error(err)
	}
}
