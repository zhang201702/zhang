package main

import (
	"github.com/zhang201702/zhang"
	"github.com/zhang201702/zhang/z"
)

// Entity
type Entity struct {
	Id uint64 `orm:"id,omitempty" json:"id,omitempty"` // id
}

func main() {

	// 数据库操作
	//db := z.DB("nft")
	//db.SetDebug(true)
	//entity := zdao.NewCommonDao[Entity](db, "entity")
	//find := entity.GetList(nil)
	//zlog.Log(find)

	// redis操作
	//r := z.Redis("abc")
	//a, err := r.Do("GET", "test")
	//zlog.LogError(err, a)

	s := zhang.Default()
	z.OpenBrowse(z.GetUrl())
	s.Run()
}
