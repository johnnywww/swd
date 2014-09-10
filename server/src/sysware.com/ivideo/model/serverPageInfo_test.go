package model

import (
	"testing"
)

func Test_GetServerPageInfo(t *testing.T) {
	serverPageInfo := ServerPageInfo{}
	var servers []ServerInfo
	servers = append(servers, ServerInfo{ServerName: "SIP服务器", Address: "http://192.168.128.24:2566", Type: 1})
	serverPageInfo.GetServerPageInfo(2, 10)
	if 0 != serverPageInfo.Page.PageNo {
		t.Error("页数不对")
	}
}
