package utils

import (
	"sysware.com/ivideo/common"
	"testing"
)

func Test_GetBaiduServerState(t *testing.T) {
	if common.SERVER_STATE_ONLINE != GetServerState("http://www.baidu.com") {
		t.Error("访问baidu失败")
	}
}

func Test_GetBaidu1ServerState(t *testing.T) {
	if common.SERVER_STATE_OFFLINE != GetServerState("http://192.168.128.167:9500") {
		t.Error("访问192.168.128.167成功")
	}
}

func Test_IsEmptyStr(t *testing.T) {
	if !IsEmptyStr("") {
		t.Error("string is not empty")
	}
}

func Test_IsLinuxOS(t *testing.T) {
	if !IsLinuxOS() {
		t.Error("is not linux os")
	}
}

func Test_IsWindowsOS(t *testing.T) {
	if !IsWindowsOS() {
		t.Error("is not windows os")
	}
}
