package zhang

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

type ServerGF struct {
	*ghttp.Server
}

func Default() *ServerGF {
	server := &ServerGF{
		Server: g.Server(),
	}
	server.SetServerRoot("html")
	server.SetPort(80)
	server.BindHandler("/health", func(r *ghttp.Request) {
		r.Response.Write("ok")
	})
	return server
}
