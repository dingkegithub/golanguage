package queueandstack

import (
	"fmt"
	"sync"
)

type QueueOptionFunc func(*queueOption)

func WithCap(c int64) QueueOptionFunc {
	return func(qp *queueOption) {
		qp.defaultCap = c
	}
}

func WithScale() QueueOptionFunc {
	return func(qp *queueOption) {
		qp.autoScale = true
	}
}

type queueOption struct {
	autoScale  bool
	defaultCap int64
}

func (qp *queueOption) apply(params ...QueueOptionFunc) {
	for _, p := range params {
		p(qp)
	}
}

type Queue struct {
	mutex   sync.RWMutex
	curCap  int64
	size    int64
	front   int64
	rear    int64
	auto    bool
	options *queueOption
	q       []interface{}
}

var (
	ErrQueueFull  error = fmt.Errorf("queue full")
	ErrQueueEmpty error = fmt.Errorf("queue empty")
)

func New(opts ...QueueOptionFunc) *Queue {
	qo := &queueOption{
		autoScale:  false,
		defaultCap: 5,
	}

	qo.apply(opts...)

	return &Queue{
		curCap:  qo.defaultCap,
		size:    0,
		rear:    -1,
		front:   0,
		options: qo,
		q:       make([]interface{}, qo.defaultCap, qo.defaultCap),
	}
}

func (q *Queue) Put(data interface{}) error {

	isFull := q.Full()
	q.mutex.Lock()
	defer q.mutex.Unlock()

	if isFull {
		if q.options.autoScale {
			scaleCap := q.curCap * 3 / 10
			if scaleCap < 3 {
				scaleCap = 3
			}
			scalQ := make([]interface{}, scaleCap)
			q.q = append(q.q, scalQ...)
			q.curCap += scaleCap
		} else {
			return ErrQueueFull
		}
	}

	q.rear = (q.rear + 1) % q.curCap
	q.q[q.rear] = data
	q.size += 1
	return nil
}

func (q *Queue) Get() (interface{}, error) {
	is := q.Empty()
	q.mutex.Lock()
	defer q.mutex.Unlock()

	// adjust cap
	if q.options.autoScale && q.curCap > 2*q.options.defaultCap && q.front > (q.curCap*7/10) {
		// ToDo
	}

	if is {
		return nil, ErrQueueEmpty
	}

	data := q.q[q.front]
	q.q[q.front] = nil
	q.front = (q.front + 1) % q.curCap
	q.size -= 1

	return data, nil
}

func (q *Queue) Size() int64 {
	q.mutex.RLock()
	defer q.mutex.RUnlock()
	return q.size
}

func (q *Queue) Cap() int64 {
	q.mutex.RLock()
	defer q.mutex.RUnlock()
	return q.curCap
}

func (q *Queue) Empty() bool {
	q.mutex.RLock()
	defer q.mutex.RUnlock()
	return q.size == 0
}

func (q *Queue) Full() bool {
	q.mutex.RLock()
	defer q.mutex.RUnlock()
	if q.size >= q.curCap {
		return true
	}

	return false
}
