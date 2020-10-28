package cmd

type Cmd interface {
	Execute()
}

type Receiver interface {
	Do()
}

type ACmd struct {
	receiver Receiver
}

func (c *ACmd) Execute() {
	c.receiver.Do()
}

type BCmd struct {
	receiver Receiver
}

func (b *BCmd) Execute() {
	b.receiver.Do()
}

type Invoker struct {
	c Cmd
}

func (i *Invoker) SetCmd(c Cmd) {
	i.c = c
}

func (i *Invoker) Ok() {
	i.c.Execute()
}
