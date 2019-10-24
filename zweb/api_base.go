package zweb

import (
	"github.com/gogf/gf/net/ghttp"
	"github.com/zhang201702/zhang/z"
	"github.com/zhang201702/zhang/zlog"
	"log"
	"runtime"
)

type ApiBase struct {
}

func (api *ApiBase) Ok(data interface{}) z.Map {
	return z.Map{
		"status": true,
		"data":   data,
	}
}

func (api *ApiBase) Error(msg string, err error) z.Map {

	pc, file, line, ok := runtime.Caller(1)
	log.Print(pc, file, line, ok)

	if err != nil {
		zlog.LogError(err)
		if msg == "" {
			msg = err.Error()
		}
	}
	return z.Map{
		"status": false,
		"msg":    msg,
	}
}

func (api *ApiBase) OkResult(data interface{}, r *ghttp.Request) error {
	return r.Response.WriteJson(api.Ok(data))
}

func (api *ApiBase) ErrorResult(msg string, error error, r *ghttp.Request) error {
	return r.Response.WriteJson(api.Error(msg, error))
}

func (api *ApiBase) Result(data interface{}, err error, r *ghttp.Request) error {
	if err != nil {
		return api.ErrorResult("", err, r)
	}
	return api.OkResult(data, r)
}

func (api *ApiBase) Result2(result *z.Result, r *ghttp.Request) error {
	return r.Response.WriteJson(*result)
}
func (api *ApiBase) GetToken(r *ghttp.Request) string {
	return r.Header.Get("token")
}
