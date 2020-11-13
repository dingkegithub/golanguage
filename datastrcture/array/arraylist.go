package array

import (
	"sync"

	"github.com/dingkegithub/golanguage/datastrcture"
	"github.com/dingkegithub/golanguage/pkg/param"
)

type arrayListOption struct {
	capacity uint64
}

func withCap(c uint64) param.ParaOption {
	return func(o interface{}) {
		if opt, ok := o.(*arrayListOption); ok {
			if c < 0 {
				return
			}

			opt.capacity = c
		}
	}
}

type ArrayList struct {
	size  uint64
	mutex sync.RWMutex
	data  []interface{}
	opts  *arrayListOption
}

func NewArrayList(options ...param.ParaOption) *ArrayList {

	opts := &arrayListOption{
		capacity: 10,
	}

	if len(options) > 0 {
		for _, option := range options {
			option(opts)
		}
	}

	return &ArrayList{
		size: 0,
		opts: opts,
		data: make([]interface{}, 0, opts.capacity),
	}
}

func (al *ArrayList) Size() uint64 {
	al.mutex.RLock()
	defer al.mutex.RUnlock()

	return al.size
}

func (al *ArrayList) Add(item interface{}) {
	al.mutex.Lock()
	defer al.mutex.Unlock()

	al.data = append(al.data, item)
	al.size += 1
}

func (al *ArrayList) Get(idx uint64) (interface{}, error) {
	al.mutex.RLock()
	defer al.mutex.RUnlock()

	if idx < al.size {
		return al.data[idx], nil
	}

	return nil, datastrcture.ErrorBoundary
}

func (al *ArrayList) Set(idx uint64, item interface{}) error {
	al.mutex.Lock()
	defer al.mutex.Unlock()

	if idx < al.size {
		al.size += 1
		newArray := make([]interface{}, 0, al.size)

		newArray = append(newArray, al.data[:idx]...)
		newArray = append(newArray, item)
		newArray = append(newArray, al.data[idx:al.size-1]...)

		al.data = newArray
		return nil
	}

	return datastrcture.ErrorBoundary
}

func (al *ArrayList) Remove(idx uint64) (interface{}, error) {
	al.mutex.Lock()
	defer al.mutex.Unlock()

	if idx < al.size {
		newArray := make([]interface{}, 0, al.size)
		newArray = append(newArray, al.data[:idx]...)

		if idx+1 < al.size {
			newArray = append(newArray, al.data[idx:])
		}

		data := al.data[idx]
		al.data = newArray

		return data, nil
	}

	return nil, datastrcture.ErrorBoundary
}
