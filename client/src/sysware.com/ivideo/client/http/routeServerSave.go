package http

import (
	"errors"
	"fmt"
	"sysware.com/ivideo/model"
	shttp "sysware.com/ivideo/server/http"
	"sysware.com/ivideo/service"
	"sysware.com/ivideo/utils"
)

type RouteServerSave struct {
	model.RouteHandle
}

func (routeServerSave *RouteServerSave) CheckServerSave(oServerName string, pServerInfo *model.ServerInfo) error {

	if !utils.IsFileExist(pServerInfo.Address) {
		return errors.New(fmt.Sprintf("没有找到服务器程序%s", pServerInfo.Address))
	}
	_, err := service.NewGetServerInfoConfigFileName().GetInfo(pServerInfo)
	return err
}

func NewRouteServerSave() model.RouteHandle {
	return &RouteServerSave{&shttp.RouteServerSave{new(RouteServerSave)}}
}
