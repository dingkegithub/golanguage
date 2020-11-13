

- 向关闭的 chan 写数据，报：panic: send on closed channed [recovered]

- 向关闭的 chan 读数据，返回对应类型的0值，int->0, bool->false, string->"", interface->nil

- 向为初始化的 chan 读写数据将阻塞
