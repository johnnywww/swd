package model

import (
	"testing"
)

func Test_GetSetupInfo(t *testing.T) {
	setupInfo := &SetupInfo{}

	if !setupInfo.GetSetupInfo(setupInfo) { //try a unit test on function
		t.Error("没通过") // 如果不是如预期的那么就报错
	}
}

func Test_SaveSetupInfo(t *testing.T) {
	setupInfo := &SetupInfo{}
	setupInfo.Users = append(setupInfo.Users, UserInfo{UserName: "admin", Password: "111111"})
	setupInfo.Servers = append(setupInfo.Servers, ServerInfo{ServerName: "SIP服务器", Address: "http://192.168.128.24:2566", Type: 1})
	if !setupInfo.SaveSetupInfo(setupInfo) { //try a unit test on function
		t.Error("没通过") // 如果不是如预期的那么就报错
	}
}

func Test_GetServerIndexByServerName(t *testing.T) {
	setupInfo := &SetupInfo{}
	setupInfo.Users = append(setupInfo.Users, UserInfo{UserName: "admin", Password: "111111"})
	setupInfo.Servers = append(setupInfo.Servers, ServerInfo{ServerName: "SIP服务器", Address: "http://192.168.128.24:2566", Type: 1})
	setupInfo.Servers = append(setupInfo.Servers, ServerInfo{ServerName: "SIP服务器2", Address: "http://192.168.128.24:2566", Type: 1})
	if 1 != setupInfo.GetServerIndexByServerName("SIP服务器2") { //try a unit test on function
		t.Error("没找到对应服务器") // 如果不是如预期的那么就报错
	}
}

func Test_DeleteServerByServerName(t *testing.T) {
	setupInfo := &SetupInfo{}
	setupInfo.Users = append(setupInfo.Users, UserInfo{UserName: "admin", Password: "111111"})
	setupInfo.Servers = append(setupInfo.Servers, ServerInfo{ServerName: "SIP服务器", Address: "http://192.168.128.24:2566", Type: 1})
	setupInfo.Servers = append(setupInfo.Servers, ServerInfo{ServerName: "SIP服务器2", Address: "http://192.168.128.24:2566", Type: 1})
	if !setupInfo.DeleteServerByServerName("SIP服务器") { //try a unit test on function
		t.Error("无法删除对应服务器") // 如果不是如预期的那么就报错
	}
	if "SIP服务器2" != setupInfo.Servers[0].ServerName {
		t.Error("数据不正确")
	}
}
