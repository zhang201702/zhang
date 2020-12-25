package zdao

import (
	"github.com/gogf/gf/frame/g"
	"github.com/zhang201702/zhang/zlog"
	"testing"
)

func TestQuery(t *testing.T) {
	query := &Query{haveWhere: false}
	p := g.Map{"id": 23}
	query.Select("*").From("a").Where("id = ?", p["id"])

	query2 := &Query{haveWhere: false}

	query2.Select("*").From(query).Where("id = ?", p["id"])
	zlog.Log(query.Sql)
	zlog.Log(query2.Sql)
}
