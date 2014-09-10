package http

import (
	"github.com/martini-contrib/render"
	"net/url"
	"sysware.com/ivideo/model"
	shttp "sysware.com/ivideo/server/http"
	"sysware.com/ivideo/utils"
)

func StartHttp() {
	clientRouteHandle := &model.RouteHandles{Handles: []model.RouteHandle{&RouteRoot{}, &RouteLogin{}, &RouteLogout{}, &shttp.RouteNotFound{}, &shttp.RouteServerSetup{}, &RouteServer{}, &shttp.RouteUser{}}}
	shttp.StartHttp(clientRouteHandle)
}

func checkLoginUser(form url.Values, re render.Render, errorUrl string) *model.UserInfo {
	result, errorMsg := shttp.CheckValidUser(form)
	retInfos := make(map[string]string)
	retInfos["userName"] = form.Get("username")
	retInfos["password"] = "******"
	if !utils.IsEmptyStr(errorMsg) {
		retInfos["errorInfo"] = "没有配置信息，请找管理员"
		re.HTML(200, errorUrl, retInfos)
	}
	return result
}

func getUniqueServerName() string {
	result := ""
	setupInfo := model.SetupInfo{}
	if !model.NewSetupHandle().GetSetupInfo(&setupInfo) {
		return result
	}
	if len(setupInfo.Servers) == 1 {
		result = setupInfo.Servers[0].ServerName
	}
	return result
}

func getFirstServerName() string {
	return getUniqueServerName()
}

func getStrMapRetInfo() map[string]string {
	result := make(map[string]string)
	result["serverName"] = getFirstServerName()
	return result
}
