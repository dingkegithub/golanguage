package classfunc

import "sync"

type Pool struct {
	// 组合匿名实现继承
	sync.Mutex
	total int
	idle  int
}

func (p *Pool) GetTotal() int {
	// 访问父类的方法
	p.Lock()
	defer p.Unlock()
	return p.total
}

func main() {
	p := new(Pool)
	p.GetTotal()
}
