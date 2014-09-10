package service

import (
	"errors"
	"fmt"
	"sysware.com/ivideo/model"
)

type getServerInfoByServerNameImpl struct {
}

func (getServerInfoByServerNameImpl *getServerInfoByServerNameImpl) getInfo(serverName string) (*model.ServerInfo, error) {
	setupInfo := model.SetupInfo{}
	if !model.NewSetupHandle().GetSetupInfo(&setupInfo) {
		return nil, errors.New("读取配置失败")
	}
	index := setupInfo.GetServerIndexByServerName(serverName)
	if index < 0 {
		return nil, errors.New(fmt.Sprintf("没有找到对应的服务器%s", serverName))
	}
	return &setupInfo.Servers[index], nil
}
