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
	port := zconfig.Conf.GetInt("port")
	if port == 0 {
		port = 80
	}
	server.SetServerRoot("html")
	server.SetPort(port)
	server.BindHandler("/health", func(r *ghttp.Request) {
		r.Response.Write("ok")
	})
	return server
}
