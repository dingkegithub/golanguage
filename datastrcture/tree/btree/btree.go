package btree

type Compare func(interface{}, interface{}) int
type Echo func(interface{})

type Node struct {
	keys    []interface{}
	degree  int
	childs  []*Node
	leaf    bool
	compare Compare
	echo    Echo
}

func NewNode(d int, leaf bool, compare Compare) *Node {
	n := &Node{}
	n.degree = d
	n.leaf = leaf
	n.keys = make([]interface{}, 0, (d<<1)-1)
	n.childs = make([]*Node, 0, d<<1)
	n.compare = compare

	return n
}

func (n *Node) Traverse() {
	var i int

	for i = 0; i < len(n.keys); i++ {
		if n.leaf == false {
			n.childs[i].Traverse()
		}
		n.echo(n.keys[i])
	}
	if n.leaf == false {
		n.childs[i].Traverse()
	}
}

func (n *Node) Search(key interface{}) *Node {
	var i int

	//      f0    f1    f2
	/// -------------------------
	//  c0     c1   c2     c3
	for i = 0; i < len(n.keys); i++ {
		if n.compare(key, n.keys[i]) != 1 {
			break
		}
	}

	if n.compare(key, n.keys[i]) == 0 {
		return n
	}

	if n.leaf {
		return nil
	}

	return n.childs[i].Search(key)
}

type Btree struct {
	degree  int
	maxKeys int
	minKeys int
	keys    int
	root    *Node
}

func NewBtree(d int) *Btree {
	return &Btree{
		degree:  d,
		root:    nil,
		minKeys: d - 1,
		maxKeys: 2*d - 1,
	}
}

func (bt *Btree) Search(k interface{}) *Node {
	if bt.root != nil {
		return bt.root.Search(k)
	}

	return nil
}

func (bt *Btree) Traverse() {
	if bt.root != nil {
		bt.root.Traverse()
	}
}
