package service

import (
	"encoding/xml"
	"errors"
	"sysware.com/ivideo/common"
	"sysware.com/ivideo/model"
	"sysware.com/ivideo/utils"
)

type getServerProcessStatusInfoLinux struct {
}

type respQueryStatusXmlInfo struct {
	XMLName xml.Name                  `xml:"Response"`
	CmdType string                    `xml:"CmdType"`
	Sn      int                       `xml:"SN"`
	Items   []respQueryStatusItemInfo `xml:"Items>Item"`
}

type respQueryStatusItemInfo struct {
	Type     string   `xml:"Type"`
	Contents []string `xml:"Content"`
}

func (getServerProcessStatusInfoLinux *getServerProcessStatusInfoLinux) GetInfo(serverId string) (map[string][]*model.ServerStatusInfo, error) {
	result := map[string][]*model.ServerStatusInfo{}
	var err error = nil
	respQueryStatusXmlInfo := respQueryStatusXmlInfo{}
	if !utils.IsEmptyStr(serverId) {
		err = getAbstarctUnixSocketXMLInfo(serverId, common.SERVER_COMMAND_QUERYSTATUS, &respQueryStatusXmlInfo)
	} else {
		err = getCmsUnixSocketXMLInfo(common.SERVER_COMMAND_QUERYSTATUS, &respQueryStatusXmlInfo)
	}

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
