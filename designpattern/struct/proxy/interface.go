package proxy

// 代理和被代理对象公共的接口
type DbInterface interface {
	Create(string)

	Query(string) interface{}

	Delete(string)
}
