>
> [1、评价标准](#1)
>
> [2、面向对象](#2)
>
> [3、设计原则](#3)
>
> [4、设计模式](#4)
>
> [5、编程规范](#5)
>
> [6、代码重构](#6)
>
> [7、面向对象、设计原则、设计模式、编程规范、代码重构联系](#7)
>


<h2 id='1'> 1、评价标准 </h2>

- 可维护性

<p> "代码易维护" 就是指, 在不破坏原有代码设计、不引入新的 bug 的情况下，能够快速地修改或者添加代码, 如果代码
分层清晰、模块化好、高内聚低耦合、遵从基于接口而非实现编程的设计原则等等，那就可能意味着代码易维护。除此之外，
代码的易维护性还跟项目代码量的多少、业务的复杂程度、利用到的技术的复杂程度、文档是否全面、团队成员的开发水平
等诸多因素有关, 如果 bug 容易修复，修改、添加功能能够轻松完成，那我们就可以主观地认为代码对我们来说易维护</p>
<p></p>
<p> “代码不易维护”就是指，修改或者添加代码需要冒着极大的引入新 bug 的风险，并且需要花费很长的时间才能完成 </p>

- 可读性

<p>需要看代码是否符合编码规范、命名是否达意、注释是否详尽、函数是否长短合适、模块划分是否清晰、是否符合高内聚低
耦合等等</p>

<p> 实际上，code review 是一个很好的测验代码可读性的手段。如果你的同事可以轻松地读懂你写的代码，那说明你的代码可
读性很好；如果同事在读你的代码时，有很多疑问，那就说明你的代码可读性有待提高了</p>

- 可扩展性

<p> 表示我们的代码应对未来需求变化的能力, 说直白点就是，代码预留了一些功能扩展点，你可以把新功能代码，直接插到扩展
点上，而不需要因为要添加一个功能而大动干戈，改动大量的原始代码 </p>

> 对修改关闭，对扩展开放

- 灵活性

- 简洁

<p> 思从深，行从简，真正的高手能云淡风轻地用最简单的方法解决最复杂的问题。这也是一个编程老手跟编程新手的本质区别之一。 </p>


- 可复用性

<p> 简单地理解为，尽量减少重复代码的编写，复用已有的代码</p>

- 可测试性

<p></p>


<h2 id='2'> 2、面向对象 </h2>

主流的编程范式或者是编程风格:

- 面向过程

- 面向对象: 封装、抽象、继承、多态

- 函数式编程

7 个大的知识点:

- 面向对象的四大特性：封装、抽象、继承、多态

- 面向对象编程与面向过程编程的区别和联系

- 面向对象分析、面向对象设计、面向对象编程

- 接口和抽象类的区别以及各自的应用场景

- 基于接口而非实现编程的设计思想

- 多用组合少用继承的设计思想

- 面向过程的贫血模型和面向对象的充血模型


<h2 id='3'> 3、设计原则 </h2>

- SRP 单一职责原则: 一个类只能承担一个事情

- OCP 开闭原则：闭修改，开扩展

- LSP 里式替换原则：子类型能够替换它们的基类型

- ISP 接口隔离原则：抽象不持有特定逻辑，应持有实现的公有逻辑

- DIP 依赖倒置原则：高底层不能互相依赖，应同时依赖抽象

- DRY 原则：不做重复的事，不写重复类似的代码

- KISS 原则：keep it simple，尽量简单的代码，让代码更容易被别人理解

- YANGNI 原则：you ain't gonna need it，只着眼必需的功能，不添加认为可能需要的功能 

- LOD 原则：类应减少被外界直接访问的机会，类与类之间避免直接通信


<h2 id='4'> 4、设计模式 </h2>

- 创建型

    - 常用： 单例模式、工厂模式（工厂方法和抽象工厂）、建造者模式。

    - 不常用：原型模式

- 结构性

    - 常用： 代理模式、桥接模式、装饰者模式、适配器模式

    - 不常用：门面模式、组合模式、享元模式

- 行为型

    - 常用： 观察者模式、模板模式、策略模式、职责链模式、迭代器模式、状态模式

    - 不常用：访问者模式、备忘录模式、命令模式、解释器模式、中介模式


<h2 id='5'> 5、编程规范 </h2>

<p>编程规范主要解决的是代码的可读性问题。编码规范相对于设计原则、设计模式，更加具体、更加偏重代码细节,
 比如，如何给变量、类、函数命名，如何写代码注释，函数不宜过长、参数不能过多等等</p>


<p> 对于编码规范，考虑到很多书籍已经讲得很好了（比如《重构》《代码大全》《代码整洁之道》等）。而且，每
条编码规范都非常简单、非常明确，比较偏向于记忆，你只要照着来做可以。</p>

<h2 id='6'> 6、代码重构 </h2>

<p>对于重构这部分内容，你需要掌握以下几个知识点：</p>

- 重构的目的（why）、对象（what）、时机（when）、方法（how）；

- 保证重构不出错的技术手段：单元测试和代码的可测试性；

- 两种不同规模的重构：大重构（大规模高层次）和小重构（小规模低层次）


<h2 id='7'> 7、五者之间的联系 </h2>

- 面向对象编程因为其具有丰富的特性（封装、抽象、继承、多态），可以实现很多复杂的设计思路，是很多设计原则、设计模式等编码实现的基础。

- 设计原则是指导我们代码设计的一些经验总结，对于某些场景下，是否应该应用某种设计模式，具有指导意义。比如，“开闭原则”是很多设计模式（策略、模板等）的指导原则。

- 设计模式是针对软件开发中经常遇到的一些设计问题，总结出来的一套解决方案或者设计思路。应用设计模式的主要目的是提高代码的可扩展性。从抽象程度上来讲，设计原则比设计模式更抽象。设计模式更加具体、更加可执行。

- 编程规范主要解决的是代码的可读性问题。编码规范相对于设计原则、设计模式，更加具体、更加偏重代码细节、更加能落地。持续的小重构依赖的理论基础主要就是编程规范。

- 重构作为保持代码质量不下降的有效手段，利用的就是面向对象、设计原则、设计模式、编码规范这些理论。

实际上，面向对象、设计原则、设计模式、编程规范、代码重构，这五者都是保持或者提高代码质量的方法论，本质上都是服务于编写高质量代码这一件事的