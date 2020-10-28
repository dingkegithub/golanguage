package doublelist

import (
	"sync"

	"github.com/dingkegithub/golanguage/datastrcture/list"
)

type NodeItem struct {
	Data interface{}
	Next *NodeItem
	Pre  *NodeItem
}

type DoubleList struct {
	Head  *NodeItem
	Tail  *NodeItem
	size  int
	mutex sync.Mutex
	Opts  *list.ListOption
}

func NewDoubleList(opts ...list.Option) *DoubleList {
	lOpt := &list.ListOption{}
	lOpt.Apply(opts...)

	return &DoubleList{
		Opts: lOpt,
	}
}

func (dl *DoubleList) Insert(data interface{}) bool {

	dl.mutex.Lock()
	defer dl.mutex.Unlock()

	node := &NodeItem{
		Data: data,
		Next: nil,
		Pre:  nil,
	}

	// keep head->pre = nil, tail->next = nil
	if dl.Head == nil {
		dl.Head = node
		dl.Tail = dl.Head
		dl.size += 1
		return true
	}

	dl.Tail.Next = node
	node.Pre = dl.Tail
	dl.Tail = node
	dl.size += 1

	return true
}

func (dl *DoubleList) Delete(data interface{}) (interface{}, error) {
	dl.mutex.Lock()
	defer dl.mutex.Unlock()

	tmp := dl.Head

	for tmp != nil {
		res, err := dl.Opts.Compare(tmp.Data, data)
		if err != nil {
			return nil, err
		}

		if res == 0 {
			if tmp == dl.Head {
				if dl.Head == dl.Tail {
					dl.Head, dl.Tail = nil, nil
				} else {
					dl.Head = dl.Head.Next
					dl.Head.Pre = nil
				}
			} else {
				tmp.Pre.Next = tmp.Next
				if tmp.Next == nil {
					dl.Tail = tmp.Pre
				} else {
					tmp.Next.Pre = tmp.Pre
				}
			}

			dl.size -= 1
			return tmp.Data, nil
		} else {
			tmp = tmp.Next
		}
	}

	return nil, nil
}

func (dl *DoubleList) Query(data interface{}) (int, error) {
	i := 0
	tmp := dl.Head

	for tmp != nil {
		res, err := dl.Opts.Compare(tmp.Data, data)
		if err != nil {
			return -1, err
		}

		if res == 0 {
			return i, nil
		}
		i += 1
		tmp = tmp.Next
	}

	return -1, nil
}

func (dl *DoubleList) Index(i int) interface{} {
	if i > dl.size {
		return nil
	}

	idx := 0
	tmp := dl.Head
	for tmp != nil {
		if idx == i {
			return tmp.Data
		}
		idx += 1
		tmp = tmp.Next
	}

	return nil
}

func (dl *DoubleList) Size() int {
	return dl.size
}
