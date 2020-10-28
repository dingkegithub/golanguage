package singlelist

import (
	"sync"
)

type NodeItem struct {
	Data interface{}
	Next *NodeItem
}

type SingleList struct {
	Head  *NodeItem
	size  int
	mutex sync.Mutex
	Opts  *listOption
}

func NewSingleList(opts ...Option) *SingleList {
	lOpt := &listOption{}
	lOpt.apply(opts...)

	return &SingleList{
		Opts: lOpt,
	}
}

func (sl *SingleList) Insert(data interface{}) bool {

	sl.mutex.Lock()
	defer sl.mutex.Unlock()

	node := &NodeItem{
		Data: data,
		Next: nil,
	}

	if sl.Head == nil {
		sl.Head = node
		sl.size = 1
	} else {
		tmp := sl.Head

		for tmp.Next != nil {
			tmp = tmp.Next
		}
		tmp.Next = node
		sl.size += 1
	}

	return true
}

func (sl *SingleList) Delete(data interface{}) (interface{}, error) {
	sl.mutex.Lock()
	defer sl.mutex.Unlock()

	var pre *NodeItem
	tmp := sl.Head

	for tmp != nil {
		res, err := sl.Opts.Compare(tmp.Data, data)
		if err != nil {
			return nil, err
		}

		if res == 0 {
			if pre == nil {
				sl.Head = sl.Head.Next
				sl.size -= 1
				return tmp.Data, nil
			}

			pre.Next = tmp.Next
			sl.size -= 1
			return tmp.Data, nil
		} else {
			pre = tmp
			tmp = tmp.Next
		}
	}

	return nil, nil
}

func (sl *SingleList) Query(data interface{}) (int, error) {
	i := 0
	tmp := sl.Head

	for tmp != nil {
		res, err := sl.Opts.Compare(tmp.Data, data)
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

func (sl *SingleList) Index(i int) interface{} {
	if i > sl.size {
		return nil
	}

	idx := 0
	tmp := sl.Head
	for tmp != nil {
		if idx == i {
			return tmp.Data
		}
		idx += 1
		tmp = tmp.Next
	}

	return nil
}

func (sl *SingleList) Size() int {
	return sl.size
}
