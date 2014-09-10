package http

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"net/http"
	"sysware.com/ivideo/log"
	"sysware.com/ivideo/model"
	shttp "sysware.com/ivideo/server/http"
	"sysware.com/ivideo/service"
	"sysware.com/ivideo/utils"
)

type routeServerConfigRoot struct {
}

func (routeServerConfigRoot *routeServerConfigRoot) Routes(m *martini.ClassicMartini) {
	m.Get("/serverconfig", func(w http.ResponseWriter, r *http.Request, re render.Render) string {
		log.WriteLog("addr: /serverconfig")
		shttp.SetResponseJsonHeader(w)
		r.ParseForm()
		serverName := r.FormValue("servername")
		if utils.IsEmptyStr(serverName) {
			return model.GetErrorDtoJson("没有服务器名称")
		}
		serverInfo, serverSetupInfo, err := service.NewGetServerSetupInfo().GetInfo(serverName)
		if nil != err {
			return model.GetErrorObjDtoJson(err)
		}
		return model.GetDataDtoJson([]interface{}{serverInfo, serverSetupInfo})
	})
}
