package tree

import (
	"github.com/dingkegithub/golanguage/datastrcture/queueandstack"
)

type Node struct {
	Left   *Node
	Right  *Node
	Parent *Node
	Data   interface{}
}

type Binary struct {
	root    *Node
	Compare func(interface{}, interface{}) int
}

func (b *Binary) Empty() bool {
	if b.root == nil {
		return true
	}

	return false
}

func (b *Binary) BuildWithPreMidOrder(pre []interface{}, mid []interface{}) {
	var idx int = 0
	b.root = b.buildWithPreMidOrder(pre, mid, 0, len(pre)-1, &idx)
}

func (b *Binary) buildWithPreMidOrder(pre, mid []interface{}, start, end int, idx *int) *Node {
	if start > end {
		return nil
	}

	// 1> pick and element from pre order and then create new node
	node := &Node{}
	node.Data = pre[*idx]
	// fmt.Println("idx: ", *idx, "start: ", start, "end: ", end, "node: ", node.Data)

	// 2> increment index to pick next element in next recursive call
	*idx += 1

	if start == end {
		return node
	}

	// 3> find index of new node in mid order, we call this index as mid-index
	midIdx := b.search(mid, node.Data, start, end)

	// 4> call _buildTreeWithPreMidOrder for elements before mid-index and make
	// built tree as left subtree of new node
	node.Left = b.buildWithPreMidOrder(pre, mid, start, midIdx-1, idx)

	// 5> call _buildTreeWithPreMidOrder for elements after mid-index and make
	//    built tree as right subtree of new node
	node.Right = b.buildWithPreMidOrder(pre, mid, midIdx+1, end, idx)

	// fmt.Println("----------------------------------------------------")
	// fmt.Println("root->", node.Data)

	// if node.Right != nil {
	// 	fmt.Println("Right-->", node.Right.Data)
	// } else {
	// 	fmt.Println("Right-->", nil)
	// }

	// if node.Left != nil {
	// 	fmt.Println("LEFT-->", node.Left.Data)
	// } else {
	// 	fmt.Println("LEFT-->", nil)
	// }
	// fmt.Println("----------------------------------------------------")

	return node
}

func (b *Binary) BuildWithPostMidOrder(post []interface{}, mid []interface{}) {
	var idx int = 0
	b.root = b.buildWithPreMidOrder(pre, mid, 0, len(pre)-1, &idx)
}

func (b *Binary) buildWithPostMidOrder(pre, mid []interface{}, start, end int, idx *int) *Node {
	if start > end {
		return nil
	}

	// 1> pick and element from pre order and then create new node
	node := &Node{}
	node.Data = pre[*idx]
	// fmt.Println("idx: ", *idx, "start: ", start, "end: ", end, "node: ", node.Data)

	// 2> increment index to pick next element in next recursive call
	*idx += 1

	if start == end {
		return node
	}

	// 3> find index of new node in mid order, we call this index as mid-index
	midIdx := b.search(mid, node.Data, start, end)

	// 4> call _buildTreeWithPreMidOrder for elements before mid-index and make
	// built tree as left subtree of new node
	node.Left = b.buildWithPreMidOrder(pre, mid, start, midIdx-1, idx)

	// 5> call _buildTreeWithPreMidOrder for elements after mid-index and make
	//    built tree as right subtree of new node
	node.Right = b.buildWithPreMidOrder(pre, mid, midIdx+1, end, idx)

	// fmt.Println("----------------------------------------------------")
	// fmt.Println("root->", node.Data)

	// if node.Right != nil {
	// 	fmt.Println("Right-->", node.Right.Data)
	// } else {
	// 	fmt.Println("Right-->", nil)
	// }

	// if node.Left != nil {
	// 	fmt.Println("LEFT-->", node.Left.Data)
	// } else {
	// 	fmt.Println("LEFT-->", nil)
	// }
	// fmt.Println("----------------------------------------------------")

	return node
}

func (b *Binary) search(mid []interface{}, data interface{}, start, end int) int {
	for i := start; i <= end; i++ {
		if b.Compare(mid[i], data) == 0 {
			return i
		}
	}

	return -1
}

func (b *Binary) PreTraverse(res *[]interface{}) {
	b.preTraverse(b.root, res)
}

func (b *Binary) preTraverse(node *Node, res *[]interface{}) {
	if node == nil {
		return
	}

	*res = append(*res, node.Data)
	b.preTraverse(node.Left, res)
	b.preTraverse(node.Right, res)
}

func (b *Binary) MidTraverse(res *[]interface{}) {
	b.midTraverse(b.root, res)
}

func (b *Binary) midTraverse(node *Node, res *[]interface{}) {
	if node == nil {
		return
	}

	b.midTraverse(node.Left, res)
	*res = append(*res, node.Data)
	b.midTraverse(node.Right, res)
}

func (b *Binary) PostTraverse(res *[]interface{}) {
	b.postTraverse(b.root, res)
}

func (b *Binary) postTraverse(node *Node, res *[]interface{}) {
	if node == nil {
		return
	}

	b.postTraverse(node.Left, res)
	b.postTraverse(node.Right, res)
	*res = append(*res, node.Data)
}

func (b *Binary) LevelTraverse(res *[]interface{}) {
	q := queueandstack.New(queueandstack.WithScale())

	q.Put(b.root)

	node, err := q.Get()

	for err == nil {
		*res = append(*res, node.(*Node).Data)
		treeNode := node.(*Node)

		if treeNode.Left != nil {
			q.Put(treeNode.Left)
		}

		if treeNode.Right != nil {
			q.Put(treeNode.Right)
		}
		node, err = q.Get()
	}
}
