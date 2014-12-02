package service

import (
	"errors"
	"net/http"
	"strconv"
	"sysware.com/ivideo/common"
	"sysware.com/ivideo/model"
	"sysware.com/ivideo/service/setup"
)

type serverConfigSaveHandleMTSImpl struct {
}

func (serverConfigSaveHandleMTSImpl *serverConfigSaveHandleMTSImpl) Save(r *http.Request, serverInfo *model.ServerInfo) error {
	if nil == r {
		return errors.New("没有请求")
	}
	setupInfo, err := NewGetServerDetailSetupInfo().GetInfo(serverInfo)
	if nil != err {
		return err
	}
	mtsSetupInfo := setupInfo.(*model.MTSServerSetupInfo)
	mtsSetupInfo.MTSInfo.Id = r.FormValue("serverId")
	mtsSetupInfo.MTSInfo.Password = r.FormValue("password")
	mtsSetupInfo.MTSInfo.RegisterInterval, err = strconv.Atoi(r.FormValue("registerInterval"))
	if nil != err {
		return err
	}
	mtsSetupInfo.MTSInfo.DefaultHeartInterval, err = strconv.Atoi(r.FormValue("defaultHeartInterval"))
	if nil != err {
		return err
	}
	mtsSetupInfo.MTSInfo.AddressInfo.IP = r.FormValue("address")
	mtsSetupInfo.MTSInfo.AddressInfo.Port, err = strconv.Atoi(r.FormValue("port"))
	if nil != err {
		return err
	}

	mtsSetupInfo.SipServerInfo.Id = r.FormValue("sipServerId")
	mtsSetupInfo.SipServerInfo.Domain = r.FormValue("sipServerDomain")
	mtsSetupInfo.SipServerInfo.AddressInfo.IP = r.FormValue("sipServerAddress")
	mtsSetupInfo.SipServerInfo.AddressInfo.Port, err = strconv.Atoi(r.FormValue("sipServerPort"))
	if nil != err {
		return err
	}

	mtsSetupInfo.CMSServerInfo.Address = r.FormValue("cmsServerAddress")
	cfgFileName, err := NewGetServerInfoConfigFileName().GetInfo(serverInfo)
	if nil != err {
		return err
	}
	err = setup.NewSetupInfoService().SaveSetupInfo(mtsSetupInfo, cfgFileName, common.XML_ENCODE_GB2312)
	return err
}
