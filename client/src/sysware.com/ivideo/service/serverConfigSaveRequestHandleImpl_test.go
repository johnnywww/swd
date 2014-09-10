package service

import (
	"net/http"
	"testing"
)

func Test_ServerConfigSaveRequestHandleImplSave(t *testing.T) {
	saveInfo := &serverConfigSaveRequestHandleImpl{}
	req, err := http.NewRequest("GET", "http://192.168.128.154/server/saveconfig?servername=报警服务器1&servertype=3&serverId=64020000002050000001&password=123456&registerInterval=4500&defaultHeartInterval=220&address=192.168.128.5&port=2700&logcfgfile=log.cfg&saveCatalog=aps-log&sipServerId=64020000002000000002&sipServerDomain=64020000&sipServerAddress=192.168.128.72&sipServerPort=2722&cmsServerAddress=http://192.168.5.138:8080/ws", nil)
	if nil != err {
		t.Error(err)
	}
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	err = req.ParseForm()
	if nil != err {
		t.Error(err)
	}
	err = saveInfo.Save(req)
	if nil != err {
		t.Error(err)
	}
}
