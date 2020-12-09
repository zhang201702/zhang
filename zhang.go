package zhang

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/zhang201702/zhang/zconfig"
)

type ServerGF struct {
	*ghttp.Server
}

func Default() *ServerGF {
	server := &ServerGF{
		Server: g.Server(),
	}

	server.SetServerRoot("html")

	server.BindHandler("/health", func(r *ghttp.Request) {
		r.Response.Write("ok")
	})
	server.BindHandler("/info", func(r *ghttp.Request) {
		r.Response.Write("ok")
	})
	port := zconfig.Conf.GetInt("port")
	if port == 0 {
		port = 80
	}
	server.SetPort(port)
	server.Server.Start()

	return server
}
