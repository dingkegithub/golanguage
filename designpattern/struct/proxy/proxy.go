package proxy

import (
	"fmt"
	"sync"
)

//
// 数据库类型
//
type DbType string

const (
	DbTypeMysql DbType = "mysql"
	DbTypeSSQL  DbType = "SqlServer"
)

func (d DbType) String() string {
	return string(d)
}

// 代理器 DbOrm:
// 代理各种数据库，提供缓存、对象延迟创建功能
type DbOrm struct {
	db          DbInterface
	cacheEnable bool
	cache       map[string]interface{}
	once        sync.Once
	dbtype      DbType
}

func NewDbOrm(cacheEable bool, t DbType) DbInterface {
	orm := &DbOrm{}
	orm.cacheEnable = cacheEable
	orm.dbtype = t
	orm.cache = make(map[string]interface{})
	return orm
}

// 延迟创建代理数据库，单例模式保证只创建一次
// 工厂模式保证创建准确的代理数据库
func (do *DbOrm) initDb() {
	do.once.Do(func() {
		switch do.dbtype {
		case DbTypeMysql:
			do.db = NewMysql()
		case DbTypeSSQL:
			fmt.Println("current orm version is unsupport")
			panic("current orm version is unsupport")
		default:
			do.db = NewMysql()
		}
	})
}

func (do *DbOrm) Create(record string) {
	if do.db == nil {
		do.initDb()
	}

	fmt.Println("Db Proxy Orm Create")
	do.db.Create(record)
}

func (do *DbOrm) Query(record string) interface{} {
	fmt.Println("Db Proxy Orm Query")

	if do.db == nil {
		do.initDb()
	}

	if do.cacheEnable {
		if v, ok := do.cache[record]; ok {
			fmt.Println("Db Proxy Orm Query from cache")
			return v
		}
	}

	v := do.db.Query(record)
	if do.cacheEnable {
		fmt.Println("Db Proxy Orm Query cache result")
		do.cache[record] = v
	}

	fmt.Println("Db Proxy Orm Query db result")
	return v
}

func (do *DbOrm) Delete(record string) {
	fmt.Println("Db Proxy Orm Delete")

	if do.db == nil {
		do.initDb()
	}

	if do.cacheEnable {
		if _, ok := do.cache[record]; ok {
			fmt.Println("Db Proxy Orm Query from cache")
			delete(do.cache, record)
		}
	}

	do.db.Delete(record)
}
