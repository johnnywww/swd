package http

import (
	"github.com/martini-contrib/render"
	"net/url"
	"sysware.com/ivideo/model"
	"sysware.com/ivideo/utils"
)

var serverRoutes *model.RouteHandles

func init() {
	serverRoutes = &model.RouteHandles{Handles: []model.RouteHandle{&RouteRoot{}, &RouteLogin{}, &RouteLogout{}, &RouteNotFound{}, &RouteServerSetup{}, &RouteUser{}, &RouteServer{}}}
}

func NewRouteHandle() model.RouteHandle {
	return serverRoutes
}

func checkLoginUser(form url.Values, re render.Render, errorUrl string) *model.UserInfo {
	result, errorMsg := CheckValidUser(form)
	retInfos := make(map[string]string)
	retInfos["userName"] = form.Get("username")
	retInfos["password"] = "******"
	if !utils.IsEmptyStr(errorMsg) {
		retInfos["errorInfo"] = "没有配置信息，请找管理员"
		re.HTML(200, errorUrl, retInfos)
	}
	return result
}
