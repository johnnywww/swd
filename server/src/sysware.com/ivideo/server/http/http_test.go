package http

import (
	"net/url"
	"testing"
)

func Test_GetValidSessionUserInfoByForm(t *testing.T) {
	form := url.Values{}
	form.Set("username", "admin")
	form.Set("password", "11")
	userInfo, errMsg := GetSessionUserInfoByForm(form)
	if nil == userInfo {
		t.Error(errMsg)
	}
}

func Test_GetNoValidSessionUserInfoByForm(t *testing.T) {
	form := url.Values{}
	form.Set("username", "asd")
	form.Set("password", "11")
	userInfo, errMsg := GetSessionUserInfoByForm(form)
	if nil != userInfo {
		t.Error("用户存在")
	}
}
