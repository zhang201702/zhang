package zdao

import (
	"database/sql"
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/util/gconv"
	"github.com/zhang201702/zhang/utils"
	"github.com/zhang201702/zhang/z"
	"github.com/zhang201702/zhang/zconfig"
	"github.com/zhang201702/zhang/zlog"
	"reflect"
	"strings"
)

type QueryDao struct {
	gdb.DB
}

type Query struct {
	dao       *QueryDao
	Sql       string
	Args      []interface{}
	haveWhere bool
}

func (dao *QueryDao) Query(query string, args ...interface{}) g.List {

	q, err := dao.GetAll(query, args...)
	if err != nil && zconfig.Debug {
		zlog.Error(err, "query 异常")
	}
	return q.List()
}
func (dao *QueryDao) QueryValue(query string, args ...interface{}) *g.Var {

	q, err := dao.GetValue(query, args...)
	if err != nil && zconfig.Debug {
		zlog.Error(err, "QueryOne 异常")
	}
	return q
}
func (dao *QueryDao) QueryOne(query string, args ...interface{}) g.Map {

	q, err := dao.GetOne(query, args...)
	if err != nil && zconfig.Debug {
		zlog.Error(err, "QueryOne 异常")
	}
	return q.ToMap()
}
func (dao *QueryDao) QueryStruct(objPointer interface{}, query string, args ...interface{}) error {

	err := dao.GetStruct(objPointer, query, args...)
	if err == sql.ErrNoRows {
		return err
	}
	if err != nil && zconfig.Debug {
		zlog.Error(err, "QueryStruct 异常")
	}
	return err
}

func (dao *QueryDao) QueryStructs(objPointerSlice interface{}, query string, args ...interface{}) error {

	err := dao.GetStructs(objPointerSlice, query, args...)

	if err == sql.ErrNoRows {
		return err
	}
	if err != nil && zconfig.Debug {
		zlog.Error(err, "QueryStructs 异常")
	}
	return err
}
func (dao *QueryDao) Condition(sql string, objects []interface{}, data interface{}, sqlConfName, sqlConf string) (string, []interface{}) {
	if data != nil {
		objects = append(objects, data)
		sql = strings.ReplaceAll(sql, "{{"+sqlConfName+"}}", sqlConf)
	}
	return sql, objects
}
func (dao *QueryDao) QueryEntity(result interface{}, query string, args ...interface{}) {

	err := dao.GetStructs(result, query, args...)

	if err != nil {
		zlog.Error(err, "queryEntity 异常")
	}
}

func (dao *QueryDao) CreateQuery(sql string, args ...interface{}) *Query {
	return &Query{
		Sql: sql, Args: args, dao: dao, haveWhere: false,
	}
}
func (query *Query) Select(fields string) *Query {
	query.Sql += utils.String("SELECT ", fields)
	return query
}
func (query *Query) From(from ...interface{}) *Query {
	if len(from) == 0 {
		return query
	}
	firstP := from[0]
	switch firstP.(type) {
	case string:
		query.Sql += utils.String(" FROM ", from)
	case *Query:
		temp := firstP.(*Query)
		alas := "t"
		if len(from) >= 2 {
			alas = gconv.String(from[1])
		}
		query.Sql += utils.String(" FROM ( ", temp.Sql, " ) ", alas)
		query.Args = append(query.Args, temp.Args...)
	default:
	}
	return query
}
func (query *Query) Where(where string, args ...interface{}) *Query {
	if query.haveWhere {
		return query.And(where, args...)
	}
	if len(args) == 0 || args[0] == nil {
		return query
	}
	if where != "" {
		query.Sql += " WHERE " + where
	}
	query.Args = append(query.Args, args...)
	query.haveWhere = true
	return query
}

func (query *Query) Eq(column string, args interface{}) *Query {
	if args == nil || args == "" {
		return query
	}
	return query.And(column+" = ?", args)
}
func (query *Query) Like(column string, args interface{}) *Query {
	if args == nil || args == "" {
		return query
	}
	return query.And(column+" like ?", z.String(args, "%"))
}

func (query *Query) In(column string, args interface{}) *Query {

	if args == nil || args == "" {
		return query
	}
	s := reflect.ValueOf(args)
	if s.Len() == 0 {
		return query
	}
	return query.And(column+" in (?)", args)
}

func (query *Query) And(where string, args ...interface{}) *Query {
	if len(args) == 0 || args[0] == nil || args[0] == "" {
		return query
	}
	if where != "" {
		query.Sql += " AND " + where
	}
	query.Args = append(query.Args, args...)
	return query
}
func (query *Query) AndDefault(defaultWhere, where string, args ...interface{}) {
	if len(args) == 0 || args[0] == nil || args[0] == "" {
		if query.haveWhere || strings.Contains(strings.ToLower(query.Sql), "where") {
			query.Append(" AND " + defaultWhere)
		} else {
			query.Append(" WHERE " + defaultWhere)
			query.haveWhere = true
		}

		return
	}
	query.And(where, args...)
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
func (query *Query) ToStructs(objPointerSlice interface{}) error {
	return query.dao.QueryStructs(objPointerSlice, query.Sql, query.Args)
}

func (query *Query) ToStruct(objPointer interface{}) error {
	return query.dao.QueryStruct(objPointer, query.Sql, query.Args)
}
func (query *Query) QueryValue() *g.Var {
	return query.dao.QueryValue(query.Sql, query.Args)
}
