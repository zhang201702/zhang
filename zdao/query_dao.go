package dao

import (
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
	"github.com/zhang201702/zhang/zlog"
	"strings"
)

type QueryDao struct {
	db gdb.DB
}

type Query struct {
	dao  *QueryDao
	Sql  string
	Args []interface{}
}

func (dao *QueryDao) Query(query string, args ...interface{}) g.List {

	dao.db.SetDebug(true)
	q, err := dao.db.GetAll(query, args...)
	if err != nil {
		zlog.Error(err, "query 异常")
	}
	return q.List()
}
func (dao *QueryDao) QueryOne(query string, args ...interface{}) g.Map {

	dao.db.SetDebug(true)
	q, err := dao.db.GetOne(query, args...)
	if err != nil {
		zlog.Error(err, "query 异常")
	}
	return q.ToMap()
}

func (dao *QueryDao) Condition(sql string, objects []interface{}, data interface{}, sqlConfName, sqlConf string) (string, []interface{}) {
	if data != nil {
		objects = append(objects, data)
		sql = strings.ReplaceAll(sql, "{{"+sqlConfName+"}}", sqlConf)
	}
	return sql, objects
}
func (dao *QueryDao) QueryEntity(result interface{}, query string, args ...interface{}) {

	dao.db.SetDebug(true)
	err := dao.db.GetStructs(result, query, args...)

	if err != nil {
		zlog.Error(err, "queryEntity 异常")
	}
}

func (dao *QueryDao) CreateQuery(sql string, args ...interface{}) *Query {
	return &Query{
		Sql: sql, Args: args, dao: dao,
	}
}
func (query *Query) And(where string, args ...interface{}) {
	if len(args) == 0 || args[0] == nil {
		return
	}
	if where != "" {
		query.Sql += " AND " + where
	}
	query.Args = append(query.Args, args...)
}

func (query *Query) GroupBy(sql string) string {
	return query.Append("GROUP BY " + sql)
}

func (query *Query) OrderBy(sql string) string {
	return query.Append("ORDER BY " + sql)
}

func (query *Query) Append(sql string) string {
	query.Sql += " " + sql
	return query.Sql
}
func (query *Query) ToSQL() (string, []interface{}) {
	return query.Sql, query.Args
}
func (query *Query) ToList() g.List {
	return query.dao.Query(query.Sql, query.Args...)
}
func (query *Query) ToMap() g.Map {
	return query.dao.QueryOne(query.Sql, query.Args...)
}
