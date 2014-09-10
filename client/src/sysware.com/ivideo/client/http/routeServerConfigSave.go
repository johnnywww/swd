package http

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"net/http"
	"sysware.com/ivideo/log"
	"sysware.com/ivideo/model"
	shttp "sysware.com/ivideo/server/http"
	"sysware.com/ivideo/service"
)

type routeServerConfigSave struct {
}

func (routeServerConfigSave *routeServerConfigSave) Routes(m *martini.ClassicMartini) {
	m.Post("/server/saveconfig", func(w http.ResponseWriter, r *http.Request, re render.Render) string {
		log.WriteLog("addr: /server/saveconfig")
		shttp.SetResponseJsonHeader(w)
		r.ParseForm()
		err := service.NewServerConfigSaveRequestHandle().Save(r)
		if nil != err {
			return model.GetErrorObjDtoJson(err)
		}
		return model.GetDataDtoJson("")
	})
}
