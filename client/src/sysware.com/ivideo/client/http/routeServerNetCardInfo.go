package http

import (
	"github.com/go-martini/martini"
	"net/http"
	"sysware.com/ivideo/log"
	"sysware.com/ivideo/model"
	shttp "sysware.com/ivideo/server/http"
	"sysware.com/ivideo/service"
)

type RouteServerNetCardInfo struct {
}

func (routeServerNetCardInfo *RouteServerNetCardInfo) Routes(m *martini.ClassicMartini) {
	m.Get("/server/info/netcard", func(w http.ResponseWriter, r *http.Request) string {
		log.WriteLog("addr: /server/info/netcard")
		shttp.SetResponseJsonHeader(w)
		result, err := service.NewGetNetStatisticsInfo().GetInfo()
		if nil != err {
			return model.GetErrorObjDtoJson(err)
		}
		return model.GetDataDtoJson(result)
	})
}
