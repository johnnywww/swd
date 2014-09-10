package service

import (
	"errors"
	"net/http"
	"sysware.com/ivideo/common"
	"sysware.com/ivideo/model"
)

type serverConfigSaveHandleImpl struct {
}

func (serverConfigSaveHandleImpl *serverConfigSaveHandleImpl) Save(r *http.Request, serverInfo *model.ServerInfo) error {
	if nil == r {
		return errors.New("没有请求")
	}
	if nil == serverInfo {
		return errors.New("没有服务器对象信息")
	}
	switch serverInfo.Type {
	case common.SERVER_TYPE_APS:
		return (&serverConfigSaveHandleAPSImpl{}).Save(r, serverInfo)
	case common.SERVER_TYPE_SIP:
		return (&serverConfigSaveHandleSIPImpl{}).Save(r, serverInfo)
	case common.SERVER_TYPE_CMS:
		return (&serverConfigSaveHandleCMSImpl{}).Save(r, serverInfo)

	}
	return errors.New("无法处理的服务器类型")
}
