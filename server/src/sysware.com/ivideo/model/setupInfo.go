package model

import (
	"encoding/xml"
	"sysware.com/ivideo/common"
	"sysware.com/ivideo/log"
	"sysware.com/ivideo/service/setup"
)

const (
	setupFileName = common.DEFAULT_CONFIG_FILE_NAME
)

type SetupHandle interface {
	GetSetupInfo(setupInfo *SetupInfo) bool
	SaveSetupInfo(setupInfo *SetupInfo) bool
}

type SetupInfo struct {
	XMLName xml.Name     `xml:"swd"`
	Users   []UserInfo   `xml:"user"`
	Servers []ServerInfo `xml:"server"`
	Port    int          `xml:"port"`
}

func NewSetupHandle() SetupHandle {
	return &SetupInfo{}
}

func (setupInfo1 *SetupInfo) GetSetupInfo(setupInfo *SetupInfo) bool {
	err := setup.NewSetupInfoService().GetSetupInfo(setupInfo, setupFileName)
	if err != nil {
		return false
	}
	return true
}

func (setupInfo1 *SetupInfo) SaveSetupInfo(setupInfo *SetupInfo) bool {
	err := setup.NewSetupInfoService().SaveSetupInfo(setupInfo, setupFileName, common.XML_ENCODE_GB2312)
	if err != nil {
		return false
	}
	return true
}

func (setupInfo *SetupInfo) GetServerIndexByServerName(serverName string) int {
	result := -1
	if nil != setupInfo {
		for index, serverInfo := range setupInfo.Servers {
			if serverInfo.ServerName == serverName {
				result = index
				break
			}
		}
	}
	return result
}

func (setupInfo *SetupInfo) AddServerInfo(serverInfo ServerInfo) {
	setupInfo.Servers = append(setupInfo.Servers, serverInfo)
}

func (setupInfo *SetupInfo) DeleteServerByServerName(serverName string) bool {
	if nil == setupInfo {
		log.WriteLog("没有配置对象")
		return false
	}
	index := setupInfo.GetServerIndexByServerName(serverName)
	if index < 0 {
		log.WriteLog("通过名称%s找不到服务器对象", serverName)
		return false
	}
	servers := setupInfo.Servers[0:index]
	setupInfo.Servers = append(servers, setupInfo.Servers[index+1:len(setupInfo.Servers)]...)
	return true
}

func (setupInfo *SetupInfo) GetUserIndexByUserName(userName string) int {
	result := -1
	if nil != setupInfo {
		for index, userInfo := range setupInfo.Users {
			if userInfo.UserName == userName {
				result = index
				break
			}
		}
	}
	return result
}

func (setupInfo *SetupInfo) AddUserInfo(userInfo UserInfo) {
	setupInfo.Users = append(setupInfo.Users, userInfo)
}

func (setupInfo *SetupInfo) DeleteUserByUserName(userName string) bool {
	if nil == setupInfo {
		log.WriteLog("没有配置对象")
		return false
	}
	index := setupInfo.GetUserIndexByUserName(userName)
	if index < 0 {
		log.WriteLog("通过名称%s找不到用户对象", userName)
		return false
	}
	users := setupInfo.Users[0:index]
	setupInfo.Users = append(users, setupInfo.Users[index+1:len(setupInfo.Users)]...)
	return true
}
