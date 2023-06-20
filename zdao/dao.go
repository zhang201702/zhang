package zdao

import (
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
	"github.com/zhang201702/zhang/zconfig"
	"github.com/zhang201702/zhang/zlog"
	"strings"
	"sync"
)

var dbDefault gdb.DB = nil
var once = sync.Once{}

func initDB() {

	defaultDbInfoName := "db"
	dbInfo := zconfig.Get(defaultDbInfoName)
	if dbInfo == nil {
		return
	}
	getLink := func(link string) (linkType string, linkValue string) {
		firstIndex := strings.Index(link, ":")
		linkType = link[0:(firstIndex)]
		linkValue = link[firstIndex+1:]
		return linkType, linkValue
	}
	switch dbInfo.(type) {
	case map[string]interface{}:
		myMap := dbInfo.(map[string]interface{})
		for name, link := range myMap {
			dbType, dbLink := getLink(link.(string))
			gdb.SetConfigGroup(name, gdb.ConfigGroup{
				gdb.ConfigNode{
					Type: dbType,
					Link: dbLink,
				},
			})
		}
	case string:
		dbType, dbLink := getLink(dbInfo.(string))
		gdb.AddDefaultConfigGroup(gdb.ConfigGroup{
			gdb.ConfigNode{
				Type: dbType,
				Link: dbLink,
			},
		})
	}
}

/**
获取 DB对象.
  默认
*/
func DB(names ...string) gdb.DB {

	once.Do(initDB)
	if len(names) == 0 {
		if dbDefault == nil {
			dbDefault = g.DB()
		}
		return dbDefault
	}
	name := names[0]

	db, err := gdb.New(name)
	if err != nil {
		zlog.Error(err, "创建db异常,name", name)
		return nil
	}
	return db

}
