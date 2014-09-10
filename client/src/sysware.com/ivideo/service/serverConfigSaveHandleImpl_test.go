package service

import (
	"net/http"
	"sysware.com/ivideo/common"
	"sysware.com/ivideo/model"
	"testing"
)

func Test_ServerConfigSaveHandleImplSave(t *testing.T) {
	saveInfo := &serverConfigSaveHandleImpl{}
	req, err := http.NewRequest("GET", "http://192.168.128.154/server/saveconfig?servername=报警服务器1&servertype=3&serverId=64020000002050000001&password=123456&registerInterval=4500&defaultHeartInterval=220&address=192.168.128.5&port=2700&logcfgfile=log.cfg&saveCatalog=aps-log&sipServerId=64020000002000000002&sipServerDomain=64020000&sipServerAddress=192.168.128.72&sipServerPort=2722&cmsServerAddress=http://192.168.5.138:8080/ws", nil)
	if nil != err {
		t.Error(err)
	}
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	err = req.ParseForm()
	if nil != err {
		t.Error(err)
	}
	serverInfo := &model.ServerInfo{ServerName: "报警服务器1", Address: "/mnt/hgfs/Source/c++/iVideo/Source/APS/Bin/aps", Type: common.SERVER_TYPE_APS}
	err = saveInfo.Save(req, serverInfo)
	if nil != err {
		t.Error(err)
	}
}
