package http

import (
	"github.com/go-martini/martini"
	"net/http"
	"sysware.com/ivideo/log"
	"sysware.com/ivideo/model"
	shttp "sysware.com/ivideo/server/http"
	"sysware.com/ivideo/service"
)

type RouteServerMemInfo struct {
}

func (routeServerMemInfo *RouteServerMemInfo) Routes(m *martini.ClassicMartini) {
	m.Get("/server/info/mem", func(w http.ResponseWriter, r *http.Request) string {
		log.WriteLog("addr: /server/info/mem")
		shttp.SetResponseJsonHeader(w)
		mem, err := service.NewGetMemInfo().GetInfo()
		if nil != err {
			return model.GetErrorObjDtoJson(err)
		}
		memFree := mem.MemFree + mem.Buffers + mem.Cached
		memUsed := mem.MemTotal - memFree
		return model.GetDataDtoJson([]interface{}{mem.MemTotal, memUsed, memFree})
	})
}
