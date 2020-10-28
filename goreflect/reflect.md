reflect.Type 永远操作的是 类型
reflect.Value 可以读取类型，也可以操作值

所有操作基于已定义的变量，但不能定义新的变量(比如知道了Type类型，想定义变量是不可能的，违反了reflect的三原则)


三原则根本

    value <-------> interface <------------> reflect object

    不是所有的Value都可以设置，要用canSet判断
