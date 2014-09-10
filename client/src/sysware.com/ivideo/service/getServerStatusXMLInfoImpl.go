package service

import (
	"sysware.com/ivideo/utils"
)

type getServerStatusXMLInfoImpl struct {
}

func (getServerStatusXMLInfoImpl *getServerStatusXMLInfoImpl) getInfo(serverId string) (*respQueryStatusXmlInfo, error) {
	var getServerStatusXMLInfo GetServerStatusXMLInfo = nil
	if !utils.IsEmptyStr(serverId) {
		getServerStatusXMLInfo = &getServerStatusXMLInfoLinuxCommon{}
	} else {
		getServerStatusXMLInfo = &getServerStatusXMLInfoLinuxCMS{}
	}
	return getServerStatusXMLInfo.getInfo(serverId)
}
