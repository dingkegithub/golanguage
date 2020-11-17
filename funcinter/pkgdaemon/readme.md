
<h2 id='1'> 1. 包的调用流程 </h2>

- 调用关系：main--->pkg2--->pkg1, 初始化次序应该是 pkg1-->pkg2---->main

- 一个包内的init 将顺序调用

- 在main函数之前执行的所有代码都在同样一个 goroutine

- 其它 goroutine 只有进入 main 之后才可能被执行, 比如pkg3 中的第二个init 启动了一个goroutine，但是并没有执行知道进入main

```
pkg1 init:  const-hello var-hello
pkg2 init:  pkg2-const-world pkg2-var-world
pkg3's first init
pkg3's second init

# pkg3 第二个init调用的携程
pkg3's second init invoked goruntine

main init:  main-var-main main-const-main
main running
pkg1 PrintName:  const-hello var-hello
pkg2 PrintName:  pkg2-const-world pkg2-var-world

// 携程进入main函数之后再执行，不再init中执行
pkg3 PrintName: ok
main exit
```
