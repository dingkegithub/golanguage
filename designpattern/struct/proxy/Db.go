package proxy

import (
	"fmt"
)

// 被代理对象
// 即 Mysql, 它将实现所有公共接口
type Mysql struct {
}

func NewMysql() DbInterface {
	fmt.Println("Db api init")
	return &Mysql{}
}

func (do *Mysql) Create(record string) {
	fmt.Println("Db api Create record", record)
}

func (do *Mysql) Query(record string) interface{} {
	fmt.Println("Db api Query")
	return fmt.Sprintf("Record:%s", record)
}

func (do *Mysql) Delete(record string) {
	fmt.Println("Db api Delete:", record)
}
