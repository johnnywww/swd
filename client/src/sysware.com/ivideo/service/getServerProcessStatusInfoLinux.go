package service

import (
	"errors"
	"sysware.com/ivideo/common"
	"sysware.com/ivideo/model"
)

type getServerProcessStatusInfoLinux struct {
}

func (getServerProcessStatusInfoLinux *getServerProcessStatusInfoLinux) GetInfo(serverId string) (map[string][]*model.ServerStatusInfo, error) {
	result := map[string][]*model.ServerStatusInfo{}
	var err error = nil
	respQueryStatusXmlInfo, err := NewGetServerStatusXMLInfo().getInfo(serverId)
	if nil != err {
		return result, err
	}
	if common.SERVER_COMMAND_QUERYSTATUS != respQueryStatusXmlInfo.CmdType {
		return result, errors.New("返回的命令不为状态查询命令")
	}
	for _, item := range respQueryStatusXmlInfo.Items {
		for _, cItem := range item.Contents {
			result[item.Type] = append(result[item.Type], &model.ServerStatusInfo{item.Type, cItem})
		}
	}
	return result, nil
}
