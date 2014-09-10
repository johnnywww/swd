package http

import (
	"github.com/go-martini/martini"
	"net/http"
	"sysware.com/ivideo/log"
	"sysware.com/ivideo/model"
	shttp "sysware.com/ivideo/server/http"
	"sysware.com/ivideo/service"
)

type RouteServerCpuInfo struct {
}

func (routeServerCpuInfo *RouteServerCpuInfo) Routes(m *martini.ClassicMartini) {
	m.Get("/server/info/cpu", func(w http.ResponseWriter, r *http.Request) string {
		log.WriteLog("addr: /server/info/cpu")
		shttp.SetResponseJsonHeader(w)
		cpuInfos, err := service.NewGetCpuStatisticsInfo().GetInfo()
		if nil != err {
			return model.GetErrorObjDtoJson(err)
		}
		return model.GetDataDtoJson(cpuInfos)
	})
}
