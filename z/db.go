package z

import (
	"github.com/gogf/gf/database/gdb"
	"github.com/zhang201702/zhang/zdao"
)

/**
获取 DB对象.
  默认
*/
func DB(names ...string) gdb.DB {
	return zdao.DB(names...)

}

/**
获取 DB对象.
  name : 配置名称,默认null
*/
func GetDB(name string) gdb.DB {
	return DB(name)
}
