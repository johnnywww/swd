package service

import (
	"encoding/xml"
	"errors"
	"sysware.com/ivideo/common"
	"sysware.com/ivideo/model"
)

type quitServerProcessInfoLinux struct {
}

type respQuitXmlInfo struct {
	XMLName xml.Name `xml:"Response"`
	CmdType string   `xml:"CmdType"`
	Sn      int      `xml:"SN"`
	Result  string   `xml:"Result"`
}

func (quitServerProcessInfoLinux *quitServerProcessInfoLinux) Quit(serverInfo *model.ServerInfo, oServerId string) error {
	if nil == serverInfo {
		return errors.New("没有服务器对象")
	}
	var quitServerProcessInfo QuitServerProcessInfo = nil
	if common.SERVER_TYPE_CMS == serverInfo.Type {
		quitServerProcessInfo = &quitServerProcessInfoLinuxCMS{}
	} else {
		quitServerProcessInfo = &quitServerProcessInfoLinuxCommon{}
	}
	return quitServerProcessInfo.Quit(serverInfo, oServerId)
}
