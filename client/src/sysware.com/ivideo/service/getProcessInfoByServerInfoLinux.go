package service

import (
	"sysware.com/ivideo/common"
	"sysware.com/ivideo/model"
)

type getProcessInfoByServerInfoLinux struct {
}

func (getProcessInfoByServerInfoLinux *getProcessInfoByServerInfoLinux) GetInfo(serverInfo *model.ServerInfo) (*model.ProcessInfo, error) {
	var getProcessInfo GetProcessInfo = nil
	if common.SERVER_TYPE_CMS == serverInfo.Type {
		getProcessInfo = &getCMSProcessInfoLinux{}
	} else {
		getProcessInfo = NewGetProcessInfo()
	}
	return getProcessInfo.GetInfo(serverInfo.Address)
}
