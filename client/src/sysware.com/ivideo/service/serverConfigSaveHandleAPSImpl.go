package service

import (
	"errors"
	"net/http"
	"strconv"
	"sysware.com/ivideo/common"
	"sysware.com/ivideo/model"
	"sysware.com/ivideo/service/setup"
	"sysware.com/ivideo/utils"
)

type serverConfigSaveHandleAPSImpl struct {
}

func (serverConfigSaveHandleAPSImpl *serverConfigSaveHandleAPSImpl) Save(r *http.Request, serverInfo *model.ServerInfo) error {
	if nil == r {
		return errors.New("没有请求")
	}
	setupInfo, err := NewGetServerDetailSetupInfo().GetInfo(serverInfo)
	if nil != err {
		return err
	}
	apsSetupInfo := setupInfo.(*model.APSServerSetupInfo)
	apsSetupInfo.APSInfo.Id = r.FormValue("serverId")
	apsSetupInfo.APSInfo.Password = r.FormValue("password")
	apsSetupInfo.APSInfo.RegisterInterval, err = strconv.Atoi(r.FormValue("registerInterval"))
	if nil != err {
		return err
	}
	apsSetupInfo.APSInfo.DefaultHeartInterval, err = strconv.Atoi(r.FormValue("defaultHeartInterval"))
	if nil != err {
		return err
	}
	apsSetupInfo.APSInfo.AddressInfo.IP = r.FormValue("address")
	apsSetupInfo.APSInfo.AddressInfo.Port, err = strconv.Atoi(r.FormValue("port"))
	if nil != err {
		return err
	}
	apsSetupInfo.APSInfo.LogInfo.Logcfgfile = r.FormValue("logcfgfile")
	if utils.IsEmptyStr(apsSetupInfo.APSInfo.LogInfo.Logcfgfile) {
		return errors.New("没有日志配置文件")
	}
	if !utils.IsFileExist(utils.GetFileDir(serverInfo.Address) + "/" + apsSetupInfo.APSInfo.LogInfo.Logcfgfile) {
		return errors.New("日志配置文件不存在")
	}
	apsSetupInfo.APSInfo.LogInfo.SaveCatalog = r.FormValue("saveCatalog")
	apsSetupInfo.SipServerInfo.Id = r.FormValue("sipServerId")
	apsSetupInfo.SipServerInfo.Domain = r.FormValue("sipServerDomain")
	apsSetupInfo.SipServerInfo.AddressInfo.IP = r.FormValue("sipServerAddress")
	apsSetupInfo.SipServerInfo.AddressInfo.Port, err = strconv.Atoi(r.FormValue("sipServerPort"))
	if nil != err {
		return err
	}
	apsSetupInfo.CMSServerInfo.Address = r.FormValue("cmsServerAddress")
	cfgFileName, err := NewGetServerInfoConfigFileName().GetInfo(serverInfo)
	if nil != err {
		return err
	}
	err = setup.NewSetupInfoService().SaveSetupInfo(apsSetupInfo, cfgFileName, common.XML_ENCODE_GB2312)
	return err

}
