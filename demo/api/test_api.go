package api

import (
	"github.com/gogf/gf/net/ghttp"
	"github.com/zhang201702/zhang/zweb"
)

type TestApi struct {
	*zweb.ApiBase
}

func (api *TestApi) Register(server *ghttp.Server) {
	server.BindHandler("test", api.test)
}
func (api *TestApi) test(r *ghttp.Request) {
	_ = api.OkResult("test", r)
}
