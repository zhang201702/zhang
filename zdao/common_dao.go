package zdao

import (
	"github.com/gogf/gf/database/gdb"
	"github.com/zhang201702/zhang/zlog"
)

type CommonDao[T any] struct {
	gdb.DB
	Name string
}

type PageResult[T any] struct {
	List     []*T `json:"list"`
	RowCount int  `json:"rowCount"`
}

func NewCommonDao[T any](db gdb.DB, name string) *CommonDao[T] {
	return &CommonDao[T]{Name: name, DB: db}
}

func (da *CommonDao[T]) Insert(data *T) (result interface{}, err error) {
	tableName := da.Name
	table := da.DB.Model(tableName)

	//data.CreateTime = z.Now()
	result, err = table.Data(data).Insert()
	if err != nil {
		zlog.LogError(err, tableName, "添加异常", data, result)
	}
	return result, err
}

func (da *CommonDao[T]) Update(data *T, where map[string]interface{}) (result interface{}, err error) {
	tableName := da.Name
	table := da.DB.Model(tableName)

	//data.CreateTime = z.Now()
	result, err = table.Data(data).Where(where).Update()
	if err != nil {
		zlog.LogError(err, tableName, "修改异常", data, result)
	}
	return result, err
}

func (da *CommonDao[T]) UpdateMap(data map[string]interface{}, where map[string]interface{}) (result interface{}, err error) {
	tableName := da.Name
	table := da.DB.Model(tableName)

	//data.CreateTime = z.Now()
	result, err = table.Update(data, where)
	if err != nil {
		zlog.LogError(err, tableName, "修改异常", data, result)
	}
	return result, err
}

func (da *CommonDao[T]) Fetch(where map[string]interface{}) *T {
	tableName := da.Name
	table := da.DB.Model(tableName)
	var data = (*T)(nil)
	err := table.Where(where).Scan(&data)
	if err != nil {
		zlog.LogError(err, "Fetch error", where)
	}
	return data
}

func (da *CommonDao[T]) GetList(where map[string]interface{}, orderBy ...string) []*T {
	tableName := da.Name
	table := da.DB.Model(tableName)
	list := make([]*T, 0)
	err := table.Where(where).Order(orderBy...).Scan(&list)
	if err != nil {
		zlog.LogError(err, "GetList error", where)
	}
	return list
}

func (da *CommonDao[T]) GetPage(where map[string]interface{}, page, size int, orderBy ...string) PageResult[T] {
	tableName := da.Name
	table := da.DB.Model(tableName)
	list := make([]*T, 0)
	result := PageResult[T]{List: list}
	table = table.Where(where)
	rowCount, err := table.Count()
	if err != nil {
		zlog.LogError(err, "GetPage.Count error", where)
		return result
	}
	result.RowCount = rowCount
	err = table.Order(orderBy...).Page(page, size).Scan(&list)
	if err != nil {
		zlog.LogError(err, "GetPage error", where, page, size)
	}
	return result
}
