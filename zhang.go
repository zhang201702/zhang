package zhang

import (
  "github.com/gogf/gf/frame/g"
  "github.com/gogf/gf/net/ghttp"
  "github.com/gogf/gf/os/gfile"
  "github.com/zhang201702/zhang/zconfig"
)

type ServerGF struct {
  *ghttp.Server
}

func Default() *ServerGF {
  server := &ServerGF{
    Server: g.Server(),
  }

  crt := zconfig.Conf.GetString("https.crt")
  key := zconfig.Conf.GetString("https.key")

  if crt != "" && key != "" {
    port4https := zconfig.Conf.GetInt("https.port", 443)
    server.SetHTTPSPort(port4https)
    server.EnableHTTPS(crt, key)
  }

  htmlPath := zconfig.Conf.GetString("html")
  if htmlPath == "" {
    htmlPath = "html"
  }
  if gfile.IsDir(htmlPath) {
    server.SetServerRoot(htmlPath)
  } else {
    server.BindHandler("/", func(r *ghttp.Request) {
      r.Response.Write("welcome api!!!")
    })
  }

  server.SetRouteOverWrite(true)

  server.BindHandler("/health", func(r *ghttp.Request) {
    r.Response.Write("ok")
  })
  server.BindHandler("/info", func(r *ghttp.Request) {
    r.Response.Write("ok")
  })
  port := zconfig.Conf.GetInt("port", 80)
  server.SetPort(port)
  return server
}

func (server *ServerGF) NoCache() {
  server.BindMiddlewareDefault(func(r *ghttp.Request) {
    r.Response.Header().Set("Cache-Control", "no-cache")
  })
}
