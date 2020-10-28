package queueandstack

import "testing"

func TestQueue(t *testing.T) {
	q := New(WithCap(3))
	q.Put(1)
	q.Put(2)
	q.Put(3)

	if !q.Full() {
		t.Error("queue Full failed, want: true, get: ", q.Full())
		t.FailNow()
	}

	err := q.Put(4)
	if err != ErrQueueFull {
		t.Error("queue full failed, want ErrQueueFull, get: ", err)
		t.FailNow()
	}

	if q.Empty() {
		t.Error("queue want: not blank, but: blank")
		t.FailNow()
	}

	for i := 0; i < 3; i++ {
		d, err := q.Get()
		if err != nil {
			t.Error("queue not blank, but get failed", err)
			t.FailNow()
		}

		if d.(int) != i+1 {
			t.Error("queue get data error, want:", i+1, ", but: ", d.(int))
			t.FailNow()
		}
	}

	if !q.Empty() {
		t.Error("queue want: blank, but: not blank")
		t.FailNow()
	}

	_, err = q.Get()
	if err != ErrQueueEmpty {
		t.Error("queue is blank, Get not found error", q.Size())
		t.FailNow()
	}

	q.Put(4)
	q.Put(5)

	for i := 4; i < 6; i++ {
		d, err := q.Get()
		if err != nil {
			t.Error("get data from queue error: ", err)
			t.FailNow()
		}

		if d.(int) != i {
			t.Error("get data want:", i, "but: ", d.(int))
			t.FailNow()
		}
	}
}

func TestAutoQueue(t *testing.T) {
	q := New(WithCap(2), WithScale())
	q.Put(1)
	q.Put(2)
	q.Put(3)

	if q.Cap() != 5 {
		t.Error("queue cap want: ", 5, "but: ", q.Cap())
		t.FailNow()
	}

	q.Put(4)
	q.Put(5)

	if !q.Full() {
		t.Error("queue is not full, but get full", q.Size(), q.Full(), q.Cap())
		t.FailNow()
	}

	q.Put(6)

	if q.Cap() != 8 {
		t.Error("queue cap want: ", 8, "but: ", q.Cap())
		t.FailNow()
	}

	q.Put(7)
	q.Put(8)
	q.Put(9)

	if q.Cap() != 11 {
		t.Error("queue cap want: ", 11, "but: ", q.Cap())
		t.FailNow()
	}

	for i := 0; i < 6; i++ {
		q.Put(10 + i)
	}

	if q.Cap() != 18 {
		t.Error("queue cap want: ", 18, "but: ", q.Cap())
		t.FailNow()
	}

	for i := 0; i < 3; i++ {
		d, err := q.Get()
		if err != nil {
			t.Error("queue not blank, but get failed", err)
			t.FailNow()
		}

		if d.(int) != i+1 {
			t.Error("queue get data error, want:", i+1, ", but: ", d.(int))
			t.FailNow()
		}
	}

	if q.Size() != 12 {
		t.Error("queue size err, want: 12, but:", q.Size())
		t.FailNow()
	}
}
