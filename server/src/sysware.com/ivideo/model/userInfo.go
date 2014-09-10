package model

import (
	"sysware.com/ivideo/common"
	"time"
)

type LoginHandle interface {
	Login(userName string, password string)
}

type UserHandle interface {
	LoginHandle
	Logout()
	IsLogin() bool
}

type UserInfo struct {
	UserName    string    `xml:"userName"`
	Password    string    `xml:"password"`
	StatusDate  time.Time `xml:"-"`
	LoginStatus int       `xml:"-"`
}

type UserSaveInfo struct {
	UserInfo
	OUserName string
	AddFlag   bool
}

func (userInfo *UserInfo) Equal(userName string, password string) bool {
	return (userName == userInfo.UserName) && (password == userInfo.Password)
}

func (userInfo *UserInfo) SetLoginState() {
	userInfo.StatusDate = time.Now()
	userInfo.LoginStatus = common.USER_STATE_LOGIN
}

func (userInfo *UserInfo) Login(userName string, password string) bool {
	if !userInfo.Equal(userName, password) {
		return false
	}
	userInfo.SetLoginState()
	return true
}

func (userInfo *UserInfo) Logout() {
	userInfo.StatusDate = time.Now()
	userInfo.LoginStatus = common.USER_STATE_LOGOUT
}

func (userInfo *UserInfo) IsLogin() bool {
	return common.USER_STATE_LOGIN == userInfo.LoginStatus
}
