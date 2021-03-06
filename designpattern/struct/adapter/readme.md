>
>[1. UML](#1)
>
>[2. 场景](#2)
>
>[3. 设计要点](#3)
>
>[4. 优缺点](#4)
>
>[5. 与其它模式关系](#5)
>

<p style="text-indent:2em">转换对象的接口使其能与其它对象进行交互，适配器隐藏复杂的转换过程。</p>

<p style="text-indent:2em">实现上可以这么做</p>

- 适配器实现与其中一个现有对象兼容的接口
- 现有对象可以使用该接口安全地调用适配器方法
- 适配器方法被调用后将以另一个对象兼容的格式和顺序将请求传递给该对象

> 设计适配器的关键考虑的首先不是适配器怎么设计，而是看清楚这两个对象如何不兼容，谁应该适配谁，然后在根据被适配对象进行设计适配器

<h2 id='1'> 1. UML </h2>

<p style="text-indent:2em">以下面的 UML 图为例子，对象 MacPro 电脑只提供了typc 接口，而现在有一个对象2 输入输出为

usb 信号，usb 协议和 typc 接口协议不兼容，直接将usb 信号不能导入到macpro 里面，因此适配器实现 Signal 接口，该接口

为对象2的接口，就是说适配器实现了被代理对象的接口, 适配器 SignalAdapter 内部通过 Cvt 完成信号的转换，最终将适配器

传递给 MacPro 的 Display 完成正确显示</p>

```
                                                        +---------------+
                                                        | 被适配对象接口|
                                                        |    Signal     |
                                                        +---------------+
                                                        |Output() string|
                                                        +------+--------+
                                                              /|\
                                                               |
                                                      +--------+---------------+
                                                      |                        | 
                   +----------------+        +--------+-------+        +-------+------+
                   |      对象1     |        |     适配器     |        |    对象2     |
                   |      MacPro    |        |  SignalAdapter |        |   UsbSignal  |
                   +----------------+        +----------------+     +--+--------------+
                   |                |        |                |     |  |              |
                   | Display(Signal)|        | us *UsbSignal<>+-----+  | Output() str |
                   +----------------+        | Output() string|        | Input() str  |
                                             | Cvt(string) str|        +--------------+
                                             +----------------+
```

<h2 id='2'> 2. 场景 </h2>

- 当你希望使用某个类， 但是其接口与其他代码不兼容时， 可以使用适配器类

    - 适配器模式允许你创建一个中间层类， 其可作为代码与遗留类、 第三方类或提供怪异接口的类之间的转换器

<h2 id='3'> 3. 设计要点 </h2>

- 确保至少有两个类的接口不兼容

    - 一个无法修改 （通常是第三方、 遗留系统或者存在众多已有依赖的类） 的功能性服务类。

    - 一个或多个将受益于使用服务类的客户端类

- 声明客户端接口， 描述客户端如何与服务交互

- 创建遵循客户端接口的适配器类。 所有方法暂时都为空。

- 在适配器类中添加一个成员变量用于保存对于服务对象的引用,或者也可以将该变量传递给适配器会更方便。

- 依次实现适配器类客户端接口的所有方法。 适配器会将实际工作委派给服务对象， 自身只负责接口或数据格式的转换。

- 客户端必须通过客户端接口使用适配器，这样一来， 你就可以在不影响客户端代码的情况下修改或扩展适配器。

<h2 id='4'> 4. 优缺点 </h2>

- 单一职责原则: 你可以将接口或数据转换代码从程序主要业务逻辑中分离

-  开闭原则: 只要客户端代码通过客户端接口与适配器进行交互， 你就能在不修改现有客户端代码的情况下在程序中添加新类型的适配器。

<h2 id='5'> 5. 与其它模式关系 </h2>

- 桥接模式通常会于开发前期进行设计， 使你能够将程序的各个部分独立开来以便开发。 另一方面， 适配器模式通常在已有程序中使用， 让相互不兼容的类能很好地合作。

- 适配器可以对已有对象的接口进行修改， 装饰模式则能在不改变对象接口的前提下强化对象功能。 此外， 装饰还支持递归组合， 适配器则无法实现。

- 适配器能为被封装对象提供不同的接口， 代理模式能为对象提供相同的接口， 装饰则能为对象提供加强的接口。

- 外观模式为现有对象定义了一个新接口， 适配器则会试图运用已有的接口。 适配器通常只封装一个对象， 外观通常会作用于整个对象子系统上。

- 桥接、 状态模式和策略模式 （在某种程度上包括适配器） 模式的接口非常相似。 实际上， 它们都基于组合模式——即将工作委派给其他对象， 不过也各自解决了不同的问题。 模式并不只是以特定方式组织代码的配方， 你还可以使用它们来和其他开发者讨论模式所解决的问题。
