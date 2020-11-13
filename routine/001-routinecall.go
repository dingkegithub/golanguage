package main

import (
	"fmt"
	"sync"
)

func main() {
	r2()
}

//
// 执行结果是什么?
//
// 执行结果:
//    每次结果不一样，可能大部分输出10，在携程可能没调度前，i已经变了
//
func r1() {
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {

		wg.Add(1)

		go func() {
			defer func() {
				wg.Done()
			}()
			fmt.Println("i:", i)
		}()
	}

	wg.Wait()
}

//
// 执行结果是什么?
//
// 执行结果: 按不同顺序输出 0~9的数字，i 通过携程函数参数传入
//           保持携程私有，外部不影响携程内部
//
func r2() {
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {

		wg.Add(1)

		go func(i int) {
			defer func() {
				wg.Done()
			}()
			fmt.Println("i:", i)
		}(i)
	}

	wg.Wait()
}
