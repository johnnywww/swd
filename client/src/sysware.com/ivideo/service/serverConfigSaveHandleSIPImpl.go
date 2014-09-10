package service

import (
	"errors"
	"net/http"
	"strconv"
	"sysware.com/ivideo/common"
	"sysware.com/ivideo/model"
	"sysware.com/ivideo/service/setup"
)

type serverConfigSaveHandleSIPImpl struct {
}

func (serverConfigSaveHandleSIPImpl *serverConfigSaveHandleSIPImpl) Save(r *http.Request, serverInfo *model.ServerInfo) error {
	if nil == r {
		return errors.New("没有请求")
	}
	setupInfo, err := NewGetServerDetailSetupInfo().GetInfo(serverInfo)
	if nil != err {
		return err
	}
	sipSetupInfo := setupInfo.(*model.SIPServerSetupInfo)
	sipSetupInfo.SIPInfo.Id = r.FormValue("serverId")
	sipSetupInfo.SIPInfo.Password = r.FormValue("password")
	sipSetupInfo.SIPInfo.Domain = r.FormValue("domain")
	sipSetupInfo.SIPInfo.DefaultHeartInterval, err = strconv.Atoi(r.FormValue("defaultHeartInterval"))
	if nil != err {
		return err
	}
	sipSetupInfo.SIPInfo.AddressInfo.IP = r.FormValue("address")
	sipSetupInfo.SIPInfo.AddressInfo.Port, err = strconv.Atoi(r.FormValue("port"))
	if nil != err {
		return err
	}
	sipSetupInfo.CMSServerInfo.Address = r.FormValue("cmsServerAddress")

	sipSetupInfo.DbInfo.AddressInfo.IP = r.FormValue("dbServerAddress")
	sipSetupInfo.DbInfo.AddressInfo.Port, err = strconv.Atoi(r.FormValue("dbServerPort"))
	if nil != err {
		return err
	}
	sipSetupInfo.DbInfo.Password = r.FormValue("dbServerPassword")
	cfgFileName, err := NewGetServerInfoConfigFileName().GetInfo(serverInfo)
	if nil != err {
		return err
	}
	err = setup.NewSetupInfoService().SaveSetupInfo(sipSetupInfo, cfgFileName, common.XML_ENCODE_GB2312)
	return err

}
