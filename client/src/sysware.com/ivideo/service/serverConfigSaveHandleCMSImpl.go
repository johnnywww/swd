package service

import (
	"errors"
	"net/http"
	"sysware.com/ivideo/common"
	"sysware.com/ivideo/model"
	"sysware.com/ivideo/service/setup"
	"sysware.com/ivideo/utils"
)

type serverConfigSaveHandleCMSImpl struct {
}

func (serverConfigSaveHandleCMSImpl *serverConfigSaveHandleCMSImpl) Save(r *http.Request, serverInfo *model.ServerInfo) error {
	if nil == r {
		return errors.New("没有请求")
	}
	cfgFileName, err := NewGetServerInfoConfigFileName().GetInfo(serverInfo)
	if nil != err {
		return err
	}
	var cMSServerSetupInfo = model.CMSServerSetupInfo{}
	err = setup.NewSetupInfoService().GetSetupInfo(&cMSServerSetupInfo, cfgFileName)
	if nil != err {
		return err
	}
	v := r.FormValue("dbPassword")
	if utils.IsEmptyStr(v) {
		return errors.New("没有数据库密码")
	}
	err = cMSServerSetupInfo.SetProperty("connection.password", v)
	if nil != err {
		return err
	}
	v = r.FormValue("dbAddress")
	if utils.IsEmptyStr(v) {
		return errors.New("没有数据库地址")
	}
	err = cMSServerSetupInfo.SetProperty("connection.url", v)
	if nil != err {
		return err
	}
	v = r.FormValue("dbUserName")
	if utils.IsEmptyStr(v) {
		return errors.New("没有数据库用户名")
	}
	err = cMSServerSetupInfo.SetProperty("connection.username", v)
	if nil != err {
		return err
	}
	v = r.FormValue("dbDefaultSchema")
	if utils.IsEmptyStr(v) {
		return errors.New("没有数据库表空间")
	}
	err = cMSServerSetupInfo.SetProperty("default_schema", v)
	if nil != err {
		return err
	}
	return setup.NewCMSSetupInfoService().SaveSetupInfo(cMSServerSetupInfo, cfgFileName, common.XML_ENCODE_UTF8)
}
