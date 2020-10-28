package proxy

import (
	"testing"
)

func TestNewDbOrm(t *testing.T) {
	var db DbInterface
	db = NewDbOrm(true, DbTypeMysql)

	db.Create("record")
	res := db.Query("record")
	t.Log("query result: ", res)
	db.Delete("record")
}
