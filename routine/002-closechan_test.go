package main

import "testing"

func TestCloseChan(t *testing.T) {
	ch := make(chan int)
	ch1 := make(chan bool)
	ch2 := make(chan string)
	ch3 := make(chan interface{})

	var nilchan chan interface{}

	close(ch)
	close(ch1)
	close(ch2)
	close(ch3)

	// 读取关闭的channel
	read := <-ch
	if read != 0 {
		t.Log("close chan read want: 0, but: ", read)
		t.FailNow()
	}

	read1 := <-ch1
	if read1 != false {
		t.Log("close chan read want: false, but: ", read1)
		t.FailNow()
	}

	read2 := <-ch2
	if read2 != "" {
		t.Log("close chan read want: '', but: ", read2)
		t.FailNow()
	}

	read3 := <-ch3
	if read3 != nil {
		t.Log("close chan read want: nil, but: ", read3)
		t.FailNow()
	}

	// 想 nil chan 写数据
	t.Log("writing to nil chan")
	nilchan <- 1111111

	t.Log("reading to nil chan")
	<-nilchan

	t.Log("readed writed to nil chan")
	// panic: send on closed channel
	ch <- 2
}
