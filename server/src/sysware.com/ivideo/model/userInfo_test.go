package model

import (
	"sysware.com/ivideo/common"
	"testing"
)

func Test_Login(t *testing.T) {
	userInfo := &UserInfo{UserName: "admin", Password: "111111"}
	if !userInfo.Login("admin", "111111") {
		t.Error("登录用户失败")
	}

	if !userInfo.IsLogin() {
		t.Error("登录状态不为已登录状态")
	}
}

func Test_LogOut(t *testing.T) {
	userInfo := &UserInfo{UserName: "admin", Password: "111111", LoginStatus: common.USER_STATE_LOGIN}
	userInfo.Logout()

	if userInfo.IsLogin() {
		t.Error("登录状态为已登录状态")
	}
}
