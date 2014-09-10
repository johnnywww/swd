package service

import (
	"errors"
	"sysware.com/ivideo/common"
	"sysware.com/ivideo/utils"
)

type getServerStatusXMLInfoLinuxCMS struct {
}

func (getServerStatusXMLInfoLinuxCMS *getServerStatusXMLInfoLinuxCMS) getInfo(serverId string) (*respQueryStatusXmlInfo, error) {
	if !utils.IsEmptyStr(serverId) {
		return nil, errors.New("有服务器编号")
	}
	respQueryStatusXmlInfo := respQueryStatusXmlInfo{}
	err := getCmsUnixSocketXMLInfo(common.SERVER_COMMAND_QUERYSTATUS, &respQueryStatusXmlInfo)
	return &respQueryStatusXmlInfo, err
}
