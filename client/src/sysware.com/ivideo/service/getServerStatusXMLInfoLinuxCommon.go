package service

import (
	"errors"
	"sysware.com/ivideo/common"
	"sysware.com/ivideo/utils"
)

type getServerStatusXMLInfoLinuxCommon struct {
}

func (getServerStatusXMLInfoLinuxCommon *getServerStatusXMLInfoLinuxCommon) getInfo(serverId string) (*respQueryStatusXmlInfo, error) {
	if utils.IsEmptyStr(serverId) {
		return nil, errors.New("没有服务器编号")
	}
	respQueryStatusXmlInfo := respQueryStatusXmlInfo{}
	err := getAbstarctUnixSocketXMLInfo(serverId, common.SERVER_COMMAND_QUERYSTATUS, &respQueryStatusXmlInfo)
	return &respQueryStatusXmlInfo, err
}
